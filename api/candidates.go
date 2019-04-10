package api

import "net/http"

// GetAllCandidates gets the list of the candidates in the election
func (api *PrimaryAPI) GetAllCandidates(w http.ResponseWriter, r *http.Request) {

	v, err := api.Election.GetAllCandidates()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	WriteJSON(w, v)

}

// GetAllCohorts gets the list of the candidates in the election
func (api *PrimaryAPI) GetAllCohorts(w http.ResponseWriter, r *http.Request) {

	v, err := api.Election.GetAllCohorts()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	WriteJSON(w, v)

}

// GetAllCandidatesWithVotes gets the list of the voters in the election
func (api *PrimaryAPI) GetAllCandidatesWithVotes(w http.ResponseWriter, r *http.Request) {
	// TODO: Needs auth
	// if (!admin) return 403

	v, err := api.Election.GetCandidateVoteCounts()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	WriteJSON(w, v)

}
