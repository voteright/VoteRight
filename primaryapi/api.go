package primaryapi

import (
	"fmt"
	"net/http"
	"strconv"

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

func (api *PrimaryAPI) PostHandler(w http.ResponseWriter, r *http.Request) {
	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	rin := r.FormValue("rin")
	rinInt, err := strconv.Atoi(rin)
	if err != nil {
		fmt.Println(err)
	}
	candidateName := r.FormValue("candidate")
	fmt.Fprintf(w, "Your name: %s %s\n", firstName, lastName)
	fmt.Fprintf(w, "Your RIN: %d\n", rinInt)
	fmt.Fprintf(w, "You voted for: %s\n", candidateName)

	fmt.Println("Your vote's hash: %s\n", election.HashVote(rinInt, candidateName))
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
	r.Post("/", api.PostHandler)
	r.Method("GET", "/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("static/assets/"))))
	return api
}
