package profile

import (
	"fmt"
	"net/http"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My Profile")
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Register for a profile")
}

func CompleteChallenge(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Completed Challenge 1")
}

func CompletedChallenges(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of completed challenges")
}
