package election

import (
	"fmt"

	"github.com/voteright/voteright/models"
)

// CastVote handles casting the vote and sending it to verification servers
func (e *Election) CastVote(voter *models.Voter, vote *models.Vote) error {
	fmt.Println(voter, vote)
	err := e.db.StoreVote(*vote)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}
