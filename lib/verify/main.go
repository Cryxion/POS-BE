package verify

import (
	"pos-be/.gen/YAPOS/public/model"
	"pos-be/.gen/YAPOS/public/table"
	db "pos-be/database"

	"github.com/go-jet/jet/v2/postgres"
)

func ShopOwnership(userID int32, shopID int32) bool {
	err := db.InitDB()
	if err != nil {
		return false
	}

	var dest *model.Shop
	getShops := table.Shop.SELECT(table.Shop.AllColumns).WHERE(table.Shop.UserID.EQ(postgres.Int32(userID))).WHERE(table.Shop.ShopID.EQ(postgres.Int32(shopID)))
	database := db.GetDB()
	defer database.Close()

	err = getShops.Query(database, &dest)
	if err != nil {
		return false
	}

	if dest != nil {
		return true
	}
	return false
}
