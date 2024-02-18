package inventory

import (
	"fmt"
	"net/http"
)

// Handler for protected resource
func Get(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome, %s!", claims.Username)
}
