package main

import (
	"fmt"
	"net/http"

	"pos-be/api/authorization"
	signin "pos-be/api/signin"
	"pos-be/api/signup"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/auth/signup", signup.SignUpHandler).Methods("POST")
	r.HandleFunc("/auth/signin", signin.SignInHandler).Methods("POST")
	r.HandleFunc("/protected", authorization.AuthorizeHandler).Methods("GET")

	handler := cors.Default().Handler(r)

	// Start server
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", handler)
}
