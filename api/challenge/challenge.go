package challenge

import (
	"fmt"
	"net/http"
)

func GetChallenge(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Challenge 1")
}
