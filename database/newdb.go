package database

import (
	"fmt"

	"github.com/asdine/storm"
	"github.com/voteright/voteright/models"
)

/*
	This is the code for the database connection to stormdb, it allows storing
	Models and handling errors in a standard way.
*/

// StormDB represents a connection to the stormdb
type StormDB struct {
	File string
	DB   *storm.DB
}

// Dump represents a dump of the database, used by importer and exporter
type Dump struct {
	Voters     []models.Voter
	Votes      []models.Vote
	Candidates []models.Candidate
	Cohorts    []models.Cohort
	Races      []models.Race
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

// Setmodels. sets if a voter has voted
func (s *StormDB) SetVoted(voter models.Voter) error {
	val := models.Voted{
		StudentID: voter.StudentID,
	}
	return s.DB.Save(&val)
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

// StoreRace stores races in the database
func (s *StormDB) StoreRace(race models.Race) error {
	return s.DB.Save(&race)

}

// GetAllRaces returns all races in the database
func (s *StormDB) GetAllRaces() ([]models.Race, error) {
	var races []models.Race
	err := s.DB.All(&races)
	return races, err
}

// StoreIntegrityViolation stores integrity violations in the database
func (s *StormDB) StoreIntegrityViolation(IntegrityViolation models.IntegrityViolation) error {
	return s.DB.Save(&IntegrityViolation)

}

// GetAllIntegrityViolations returns all integriy violations in the database
func (s *StormDB) GetAllIntegrityViolations() ([]models.IntegrityViolation, error) {
	var races []models.IntegrityViolation
	err := s.DB.All(&races)
	return races, err
}

// StoreBallot stores a ballot in the database
func (s *StormDB) StoreBallot(Ballot models.Ballot) error {
	return s.DB.Save(&Ballot)
}

// GetAllBallots gets all ballots in the database
func (s *StormDB) GetAllBallots() ([]models.Ballot, error) {
	var ballots []models.Ballot
	err := s.DB.All(&ballots)
	return ballots, err
}
