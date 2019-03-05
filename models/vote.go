package models

// Vote represents a single vote
type Vote struct {
	Hash      string
	Candidate int
	StudentID int
}

// calcluate if it does not exist some hash of the vote, must be repeatable
func (v *Vote) hash(voter *Voter) {
	// TODO: implement hash

}
