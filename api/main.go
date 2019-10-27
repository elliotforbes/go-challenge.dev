package main

import (
	"fmt"
	"net/http"
	"github.com/elliotforbes/go-challenge.dev/auth"
	"github.com/elliotforbes/go-challenge.dev/execute"
	"github.com/elliotforbes/go-challenge.dev/profile"
	"github.com/elliotforbes/go-challenge.dev/challenge"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.Handle("/execute", auth.Auth(execute.Execute))
	http.Handle("/challenge", auth.Auth(challenge.GetChallenge))
	http.Handle("/profile", auth.Auth(profile.GetProfile))
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}