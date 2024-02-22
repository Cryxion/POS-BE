package shop

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pos-be/.gen/YAPOS/public/model"
	"pos-be/.gen/YAPOS/public/table"
	db "pos-be/database"
	"pos-be/lib/authentication"
	"pos-be/lib/result"
	"time"

	"github.com/go-jet/jet/v2/postgres"
)

func Create(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var shopDetail model.Shop
	json.NewDecoder(r.Body).Decode(&shopDetail)

	err := db.InitDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(result.Json_return(false, "3000: Please try again!", nil))
		return
	}

	claim, err := authentication.ParseJWTToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(result.Json_return(false, "Not authorized to perform this action", nil))
		return
	}

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
		w.Write(result.Json_return(false, fmt.Sprintf("Creation for shop %s failed!", shopDetail.ShopName), nil))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(result.Json_return(true, fmt.Sprintf("Shop %s successfully created", shopDetail.ShopName), nil))

}

func Get(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	err := db.InitDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(result.Json_return(false, "3000: Please try again!", nil))
		return
	}

	claim, err := authentication.ParseJWTToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(result.Json_return(false, "Not authorized to perform this action", nil))
		return
	}

	var dest []model.Shop

	getShops := table.Shop.SELECT(table.Shop.AllColumns).WHERE(table.Shop.UserID.EQ(postgres.Int32(claim.UserId)))

	// Retrieve the database connection
	database := db.GetDB()
	defer database.Close()

	err = getShops.Query(database, &dest)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write(result.Json_return(false, "Unable to retrieve shop(s)!", nil))
		return
	} else {
		// Print the JSON data
		w.Write(result.Json_return(true, "Inventory retrieved!", dest))
	}

}
