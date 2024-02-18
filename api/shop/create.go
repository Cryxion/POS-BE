package shop

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pos-be/.gen/YAPOS/public/model"
	"pos-be/.gen/YAPOS/public/table"
	db "pos-be/database"
	authentication "pos-be/lib"
	"time"
)

func Create(w http.ResponseWriter, r *http.Request) {

	var shopDetail model.Shop
	json.NewDecoder(r.Body).Decode(&shopDetail)

	err := db.InitDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "3000: Please try again!"})
		return
	}

	claim, err := authentication.ParseJWTToken(r)

	newShop := model.Shop{
		UserID:          int32(claim.UserId),
		ShopName:        shopDetail.ShopName,
		ShopDescription: shopDetail.ShopDescription,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	userInsertion := table.Shop.INSERT(table.Shop.UserID, table.Shop.ShopName, table.Shop.ShopDescription, table.Shop.CreatedAt, table.Shop.UpdatedAt).MODEL(newShop)

	// Retrieve the database connection
	database := db.GetDB()
	defer database.Close()

	_, err = userInsertion.Exec(database)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Creation for shop %s failed!", shopDetail.ShopName)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Shop %s successfully created", shopDetail.ShopName)

}
