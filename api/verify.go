package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"

	"github.com/voteright/voteright/models"
)

// HandleVerificationPost handles the posting of a ballot for verficiation
func (api *PrimaryAPI) HandleVerificationPost(w http.ResponseWriter, r *http.Request) {
	var b models.Ballot
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&b)
	fmt.Println(b)
	if err != nil {
		w.Write([]byte("Invalid format"))
		w.WriteHeader(403)
		api.Database.StoreIntegrityViolation(models.IntegrityViolation{
			Message: "Invalid ballot posted",
			Time:    time.Now(),
		})
		return
	}
	err = api.Database.StoreBallot(b)
	if err != nil {
		fmt.Println(err)
	}

}

// HandleVerificationCounts handles returning counts for verification
func (api *PrimaryAPI) HandleVerificationCounts(w http.ResponseWriter, r *http.Request) {
	x := []models.CandidateVotes{}
	ballots, err := api.Database.GetAllBallots()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	// Populate the array of candidate votes to count totals
	for _, ballot := range ballots {
		for _, candidate := range ballot.Candidates {
			found := false
			for _, c := range x {
				if c.Candidate == candidate {
					found = true
					break
				}
			}
			if !found {
				x = append(x, models.CandidateVotes{Candidate: candidate})
			}
		}
	}

	for _, ballot := range ballots {
		for _, ballotCandidate := range ballot.Candidates {
			for i := range x {
				if ballotCandidate == x[i].Candidate {
					x[i].Votes++
				}
			}
		}

	}

	WriteJSON(w, x)

}

func (api *PrimaryAPI) GetBallot(w http.ResponseWriter, r *http.Request) {
	idstr := chi.URLParam(r, "id")
	idNum, err := strconv.Atoi(idstr)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}
	ballots, _ := api.Database.GetAllBallots()
	fmt.Println(ballots)
	for _, ballot := range ballots {
		if idNum == ballot.RandomID {
			WriteJSON(w, ballot)
			return
		}
	}
	w.Write([]byte("Not found"))
	w.WriteHeader(404)
}
