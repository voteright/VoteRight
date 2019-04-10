package database

import (
	"fmt"

	"github.com/asdine/storm"
	"github.com/voteright/voteright/models"
)

// StormDB represents a connection to the stormdb
type StormDB struct {
	File string
	DB   *storm.DB
}

// Connect connects to the database
func (s *StormDB) Connect() error {
	DB, err := storm.Open(s.File)
	s.DB = DB
	return err
}

// Close closes the connection to the databae
func (s *StormDB) Close() error {
	return s.DB.Close()
}

// StoreCandidate stores a candidate in the database
func (s *StormDB) StoreCandidate(candidate models.Candidate) error {
	return s.DB.Save(&candidate)
}

// StoreCandidates stores candidates in the database
func (s *StormDB) StoreCandidates(candidates []models.Candidate) error {
	n, err := s.DB.Begin(true)
	if err != nil {
		return err
	}

	for _, c := range candidates {
		err := n.Save(&c)
		if err != nil {
			n.Rollback()
			return err
		}
	}

	return n.Commit()

}

// GetAllCandidates returns all candidates in the database
func (s *StormDB) GetAllCandidates() ([]models.Candidate, error) {
	var ret []models.Candidate
	err := s.DB.All(&ret)
	return ret, err
}

// StoreCohort stores a cohort in the database
func (s *StormDB) StoreCohort(cohort models.Cohort) error {
	return s.DB.Save(&cohort)
}

// StoreCohorts stores cohorts in the database
func (s *StormDB) StoreCohorts(cohorts []models.Cohort) error {
	tx, err := s.DB.Begin(true)
	if err != nil {
		return err
	}
	for _, cohort := range cohorts {
		tx.Save(&cohort)
	}

	return tx.Commit()
}

// GetAllCohorts returns all cohorts in the database
func (s *StormDB) GetAllCohorts() ([]models.Cohort, error) {
	var ret []models.Cohort
	err := s.DB.All(&ret)
	return ret, err
}

// StoreVote stores a vote in the database
func (s *StormDB) StoreVote(vote models.Vote) error {
	return s.DB.Save(&vote)
}

// GetAllVotes returns all votes in the database
func (s *StormDB) GetAllVotes() ([]models.Vote, error) {
	var ret []models.Vote
	err := s.DB.All(&ret)
	return ret, err
}

// StoreVoter stores a voter in the database
func (s *StormDB) StoreVoter(voter models.Voter) error {

	return s.DB.Save(&voter)
}

// Voted is a storm struct
type Voted struct {
	ID        int `storm:"id,increment"`
	StudentID int
}

// SetVoted sets if a voter has voted
func (s *StormDB) SetVoted(voter models.Voter) error {
	val := Voted{
		StudentID: voter.StudentID,
	}
	return s.DB.Save(&val)
}

// HasVoted checks if a voter has voted
func (s *StormDB) HasVoted(voter models.Voter) (*bool, error) {

	var voted []Voted

	s.DB.All(&voted)

	var voters []models.Voter
	s.DB.All(&voters)
	retval := false
	for _, val := range voted {
		if voter.StudentID == val.StudentID {
			retval = true
			return &retval, nil
		}
	}
	return &retval, nil
}

// StoreVoters stores voters in the database
func (s *StormDB) StoreVoters(voters []models.Voter) error {
	tx, _ := s.DB.Begin(true)
	for _, voter := range voters {
		tx.Save(&voter)
	}

	return tx.Commit()

}

// GetAllVoters returns all voters in the database
func (s *StormDB) GetAllVoters() ([]models.Voter, error) {
	var voters []models.Voter
	err := s.DB.All(&voters)
	fmt.Println(len(voters))
	return voters, err
}

// StoreRace stores voters in the database
func (s *StormDB) StoreRace(race models.Race) error {
	return s.DB.Save(&race)

}

// GetAllRaces returns all voters in the database
func (s *StormDB) GetAllRaces() ([]models.Race, error) {
	var races []models.Race
	err := s.DB.All(&races)
	return races, err
}
