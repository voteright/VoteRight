package election

import (
	"github.com/voteright/voteright/database"
	"github.com/voteright/voteright/models"
)

// Election represents entities required to run an election, will eventually contain
// the database interaction, and potentially remote servers
type Election struct {
	db *database.Database
}

// New returns a new election struct
func New(db *database.Database) *Election {
	return &Election{
		db: db,
	}
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
