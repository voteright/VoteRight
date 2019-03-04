package election

import "github.com/voteright/voteright/database"

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
