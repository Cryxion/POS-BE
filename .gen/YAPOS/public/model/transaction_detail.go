//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type TransactionDetail struct {
	TransactionDetailID int32 `sql:"primary_key"`
	TransactionID       *int32
	ShopID              *int32
	ItemID              *int32
	Quantity            *int32
	PerPrice            *string
	TotalPrice          *string
}