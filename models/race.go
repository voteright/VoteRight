package models

// Race contains all information for a single race
type Race struct {
	ID         int `storm:"id,increment"`
	Name       string
	Votes      []Vote
	Candidates []int
}

// RaceWithCandidates is the value that is returned to the frontend, containing candidates and no votes
type RaceWithCandidates struct {
	ID         int
	Name       string
	Candidates []Candidate
}
