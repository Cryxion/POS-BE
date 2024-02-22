package inventory

import (
	"encoding/json"
	"net/http"
	"pos-be/.gen/YAPOS/public/model"
	. "pos-be/.gen/YAPOS/public/table"
	db "pos-be/database"

	. "github.com/go-jet/jet/v2/postgres"
)

// TODO : Get inventories based on shop
func Get(w http.ResponseWriter, r *http.Request) {

	var shopDetail model.Shop
	json.NewDecoder(r.Body).Decode(&shopDetail)

	err := db.InitDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "3000: Please try again!"})
		return
	}

	userInsertion := Inventory.SELECT(Inventory.AllColumns).WHERE(Inventory.ShopID.EQ(Int32(shopDetail.ShopID)))

	// Retrieve the database connection
	database := db.GetDB()
	defer database.Close()

	var dest []struct{ model.Inventory }

	err = userInsertion.Query(database, &dest)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unable to retrieve item for shop!"})
		return
	} else {
		jsonData, err := json.Marshal(dest)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"error": "Unable to retrieve item for shop!"})
		}
		// Print the JSON data
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
