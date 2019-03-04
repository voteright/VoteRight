package database

import "github.com/voteright/voteright/models"

// StoreCandidate stores a candidate in the database
func (d *Database) StoreCandidate(candidate models.Candidate) error {
	st, err := d.db.Prepare("INSERT INTO candidates VALUES (?,?,?)")
	if err != nil {
		return err
	}
	_, err = st.Exec(candidate.ID, candidate.Name, candidate.Cohort)
	if err != nil {
		return err
	}
	return nil
}

// StoreCandidates stores candidates in the database
func (d *Database) StoreCandidates(candidates []models.Candidate) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	for _, candidate := range candidates {
		st, err := d.db.Prepare("INSERT INTO candidates VALUES (?,?,?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = st.Exec(candidate.ID, candidate.Name, candidate.Cohort)

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

// GetAllCandidates returns all candidates in the database
func (d *Database) GetAllCandidates() ([]models.Candidate, error) {
	res, err := d.db.Query("SELECT * from candidates")
	if err != nil {
		return nil, err
	}
	defer res.Close()
	ret := []models.Candidate{}

	for res.Next() {
		v := &models.Candidate{}
		err := res.Scan(&v.ID, &v.Name, &v.Cohort)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *v)
	}
	return ret, nil
}
