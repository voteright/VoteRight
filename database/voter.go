package database

import "github.com/voteright/voteright/models"

// StoreVoter stores a voter in the database
func (d *Database) StoreVoter(voter models.Voter) error {
	st, err := d.db.Prepare("INSERT INTO voters VALUES (?,?,?)")
	if err != nil {
		return err
	}
	_, err = st.Exec(voter.StudentID, voter.Name, voter.Cohort)
	if err != nil {
		return err
	}
	return nil
}

// StoreVoters stores voters in the database
func (d *Database) StoreVoters(voters []models.Voter) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	for _, voter := range voters {
		st, err := tx.Prepare("INSERT INTO voters VALUES (?,?,?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = st.Exec(voter.StudentID, voter.Name, voter.Cohort)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}

// GetAllVoters returns all voters in the database
func (d *Database) GetAllVoters() ([]models.Voter, error) {
	res, err := d.db.Query("SELECT * from voters")
	if err != nil {
		return nil, err
	}
	defer res.Close()
	ret := []models.Voter{}

	for res.Next() {
		v := &models.Voter{}
		err := res.Scan(&v.StudentID, &v.Name, &v.Cohort)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *v)
	}
	return ret, nil
}
