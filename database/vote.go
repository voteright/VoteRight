package database

import "github.com/voteright/voteright/models"

// StoreVote stores a vote in the database
func (d *Database) StoreVote(vote models.Vote) error {
	st, err := d.db.Prepare("INSERT INTO votes VALUES (?,?)")
	if err != nil {
		return err
	}
	_, err = st.Exec(vote.Hash, vote.Candidate)
	if err != nil {
		return err
	}
	return nil
}

// GetAllVotes returns all votes in the database
func (d *Database) GetAllVotes() ([]models.Vote, error) {
	res, err := d.db.Query("SELECT * from votes")
	if err != nil {
		return nil, err
	}
	defer res.Close()
	ret := []models.Vote{}

	for res.Next() {
		v := &models.Vote{}
		err := res.Scan(&v.Hash, &v.Candidate)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *v)
	}
	return ret, nil
}
