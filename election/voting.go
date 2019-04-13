package election

import (
	"fmt"

	"github.com/voteright/voteright/models"
)

type idpost struct {
	ID int
}

func (e *Election) castVote(voter *models.Voter, vote *models.Vote) error {
	fmt.Println(voter, vote)
	err := e.db.StoreVote(*vote)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// CastVotes handles casting the votes
func (e *Election) CastVotes(voter *models.Voter, votes []models.Vote) error {
	for _, v := range votes {
		err := e.castVote(voter, &v)
		if err != nil {
			return err
		}
	}
	return e.db.SetVoted(*voter)

}
