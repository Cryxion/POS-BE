//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Inventory struct {
	ItemID          int32 `sql:"primary_key"`
	ShopID          *int32
	ItemName        *string
	ItemDescription *string
	Quantity        *int32
	MinimumQuantity *int32
	IsCount         *bool
	IsObsolete      *bool
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	ActualPrice     *string
	DiscountedPrice *string
	CostPrice       *string
}
