package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/voteright/voteright/config"
	"github.com/voteright/voteright/models"
)

// Database represents a connection to the database
type Database struct {
	db *sql.DB
}

// Dump represents a dump of the database, used by importer and exporter
type Dump struct {
	Voters     []models.Voter
	Votes      []models.Vote
	Candidates []models.Candidate
	Cohorts    []models.Cohort
}

// New establishes the connection to the database and returns the driver
func New(cfg *config.Config) (*Database, error) {
	dbs := &Database{}
	db, err := sql.Open("sqlite3", cfg.DatabaseFile)
	if err != nil {
		return nil, err
	}
	dbs.db = db
	return dbs, nil
}

// QueryStatement Run sql query, should probably use wrappers instead of this directly
func (d *Database) QueryStatement(statement string, vals ...interface{}) (*sql.Rows, error) {
	rows, err := d.db.Query(statement, vals)
	return rows, err
}

// ExecStatement Run sql statement, should probably use wrappers instead of this directly
func (d *Database) ExecStatement(statement string) (sql.Result, error) {
	st, err := d.db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	res, err := st.Exec()
	if err != nil {
		return nil, err
	}

	return res, nil

}

// StoreVoter stores a voter in the database
func (d *Database) StoreVoter(voter models.Voter) error {
	st, err := d.db.Prepare("INSERT INTO voters VALUES (?,?,?)")
	if err != nil {
		return err
	}
	_, err = st.Exec(voter.StudentID, voter.Name, voter.Cohort)
	if err != nil {
		return err
	}
	return nil
}

// StoreVoters stores voters in the database
func (d *Database) StoreVoters(voters []models.Voter) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	for _, voter := range voters {
		st, err := tx.Prepare("INSERT INTO voters VALUES (?,?,?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = st.Exec(voter.StudentID, voter.Name, voter.Cohort)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}

// StoreCohort stores a cohort in the database
func (d *Database) StoreCohort(cohort models.Cohort) error {
	st, err := d.db.Prepare("INSERT INTO cohorts VALUES (?,?)")
	if err != nil {
		return err
	}
	_, err = st.Exec(cohort.ID, cohort.Name)
	if err != nil {
		return err
	}
	return nil
}

// StoreCohorts stores cohorts in the database
func (d *Database) StoreCohorts(cohorts []models.Cohort) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	for _, cohort := range cohorts {
		st, err := tx.Prepare("INSERT INTO cohorts VALUES (?,?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = st.Exec(cohort.ID, cohort.Name)

		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}

// StoreCandidate stores a candidate in the database
func (d *Database) StoreCandidate(candidate models.Candidate) error {
	st, err := d.db.Prepare("INSERT INTO candidates VALUES (?,?,?)")
	if err != nil {
		return err
	}
	_, err = st.Exec(candidate.ID, candidate.Name, candidate.Cohort)
	if err != nil {
		return err
	}
	return nil
}

// StoreCandidates stores candidates in the database
func (d *Database) StoreCandidates(candidates []models.Candidate) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	for _, candidate := range candidates {
		st, err := d.db.Prepare("INSERT INTO candidates VALUES (?,?,?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = st.Exec(candidate.ID, candidate.Name, candidate.Cohort)

		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}

// StoreVote stores a vote in the database
func (d *Database) StoreVote(vote models.Vote) error {
	st, err := d.db.Prepare("INSERT INTO votes VALUES (?,?)")
	if err != nil {
		return err
	}
	_, err = st.Exec(vote.Hash, vote.Candidate)
	if err != nil {
		return err
	}
	return nil
}

// GetAllVoters returns all voters in the database
func (d *Database) GetAllVoters() ([]models.Voter, error) {
	res, err := d.db.Query("SELECT * from voters")
	if err != nil {
		return nil, err
	}
	defer res.Close()
	ret := []models.Voter{}

	for res.Next() {
		v := &models.Voter{}
		err := res.Scan(&v.StudentID, &v.Name, &v.Cohort)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *v)
	}
	return ret, nil
}

// GetAllCohorts returns all cohorts in the database
func (d *Database) GetAllCohorts() ([]models.Cohort, error) {
	res, err := d.db.Query("SELECT * from cohorts")
	if err != nil {
		return nil, err
	}
	defer res.Close()
	ret := []models.Cohort{}

	for res.Next() {
		v := &models.Cohort{}
		err := res.Scan(&v.ID, &v.Name)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *v)
	}
	return ret, nil
}

// GetAllCandidates returns all candidates in the database
func (d *Database) GetAllCandidates() ([]models.Candidate, error) {
	res, err := d.db.Query("SELECT * from candidates")
	if err != nil {
		return nil, err
	}
	defer res.Close()
	ret := []models.Candidate{}

	for res.Next() {
		v := &models.Candidate{}
		err := res.Scan(&v.ID, &v.Name, &v.Cohort)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *v)
	}
	return ret, nil
}

// GetAllVotes returns all votes in the database
func (d *Database) GetAllVotes() ([]models.Vote, error) {
	res, err := d.db.Query("SELECT * from votes")
	if err != nil {
		return nil, err
	}
	defer res.Close()
	ret := []models.Vote{}

	for res.Next() {
		v := &models.Vote{}
		err := res.Scan(&v.Hash, &v.Candidate)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *v)
	}
	return ret, nil
}
