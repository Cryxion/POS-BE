package signin

import (
	"encoding/json"
	"fmt"
	"net/http"
	authentication "pos-be/lib"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define a map to store users (in memory for simplicity)
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Define a secret key for signing JWT tokens
var jwtKey = []byte("your-secret-key")

// Handler for user login
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var signInDetail authentication.SignInDetail
	json.NewDecoder(r.Body).Decode(&signInDetail)

	// Validate user credentials
	password, exists := users[signInDetail.Username]

	if !exists || password != signInDetail.Password {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid username or password")
		return
	}

	// Create JWT token
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &authentication.Claims{
		Username: signInDetail.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error generating token")
		return
	}

	// Return JWT token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
