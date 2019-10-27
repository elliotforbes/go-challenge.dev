package execute

import (
	"fmt"
	"net/http"
)

func Execute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing Go Code")
	fmt.Fprintf(w, "Successfully Executed Go Code")
}