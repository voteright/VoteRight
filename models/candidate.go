package models

// Candidate represents one candidate to vote on
type Candidate struct {
	Name   string
	Cohort int
	ID     int
}

// CandidateVotes returns the number of votes for a candidate
type CandidateVotes struct {
	Candidate Candidate
	Votes     int
}
