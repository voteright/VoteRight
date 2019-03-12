package primaryapi

import "net/http"

// GetAllVoters gets the list of the voters in the election
func (api *PrimaryAPI) GetAllVoters(w http.ResponseWriter, r *http.Request) {
	// TODO: Needs auth
	// if (!admin) return 403

	v, err := api.Election.GetAllVoters()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	WriteJSON(w, v)

}
