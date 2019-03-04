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
