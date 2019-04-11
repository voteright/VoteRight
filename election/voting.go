package election

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"time"

	"github.com/voteright/voteright/models"
)

type idpost struct {
	ID int
}

// CastVote handles casting the vote and sending it to verification servers
func (e *Election) CastVote(voter *models.Voter, vote *models.Vote) error {
	fmt.Println(voter, vote)
	err := e.db.StoreVote(*vote)
	if err != nil {
		fmt.Println(err)
		return err
	}
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	client := &http.Client{
		Jar: cookieJar,
	}
	for _, str := range e.VerificationServers {
		c := http.Cookie{
			Name:   "session_token",
			Value:  strconv.Itoa(voter.StudentID),
			Domain: str,
		}
		url, err := url.Parse(str)
		if err != nil {
			fmt.Println("err", err.Error())
		}
		client.Jar.SetCookies(url, []*http.Cookie{&c})
		fmt.Println("Performing a post request on ", str)
		b, _ := json.Marshal(&[]idpost{{
			ID: vote.ID,
		}})
		req, err := http.NewRequest("POST", str+"/voters/vote", bytes.NewReader(b))
		if err != nil {
			e.db.StoreIntegrityViolation(models.IntegrityViolation{
				Message: "Failed to post to verification server",
				Time:    time.Now(),
			})
			fmt.Println("err", err.Error())
		}
		_, err = client.Do(req)
		if err != nil {
			e.db.StoreIntegrityViolation(models.IntegrityViolation{
				Message: "Failed to post to verification server",
				Time:    time.Now(),
			})
			fmt.Println("err", err.Error())
		}
	}
	return err
}
