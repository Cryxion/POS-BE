package inventory

import (
	"encoding/json"
	"net/http"
	"pos-be/.gen/YAPOS/public/model"
	"pos-be/.gen/YAPOS/public/table"
	db "pos-be/database"
	"time"

	"pos-be/lib/authentication"
	"pos-be/lib/result"
	"pos-be/lib/verify"

	"github.com/go-jet/jet/v2/postgres"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var inventoryDetail model.Inventory
	json.NewDecoder(r.Body).Decode(&inventoryDetail)

	//Verify DB connection is able to be establish before continuing further
	err := db.InitDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(result.Json_return(false, "3000: Please try again!", nil))
		return
	}

	//Verify user is truly login, though initial checking already done in the middleware
	claim, err := authentication.ParseJWTToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(result.Json_return(false, "Not authorized to perform this action", nil))
		return
	}

	//Verify Shop belongs to this user, to ensure no ill intention or knowledgeable user.
	if !verify.ShopOwnership(claim.UserId, inventoryDetail.ShopID) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(result.Json_return(false, "Not authorized to perform this action", nil))
		return
	}

	//New class for insertion
	newInventory := model.Inventory{
		ShopID:          inventoryDetail.ShopID,
		ItemName:        inventoryDetail.ItemName,
		ItemDescription: inventoryDetail.ItemDescription,
		IsObsolete:      false,
		Quantity:        inventoryDetail.Quantity,
		MinimumQuantity: inventoryDetail.Quantity,
		IsCount:         inventoryDetail.IsCount,
		ActualPrice:     inventoryDetail.ActualPrice,
		DiscountedPrice: inventoryDetail.DiscountedPrice,
		CostPrice:       inventoryDetail.CostPrice,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	inventoryInsertion := table.Inventory.INSERT(table.InventoryHistory.AllColumns).MODEL(newInventory)

	// Retrieve the database connection
	database := db.GetDB()
	defer database.Close()

	_, err = inventoryInsertion.Exec(database)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write(result.Json_return(false, "Failed to create item, please try again!", nil))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(result.Json_return(true, "Item has been created!", nil))

}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var shopDetail model.Shop
	json.NewDecoder(r.Body).Decode(&shopDetail)

	err := db.InitDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(result.Json_return(false, "3000: Please try again!", nil))
		return
	}

	///TODO: verify user with shop ownership before calling the list
	getInventory := table.Inventory.SELECT(table.Inventory.AllColumns).WHERE(table.Inventory.ShopID.EQ(postgres.Int32(shopDetail.ShopID)))

	// Retrieve the database connection
	database := db.GetDB()
	defer database.Close()

	var dest []model.Inventory

	err = getInventory.Query(database, &dest)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write(result.Json_return(false, "Unable to retrieve inventory for shop!", nil))
		return
	} else {
		w.Write(result.Json_return(true, "Inventory retrieved!", dest))
	}
}
