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

// HashVote calcluate if it does not exist some hash of the vote, must be repeatable
func (v *Vote) HashVote(voter *Voter) {
	// TODO: implement hash
	h := fnv.New32a()
	h.Write([]byte(strconv.Itoa(voter.StudentID + v.Candidate)))
	v.Hash = fmt.Sprint(h.Sum32())
}
