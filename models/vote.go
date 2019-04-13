package models

import (
	"fmt"
	"hash/fnv"
	"strconv"
)

// Vote represents a single vote
type Vote struct {
	Hash      string
	Candidate int
	StudentID int
	ID        int `storm:"id,increment"`
}

// Voted is a storm struct for storing if a user has voted, it is here because it has not
type Voted struct {
	ID        int `storm:"id,increment"`
	StudentID int
}

// HashVote calcluate if it does not exist some hash of the vote, must be repeatable
func (v *Vote) HashVote(voter *Voter) {
	h := fnv.New32a()
	h.Write([]byte(strconv.Itoa(voter.StudentID + v.Candidate)))
	v.Hash = fmt.Sprint(h.Sum32())
}
