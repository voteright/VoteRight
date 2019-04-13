package election

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/voteright/voteright/database"
	"github.com/voteright/voteright/models"
)

// Election represents entities required to run an election, will eventually contain
// the database interaction, and potentially remote servers
type Election struct {
	db                  *database.StormDB
	Verification        bool
	VerificationServers []string
}

// New returns a new election struct
func New(db *database.StormDB, Verification bool, VerificationServers []string) *Election {
	return &Election{
		db:                  db,
		Verification:        Verification,
		VerificationServers: VerificationServers,
	}
}

func (e *Election) GetCandidateByID(id int) (*models.Candidate, error) {
	candidates, err := e.db.GetAllCandidates()
	var ret *models.Candidate
	for _, c := range candidates {
		if c.ID == id {
			ret = &c
			break
		}
	}
	return ret, err

}

// HasVoted returns true if the voter has voted, false if they have not, and nil, error if an error occurs
func (e *Election) HasVoted(voter models.Voter) (*bool, error) {
	var voted []models.Voted

	e.db.DB.All(&voted)

	var voters []models.Voter
	e.db.DB.All(&voters)
	retval := false
	for _, val := range voted {
		if voter.StudentID == val.StudentID {
			retval = true
			return &retval, nil
		}
	}
	return &retval, nil
}

// GetVoterByID returns the voter with the given id
func (e *Election) GetVoterByID(id int) (*models.Voter, error) {
	voters, err := e.db.GetAllVoters()
	var ret *models.Voter

	for _, v := range voters {
		if v.StudentID == id {
			ret = &v
			break
		}
	}
	return ret, err

}

// GetAllCohorts returns all voters in the database
func (e *Election) GetAllCohorts() ([]models.Cohort, error) {
	return e.db.GetAllCohorts()
}

// GetAllVoters returns all voters in the database
func (e *Election) GetAllVoters() ([]models.Voter, error) {
	return e.db.GetAllVoters()
}

// GetAllCandidates returns all candidates in the database
func (e *Election) GetAllCandidates() ([]models.Candidate, error) {
	return e.db.GetAllCandidates()
}

// GetAllRaces returns all races in the database
func (e *Election) GetAllRaces() ([]models.RaceWithCandidates, error) {
	candidates, err := e.db.GetAllCandidates()
	if err != nil {
		return nil, err
	}
	races, err := e.db.GetAllRaces()
	if err != nil {
		return nil, err
	}
	var ret []models.RaceWithCandidates
	for _, race := range races {
		r1 := models.RaceWithCandidates{
			Name: race.Name,
			ID:   race.ID,
		}
		for _, candidate := range candidates {
			for _, id := range race.Candidates {
				if id == candidate.ID {
					r1.Candidates = append(r1.Candidates, candidate)
				}
			}
		}
		ret = append(ret, r1)
	}
	return ret, nil
}

// GetCandidateVoteCounts returns the candidates with their vote totals
func (e *Election) GetCandidateVoteCounts() (*[]models.CandidateVotes, error) {
	ret := []models.CandidateVotes{}
	c, err := e.GetAllCandidates()
	if err != nil {
		return nil, err
	}
	v, err := e.db.GetAllVotes()
	if err != nil {
		return nil, err
	}

	// iterate through all candidates and all votes and tally them up
	for _, cand := range c {
		curr := models.CandidateVotes{
			Candidate: cand,
			Votes:     0,
		}
		for _, vote := range v {
			if cand.ID == vote.Candidate {
				curr.Votes = curr.Votes + 1
			}
		}
		ret = append(ret, curr)
	}

	return &ret, nil
}

// GetCountsFromVerificationServers returns all of the counts from the verfication servers
func (e *Election) GetCountsFromVerificationServers() ([][]models.CandidateVotes, error) {
	var vsVals [][]models.CandidateVotes
	for _, s := range e.VerificationServers {
		r, err := http.Get(s + "/integrity/totals")

		if err != nil {
			e.db.StoreIntegrityViolation(models.IntegrityViolation{
				Message: "[SEVERE]: Could not communicate with verification server " + s,
				Time:    time.Now(),
			})
			return nil, err
		}
		var ret []models.CandidateVotes
		dec := json.NewDecoder(r.Body)
		err = dec.Decode(&ret)
		if err != nil {
			e.db.StoreIntegrityViolation(models.IntegrityViolation{
				Message: "[SEVERE]: Could not decode response from verification server " + s,
				Time:    time.Now(),
			})
			fmt.Println("failed to decode", err.Error())
			return nil, err
		}
		vsVals = append(vsVals, ret)
	}
	fmt.Println(vsVals)
	return vsVals, nil
}

// ByName is a type for sorting the slices of votes returned by verification servers
type ByName []models.CandidateVotes

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Candidate.Name < a[j].Candidate.Name }

// CheckVerificationCountsMatch checks if the veification counts from all servers match
func (e *Election) CheckVerificationCountsMatch(votes [][]models.CandidateVotes) bool {

	for _, i := range votes {
		sort.Sort(ByName(i))
		for _, k := range i {
			for _, j := range votes {
				sort.Sort(ByName(j))

				// If one server has a candidate the others dont return false
				if len(i) != len(j) {
					return false
				}
				// Slices should be identical after sorting
				for _, l := range j {
					if l != k {
						return false
					}
				}
			}
		}
	}
	return true
}
