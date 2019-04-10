package api

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

type voted struct {
	Voted bool
}

func (api *PrimaryAPI) VerifySelf(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(403)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(c.Value)

	me, err := api.Election.GetVoterByID(id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	ret, _ := api.Election.HasVoted(*me)
	if ret != nil {
		if *ret {
			WriteJSON(w, &voted{
				Voted: true,
			})
			return
		}

	}
	WriteJSON(w, &voted{
		Voted: false,
	})
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
	print("here")
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
	fmt.Println(voter)
	ret, err := api.Election.HasVoted(*voter)
	fmt.Println(*ret)
	if *ret {
		w.Write([]byte("voted"))
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

func (api *PrimaryAPI) CastVote(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var s []idpost
	err := dec.Decode(&s)
	if err != nil {

	}

	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(403)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(c.Value)
	me, _ := api.Election.GetVoterByID(id)
	voted, err := api.Election.HasVoted(*me)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	if *voted {
		w.WriteHeader(400)
		w.Write([]byte("This voter has alredy voted"))
		return
	}

	for _, val := range s {
		candidate, err := api.Election.GetCandidateByID(val.ID)

		vote := &models.Vote{
			StudentID: me.StudentID,
			Candidate: candidate.ID,
		}
		_ = vote
		vote.HashVote(me)
		fmt.Println(vote.Hash)

		err = api.Election.CastVote(me, vote)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))

		}
	}
	err = api.Database.SetVoted(*me)

}