package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/voteright/voteright/config"
)

// Database represents a connection to the database
type Database struct {
	db *sql.DB
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
