package models

import "testing"

func TestVote_HashVote(t *testing.T) {
	v1 := Voter{
		StudentID: 1,
	}
	v := &Vote{
		Candidate: 1,
		StudentID: 1,
		ID:        1,
	}
	v.HashVote(&v1)
	if v.Hash != "923577301" {
		t.Errorf("Incorrect hash, expected 923577301 got %s\n", v.Hash)
	}

}
