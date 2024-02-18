package signup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pos-be/.gen/YAPOS/public/model"
	. "pos-be/.gen/YAPOS/public/table"
	db "pos-be/database"
	authentication "pos-be/lib"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Define a map to store users (in memory for simplicity)
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Handler for user registration
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var signUpDetail authentication.SignUpDetail
	json.NewDecoder(r.Body).Decode(&signUpDetail)

	// Store user information (in memory for simplicity)
	// users[loginDetail.Username] = signUpDetail.Password
	if signUpDetail.Password != signUpDetail.Confirm_Password {

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpDetail.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = db.InitDB()
	if err != nil {
		fmt.Println("Error initializing DB:", err)
		return
	}

	newUser := model.User{
		Username:     signUpDetail.Username,
		PasswordHash: string(hashedPassword),
		FirstName:    signUpDetail.First_Name,
		LastName:     signUpDetail.Last_Name,
		Email:        signUpDetail.Email,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	userInsertion := User.INSERT(User.Username, User.PasswordHash, User.FirstName, User.LastName, User.Email, User.CreatedAt, User.UpdatedAt).MODEL(newUser)

	// Retrieve the database connection
	database := db.GetDB()
	defer database.Close()

	test, err := userInsertion.Exec(database)

	fmt.Println(err)
	fmt.Println(test)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User %s registered successfully", signUpDetail.Username)
}
