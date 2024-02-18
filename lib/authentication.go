package authentication

import (
	"github.com/dgrijalva/jwt-go"
)

// Define a struct to represent a user
type SignInDetail struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Define a struct to represent a user
type SignUpDetail struct {
	Username         string `json:"username"`
	Password         string `json:"password"`
	Confirm_Password string `json:"confirm_password"`
	First_Name       string `json:"first_name"`
	Last_Name        string `json:"last_name"`
	Email            string `json:"email"`
}

// Define a struct to represent JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
