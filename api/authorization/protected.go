package authorization

import (
	"fmt"
	"net/http"
	user "pos-be/lib"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your-secret-key")

// Handler for protected resource
func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	// Extract JWT token from request header
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Authorization header missing")
		return
	}

	// Parse JWT token
	token, err := jwt.ParseWithClaims(tokenString, &user.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid token")
		fmt.Fprintf(w, err.Error())
		return
	}

	// Check if token is valid
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid token")
		return
	}

	// Token is valid, extract claims and respond with protected data
	claims := token.Claims.(*user.Claims)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome, %s!", claims.Username)
}
