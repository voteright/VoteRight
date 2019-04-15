package api

import "net/http"

// GetAllRaces gets the list of the candidates in the election
func (api *PrimaryAPI) GetAllRaces(w http.ResponseWriter, r *http.Request) {

	v, err := api.Election.GetAllRaces()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	WriteJSON(w, v)

}
