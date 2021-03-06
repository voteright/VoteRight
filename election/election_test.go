package election

import (
	"fmt"
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

func TestElection_GetCandidateByID(t *testing.T) {
	testDatabase := database.StormDB{
		File: "election_testdb.db",
	}
	err := testDatabase.Connect()
	if err != nil {
		t.Errorf("Couldn't connect to database. Error: %s", err.Error())
	}

	testCandidate := models.Candidate{
		ID:     1,
		Cohort: 1,
		Name:   "Joey Lyon",
	}

	testCandidateTwo := models.Candidate{
		ID:     2,
		Cohort: 1,
		Name:   "Grace Roller",
	}

	err = testDatabase.StoreCandidate(testCandidate)
	if err != nil {
		t.Errorf("Couldn't add test candidate to database")
	}
	err = testDatabase.StoreCandidate(testCandidateTwo)
	if err != nil {
		t.Errorf("Couldn't add test candidate to database")
	}

	e := New(&testDatabase, false, []string{})

	candidate, err := e.GetCandidateByID(1)

	if err != nil {
		t.Errorf("Unexpected error reading candidates")
	}

	if candidate.Name != "Joey Lyon" || candidate.ID != 1 {
		t.Errorf("Candidate information not as expected")
	}

	candidate, err = e.GetCandidateByID(500)
	if err != nil {
		t.Errorf("Unexpected error reading candidates")
	}
	if candidate != nil {
		t.Errorf("Got a candidate for an invalid id")
	}
	err = os.Remove("election_testdb.db")
	if err != nil {
		t.Log("Cleanup failed")
	}
}

func TestElection_GetCandidateVoteCounts(t *testing.T) {
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

	votes, err := e.GetCandidateVoteCounts()
	if err != nil {
		t.Errorf("Error reading candidate votes")
	}
	v := *votes
	if len(v) != 1 {
		t.Errorf("vote array length not long enough")
	} else {
		fmt.Println(v[0])
		if v[0].Votes != 0 {
			t.Errorf("Vote count inacurate before voting")
		}
	}

	vote.HashVote(&testVoter)
	e.CastVotes(&testVoter, []models.Vote{*vote})

	votes, err = e.GetCandidateVoteCounts()
	if err != nil {
		t.Errorf("Error reading candidate votes")
	}
	v = *votes
	if len(v) != 1 {
		t.Errorf("vote array length not long enough")
	} else {
		fmt.Println(v[0])
		if v[0].Votes != 1 {
			t.Errorf("Vote count inacurate after voting")
		}
	}

	err = os.Remove("election_testdb.db")
	if err != nil {
		t.Log("Cleanup failed")
	}
}

func TestElection_CheckVerificationCountsMatch(t *testing.T) {
	v := []models.CandidateVotes{
		{
			Candidate: models.Candidate{
				Name: "Test Candidate",
				ID:   1,
			},
			Votes: 1,
		},
		{
			Candidate: models.Candidate{
				Name: "Second Candidate",
				ID:   2,
			},
			Votes: 2,
		},
	}
	v2 := []models.CandidateVotes{
		{
			Candidate: models.Candidate{
				Name: "Test Candidate",
				ID:   1,
			},
			Votes: 1,
		},
		{
			Candidate: models.Candidate{
				Name: "Second Candidate",
				ID:   2,
			},
			Votes: 2,
		},
	}
	e := Election{}
	par := [][]models.CandidateVotes{v, v2}

	ret := e.CheckVerificationCountsMatch(par)
	if !ret {
		t.Errorf("Vote totals do not match but they should")
	}
	par2 := [][]models.CandidateVotes{v, []models.CandidateVotes{}}

	ret = e.CheckVerificationCountsMatch(par2)
	if ret {
		t.Errorf("Vote totals match but they shouldnt")
	}
	v2[0].Votes = 5
	par = [][]models.CandidateVotes{v, v2}

	ret = e.CheckVerificationCountsMatch(par)
	if ret {
		t.Errorf("Vote totals do not match but they should")
	}
}

func TestElection_GetAllRaces(t *testing.T) {
	testDatabase := database.StormDB{
		File: "election_testdb.db",
	}
	err := testDatabase.Connect()
	if err != nil {
		t.Errorf("Couldn't connect to database. Error: %s", err.Error())
	}

	testCandidate := models.Candidate{
		ID:     1,
		Cohort: 1,
		Name:   "Joey Lyon",
	}
	testCandidateTwo := models.Candidate{
		ID:     2,
		Cohort: 1,
		Name:   "Joey Lyon Two",
	}

	err = testDatabase.StoreCandidate(testCandidate)
	if err != nil {
		t.Errorf("Couldn't add test candidate to database")
	}
	err = testDatabase.StoreCandidate(testCandidateTwo)
	if err != nil {
		t.Errorf("Couldn't add test candidate to database")
	}
	race := models.Race{
		ID:         1,
		Name:       "Test",
		Candidates: []int{1, 2},
	}

	e := New(&testDatabase, false, []string{})
	resp, err := e.GetAllRaces()
	if err != nil {
		t.Errorf("Couldn't add test candidate to database")
	}
	if len(resp) != 0 {
		t.Errorf("Got races when there were none in the database")
	}
	e.db.StoreRace(race)

	err = os.Remove("election_testdb.db")
	if err != nil {
		t.Log("Cleanup failed")
	}
}

func TestElection_GetVoterByID(t *testing.T) {
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

	err = testDatabase.StoreVoter(testVoter)
	if err != nil {
		t.Errorf("Couldn't add test voter to database")
	}
	err = testDatabase.StoreVoter(testVoterWontVote)
	if err != nil {
		t.Errorf("Couldn't add test voter to database")
	}

	e := New(&testDatabase, false, []string{})

	voter, err := e.GetVoterByID(1)
	if err != nil {
		t.Errorf("Unexpected error in grabbing voter")
	}
	if *voter != testVoter {
		t.Errorf("Failed to retreive proper voter")
	}
	voter, err = e.GetVoterByID(400)
	if err != nil {
		t.Errorf("unexpected error grabbing voter")
	}
	if voter != nil {
		t.Errorf("voter was returned for invalid id")
	}
	err = os.Remove("election_testdb.db")
	if err != nil {
		t.Log("Cleanup failed")
	}
}
