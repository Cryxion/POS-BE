//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var TransactionDetail = newTransactionDetailTable("public", "transaction_detail", "")

type transactionDetailTable struct {
	postgres.Table

	// Columns
	TransactionDetailID postgres.ColumnInteger
	TransactionID       postgres.ColumnInteger
	ShopID              postgres.ColumnInteger
	ItemID              postgres.ColumnInteger
	Quantity            postgres.ColumnInteger
	PerPrice            postgres.ColumnString
	TotalPrice          postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TransactionDetailTable struct {
	transactionDetailTable

	EXCLUDED transactionDetailTable
}

// AS creates new TransactionDetailTable with assigned alias
func (a TransactionDetailTable) AS(alias string) *TransactionDetailTable {
	return newTransactionDetailTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TransactionDetailTable with assigned schema name
func (a TransactionDetailTable) FromSchema(schemaName string) *TransactionDetailTable {
	return newTransactionDetailTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TransactionDetailTable with assigned table prefix
func (a TransactionDetailTable) WithPrefix(prefix string) *TransactionDetailTable {
	return newTransactionDetailTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TransactionDetailTable with assigned table suffix
func (a TransactionDetailTable) WithSuffix(suffix string) *TransactionDetailTable {
	return newTransactionDetailTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTransactionDetailTable(schemaName, tableName, alias string) *TransactionDetailTable {
	return &TransactionDetailTable{
		transactionDetailTable: newTransactionDetailTableImpl(schemaName, tableName, alias),
		EXCLUDED:               newTransactionDetailTableImpl("", "excluded", ""),
	}
}

func newTransactionDetailTableImpl(schemaName, tableName, alias string) transactionDetailTable {
	var (
		TransactionDetailIDColumn = postgres.IntegerColumn("transaction_detail_id")
		TransactionIDColumn       = postgres.IntegerColumn("transaction_id")
		ShopIDColumn              = postgres.IntegerColumn("shop_id")
		ItemIDColumn              = postgres.IntegerColumn("item_id")
		QuantityColumn            = postgres.IntegerColumn("quantity")
		PerPriceColumn            = postgres.StringColumn("per_price")
		TotalPriceColumn          = postgres.StringColumn("total_price")
		allColumns                = postgres.ColumnList{TransactionDetailIDColumn, TransactionIDColumn, ShopIDColumn, ItemIDColumn, QuantityColumn, PerPriceColumn, TotalPriceColumn}
		mutableColumns            = postgres.ColumnList{TransactionIDColumn, ShopIDColumn, ItemIDColumn, QuantityColumn, PerPriceColumn, TotalPriceColumn}
	)

	return transactionDetailTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		TransactionDetailID: TransactionDetailIDColumn,
		TransactionID:       TransactionIDColumn,
		ShopID:              ShopIDColumn,
		ItemID:              ItemIDColumn,
		Quantity:            QuantityColumn,
		PerPrice:            PerPriceColumn,
		TotalPrice:          TotalPriceColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}