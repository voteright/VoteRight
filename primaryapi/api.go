package primaryapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/voteright/voteright/database"

	"github.com/voteright/voteright/config"
	"github.com/voteright/voteright/election"

	"github.com/go-chi/chi"
)

// PrimaryAPI represents the configuration for the primary vote server api
type PrimaryAPI struct {
	ListenURL string
	Election  *election.Election
	Database  *database.Database
	r         chi.Router
}

// IndexHandler serves the main vote page
func (api *PrimaryAPI) IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

// AdminHandler serves the main Admin page
func (api *PrimaryAPI) AdminHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/admin.html")
}

// Serve begins the server
func (api *PrimaryAPI) Serve() {
	fmt.Printf("Serving on: %s \n", api.ListenURL)
	if err := http.ListenAndServe(api.ListenURL, api.r); err != nil {
		fmt.Println("Unable to serve.")
	}
}

// WriteJSON writes the data as JSON.
func WriteJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Write(b)
	return nil
}

// New returns a new PrimaryAPI object
func New(cfg *config.Config, e *election.Election, d *database.Database) *PrimaryAPI {
	r := chi.NewRouter()

	api := &PrimaryAPI{
		ListenURL: cfg.ListenURL,
		Election:  e,
		Database:  d,
		r:         r,
	}

	r.Get("/", api.IndexHandler)
	// we're going to need to add mock auth here at some point
	r.Get("/admin", api.AdminHandler)
	r.Get("/cohorts", api.GetAllCohorts)
	r.Route("/voters", func(r chi.Router) {
		r.Get("/", api.GetAllVoters)
		r.Post("/validate", api.ValidateVoter)
		r.Post("/login", api.LoginVoter)
	})
	r.Route("/candidates", func(r chi.Router) {
		r.Get("/", api.GetAllCandidates)
		r.Get("/votes", api.GetAllCandidatesWithVotes)
	})

	r.Method("GET", "/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("static/assets/"))))
	return api
}
