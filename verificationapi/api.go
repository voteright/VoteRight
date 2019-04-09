package verificationapi

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/voteright/voteright/config"
	"github.com/voteright/voteright/database"
	"github.com/voteright/voteright/election"
)

// VerificationAPI represents an instance of the verification api server
type VerificationAPI struct {
	APIKey    string
	ListenURL string
	Election  *election.Election
	Database  *database.Database
	r         chi.Router
}

// IndexHandler handles serving the index page for the verification server
func (api *VerificationAPI) IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}

// New returns a new PrimaryAPI object
func New(cfg *config.Config, e *election.Election, d *database.Database) *VerificationAPI {
	r := chi.NewRouter()

	api := &VerificationAPI{
		ListenURL: cfg.ListenURL,
		Election:  e,
		Database:  d,
		r:         r,
	}

	r.Get("/", api.IndexHandler)

	r.Method("GET", "/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("static/assets/"))))
	return api
}

// Serve begins the server
func (api *VerificationAPI) Serve() {
	fmt.Printf("Serving on: %s \n", api.ListenURL)
	if err := http.ListenAndServe(api.ListenURL, api.r); err != nil {
		fmt.Println("Unable to serve.")
	}
}
