package database

import "github.com/voteright/voteright/models"

// StoreCohort stores a cohort in the database
func (d *Database) StoreCohort(cohort models.Cohort) error {
	st, err := d.db.Prepare("INSERT INTO cohorts VALUES (?,?)")
	if err != nil {
		return err
	}
	_, err = st.Exec(cohort.ID, cohort.Name)
	if err != nil {
		return err
	}
	return nil
}

// StoreCohorts stores cohorts in the database
func (d *Database) StoreCohorts(cohorts []models.Cohort) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	for _, cohort := range cohorts {
		st, err := tx.Prepare("INSERT INTO cohorts VALUES (?,?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = st.Exec(cohort.ID, cohort.Name)

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

// GetAllCohorts returns all cohorts in the database
func (d *Database) GetAllCohorts() ([]models.Cohort, error) {
	res, err := d.db.Query("SELECT * from cohorts")
	if err != nil {
		return nil, err
	}
	defer res.Close()
	ret := []models.Cohort{}

	for res.Next() {
		v := &models.Cohort{}
		err := res.Scan(&v.ID, &v.Name)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *v)
	}
	return ret, nil
}
