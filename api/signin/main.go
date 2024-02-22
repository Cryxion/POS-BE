package signin

import (
	"encoding/json"
	"net/http"
	"pos-be/.gen/YAPOS/public/model"
	"pos-be/.gen/YAPOS/public/table"
	db "pos-be/database"
	"pos-be/lib/authentication"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-jet/jet/v2/postgres"

	_ "github.com/lib/pq"
)

// Define a secret key for signing JWT tokens
var jwtKey = []byte("your-secret-key")

// Handler for user login
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var signInDetail authentication.SignInDetail
	json.NewDecoder(r.Body).Decode(&signInDetail)

	err := db.InitDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "3000: Please try again!"})
		return
	}

	statement := postgres.SELECT(table.User.AllColumns).FROM(table.User).WHERE(table.User.Username.EQ(postgres.String(signInDetail.Username)))
	database := db.GetDB()

	var res model.User
	err = statement.Query(database, &res)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"error": "Incorrect username or password!"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.PasswordHash), []byte(signInDetail.Password))
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"error": "Incorrect username or password!"})
		return
	}
	// Create JWT token
	tokenString, err := authentication.NewJWTToken(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unable to generate token!"})
		return
	}

	// Return JWT token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
