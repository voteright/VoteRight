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

// Candidate represents one candidate to vote on
type Candidate struct {
	Name   string
	Cohort string
	ID     int
}

// Voter represents a voter in the election
type Voter struct {
	StudentID int
	Cohort    string
}

// Vote represents a single vote
type Vote struct {
	Hash      string
	Candidate int
}
