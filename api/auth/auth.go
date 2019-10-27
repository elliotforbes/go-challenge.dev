package auth

import (
	"fmt"
	"net/http"
)

func Auth(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Authenticating Endpoint")
		endpoint(w, r)
	})
}
