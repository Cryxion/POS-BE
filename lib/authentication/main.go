package authentication

import (
	"errors"
	"net/http"
	"pos-be/.gen/YAPOS/public/model"
	"time"

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
	UserId    int32  `json:"number"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	jwt.StandardClaims
}

var jwtKey = []byte("your-secret-key")

func ParseJWTToken(r *http.Request) (*Claims, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return nil, errors.New("Unauthorized")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func NewJWTToken(res model.User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := Claims{
		UserId:    int32(res.UserID),
		Username:  res.Username,
		FirstName: res.FirstName,
		Email:     res.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
