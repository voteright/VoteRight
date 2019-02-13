package primaryapi

import (
	"fmt"
	"net/http"

	"github.com/voteright/voteright/config"
	"github.com/voteright/voteright/election"

	"github.com/go-chi/chi"
)

// PrimaryAPI represents the configuration for the primary vote server api
type PrimaryAPI struct {
	ListenURL string
	Election  *election.Election
	r         chi.Router
}

// IndexHandler serves the main vote page
func (api *PrimaryAPI) IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

// Serve begins the server
func (api *PrimaryAPI) Serve() {
	fmt.Printf("Serving on: %s \n", api.ListenURL)
	if err := http.ListenAndServe(api.ListenURL, api.r); err != nil {
		fmt.Println("Unable to serve.")
	}
}

// New returns a new PrimaryAPI object
func New(cfg *config.Config, e *election.Election) *PrimaryAPI {
	r := chi.NewRouter()

	api := &PrimaryAPI{
		ListenURL: cfg.ListenURL,
		Election:  e,
		r:         r,
	}

	r.Get("/", api.IndexHandler)
	r.Method("GET", "/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("static/assets/"))))
	return api
}
