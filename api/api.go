package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/voteright/voteright/database"

	"github.com/voteright/voteright/config"
	"github.com/voteright/voteright/election"

	"github.com/go-chi/chi"
)

// PrimaryAPI represents the configuration for the primary vote server api
type PrimaryAPI struct {
	ListenURL          string
	Election           *election.Election
	Database           *database.StormDB
	VerificationAPIKey string // Allow the applciation to authenticate with the verification api
	r                  chi.Router
}

// IndexHandler serves the main vote page
func (api *PrimaryAPI) IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

// ThanksHandler serves the main thanks page
func (api *PrimaryAPI) ThanksHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/thanksforvoting.html")
}

// VoteBoothHandler serves the main vote page
func (api *PrimaryAPI) VoteBoothHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/vote.html")
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

// GetintegrityViolations gets the integrity violations that occured
func (api *PrimaryAPI) GetintegrityViolations(w http.ResponseWriter, r *http.Request) {

	v, err := api.Database.GetAllIntegrityViolations()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	WriteJSON(w, v)

}

// Whoamitestpage is a Proof of concept to get session token
func (api *PrimaryAPI) Whoamitestpage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(403)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(c.Value)
	me, _ := api.Election.GetVoterByID(id)
	WriteJSON(w, me)

}

// New returns a new PrimaryAPI object
func New(cfg *config.Config, e *election.Election, d *database.StormDB) *PrimaryAPI {
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

	r.Route("/voters", func(r chi.Router) {
		r.Get("/", api.GetAllVoters)
		r.Post("/validate", api.ValidateVoter)
		r.Post("/verifyself", api.VerifySelf)
		r.Post("/login", api.LoginVoter)
		r.Get("/whoami", api.Whoamitestpage)
		r.Post("/vote", api.CastVote)
	})

	// if !api.Election.Verification {
	r.Get("/votingbooth", api.VoteBoothHandler)
	r.Get("/cohorts", api.GetAllCohorts)
	r.Get("/thanks", api.ThanksHandler)

	r.Route("/candidates", func(r chi.Router) {
		r.Get("/", api.GetAllCandidates)
		r.Get("/votes", api.GetAllCandidatesWithVotes)
	})

	r.Route("/races", func(r chi.Router) {
		r.Get("/", api.GetAllRaces)
	})
	// }

	r.Route("/integrity", func(r chi.Router) {
		r.Get("/", api.GetintegrityViolations)
	})

	r.Method("GET", "/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("static/assets/"))))
	return api
}
