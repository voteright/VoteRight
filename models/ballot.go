package models

import "math/rand"

// Ballot represents all of the votes cast by a voter, indexed by a unique id that is given to the voter upon voting
type Ballot struct {
	RandomID   int `storm:"id"`
	Candidates []Candidate
}

// GenerateRandomID will generate the ID for the ballot
func (b *Ballot) GenerateRandomID() {
	b.RandomID = rand.Int()
	for b.RandomID < 100000 {
		b.RandomID = rand.Int() + 1
	}
	b.RandomID /= 10000
}
