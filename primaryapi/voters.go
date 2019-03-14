package primaryapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/voteright/voteright/models"
)

// GetAllVoters gets the list of the voters in the election
func (api *PrimaryAPI) GetAllVoters(w http.ResponseWriter, r *http.Request) {
	// TODO: Needs auth
	// if (!admin) return 403

	v, err := api.Election.GetAllVoters()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	WriteJSON(w, v)

}

type idpost struct {
	ID int
}

func (api *PrimaryAPI) ValidateVoter(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var s idpost
	err := dec.Decode(&s)
	if err != nil {
		fmt.Println("Error", err.Error())
		w.WriteHeader(400)
		w.Write([]byte("Could not understand request body"))
		return
	}
	voter, err := api.Election.GetVoterByID(s.ID)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Failed to read the database"))
		return
	}

	WriteJSON(w, voter)
}

// TODO: finish this
func (api *PrimaryAPI) LoginVoter(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var s models.Voter
	err := dec.Decode(&s)
	if err != nil {
		fmt.Println("Error", err.Error())
		w.WriteHeader(400)
		w.Write([]byte("Could not understand request body"))
		return
	}
	voter, err := api.Election.GetVoterByID(s.StudentID)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Failed to read dbs"))
		return
	}
	if *voter != s {
		w.WriteHeader(403)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   strconv.Itoa(voter.StudentID),
		Expires: time.Now().Add(5 * 60 * time.Second),
	})
	fmt.Println("voter logged in", voter.Name)
	WriteJSON(w, voter)
}
