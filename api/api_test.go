package api

import (
	"net/http"
	"os"
	"testing"

	"github.com/voteright/voteright/config"
	"github.com/voteright/voteright/database"
	"github.com/voteright/voteright/election"
)

func TestVariousEndpoints(t *testing.T) {
	testDatabase := database.StormDB{
		File: "election_testdb.db",
	}
	err := testDatabase.Connect()
	if err != nil {
		t.Errorf("Couldn't connect to database. Error: %s", err.Error())
	}

	e := election.New(&testDatabase, false, []string{})
	c := config.Config{
		ListenURL: "0.0.0.0:8080",
	}
	a := New(&c, e, &testDatabase)
	go a.Serve()
	// <-time.After(5 * time.Second)
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Error("failed to hit api endpoint")
	} else if resp.StatusCode != 200 {
		t.Error("Unexpected status code main")
	}
	resp, err = http.Get("http://localhost:8080/candidates/")
	if err != nil {
		t.Error("failed to hit api endpoint")
	} else if resp.StatusCode != 200 {
		t.Errorf("Unexpected status code candidates")
	}
	resp, err = http.Get("http://localhost:8080/candidates/votes")
	if err != nil {
		t.Error("failed to hit api endpoint", err.Error())
	} else if resp.StatusCode != 200 {
		t.Errorf("Unexpected status code votes")
	}
	err = os.Remove("election_testdb.db")
	if err != nil {
		t.Log("Cleanup failed")
	}
}

// func Test_Candidates_Endpoints(t *testing.T) {
// testDatabase := database.StormDB{
// 	File: "election_testdb.db",
// }
// err := testDatabase.Connect()
// if err != nil {
// 	t.Errorf("Couldn't connect to database. Error: %s", err.Error())
// }

// testCandidate := models.Candidate{
// 	ID:     1,
// 	Cohort: 1,
// 	Name:   "Joey Lyon",
// }

// testCandidateTwo := models.Candidate{
// 	ID:     2,
// 	Cohort: 1,
// 	Name:   "Grace Roller",
// }

// err = testDatabase.StoreCandidate(testCandidate)
// if err != nil {
// 	t.Errorf("Couldn't add test candidate to database")
// }
// err = testDatabase.StoreCandidate(testCandidateTwo)
// if err != nil {
// 	t.Errorf("Couldn't add test candidate to database")
// }

// e := New(&testDatabase, false, []string{})

// err = os.Remove("election_testdb.db")
// if err != nil {
// 	t.Log("Cleanup failed")
// }
// }
