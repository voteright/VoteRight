package models

// Voter represents a voter in the election
type Voter struct {
	StudentID int `storm:"id"`
	Cohort    int
	Name      string
}
