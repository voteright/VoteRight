package api

import (
	"cloud.google.com/go/spanner/admin/database/apiv1"
	"github.com/voteright/voteright/election"
	"testing"
)

func Test_Candidates_Endpoints(t *testing.T) {
	d := database.StormDB{
		File: "testdb.dbtest"
	}
	d.Connect()
	
	e := election.New()
	a := PrimaryAPI{
		ListenURL: "0.0.0.0:5555",
		Election: 
	}

}
