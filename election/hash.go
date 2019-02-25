package election

import (
	"strings"
)

// return some hash of the vote, must be repeatable
func (v *Vote) hash() {
	// TODO: implement hash

}

func HashVote(rin int, candidateName string) string {
	var str strings.Builder

	for c := range candidateName {
		str.WriteString(string(int(c)))
	}
	return str.String()
}
