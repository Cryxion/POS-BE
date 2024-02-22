package main

import (
	"fmt"
	"net/http"

	"pos-be/api/inventory"
	"pos-be/api/signin"
	"pos-be/api/signup"

	"pos-be/lib/authentication"
	"pos-be/lib/result"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/auth/signup", signup.SignUpHandler).Methods("POST")
	r.HandleFunc("/auth/signin", signin.SignInHandler).Methods("POST")

	secure := r.PathPrefix("/api/v1").Subrouter()
	secure.Use(JwtVerify)
	secure.HandleFunc("/inventory", inventory.Get).Methods("GET")

	handler := cors.Default().Handler(r)

	// Start server
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", handler)
}

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Your JWT verification logic here
		// For simplicity, let's assume it's just printing for demonstration

		// Parse JWT token
		token, err := authentication.ParseJWTToken(r)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(result.Json_return(false, "Invalid token!", nil))
			return
		}

		// Check if token is valid
		if token == nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(result.Json_return(false, "Invalid token!", nil))
			return
		}

		// Verify JWT token validity here...
		// If valid, proceed to next handler
		// Otherwise, return an error response
		next.ServeHTTP(w, r)
	})
}
