package election

import (
	"os"
	"testing"

	"github.com/voteright/voteright/database"
	"github.com/voteright/voteright/models"
)

// Test the HasVoted Function
func TestElection_HasVoted(t *testing.T) {
	testDatabase := database.StormDB{
		File: "election_testdb.db",
	}
	err := testDatabase.Connect()
	if err != nil {
		t.Errorf("Couldn't connect to database. Error: %s", err.Error())
	}
	testVoter := models.Voter{
		StudentID: 1,
		Cohort:    1,
		Name:      "Prof Sturman",
	}
	testVoterWontVote := models.Voter{
		StudentID: 2,
		Cohort:    1,
		Name:      "Prof Goldschmidt",
	}
	testCandidate := models.Candidate{
		ID:     1,
		Cohort: 1,
		Name:   "Joey Lyon",
	}

	err = testDatabase.StoreVoter(testVoter)
	if err != nil {
		t.Errorf("Couldn't add test voter to database")
	}
	err = testDatabase.StoreVoter(testVoterWontVote)
	if err != nil {
		t.Errorf("Couldn't add test voter to database")
	}
	err = testDatabase.StoreCandidate(testCandidate)
	if err != nil {
		t.Errorf("Couldn't add test candidate to database")
	}

	e := New(&testDatabase, false, []string{})
	// Begin testing HasVoted function
	ret, err := e.HasVoted(testVoter)
	if err != nil {
		t.Errorf("unexpected error in checking if voter has voted")
	}
	if *ret {
		t.Errorf("HasVoted returned true when a voter hasn't voted")
	}

	vote := &models.Vote{
		Candidate: 1,
		StudentID: 1,
	}
	vote.HashVote(&testVoter)
	e.CastVotes(&testVoter, []models.Vote{*vote})
	ret, err = e.HasVoted(testVoter)
	if err != nil {
		t.Errorf("unexpected error in checking if voter has voted")
	}
	if *ret == false {
		t.Errorf("HasVoted returned false when a voter has voted")
	}

	ret, err = e.HasVoted(testVoterWontVote)
	if err != nil {
		t.Errorf("unexpected error in checking if voter has voted")
	}
	if *ret {
		t.Errorf("HasVoted returned true when a voter has not voted")
	}
	err = os.Remove("election_testdb.db")
	if err != nil {
		t.Log("Cleanup failed")
	}
}
