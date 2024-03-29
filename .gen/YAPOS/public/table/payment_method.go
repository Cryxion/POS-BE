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

var PaymentMethod = newPaymentMethodTable("public", "payment_method", "")

type paymentMethodTable struct {
	postgres.Table

	// Columns
	PaymentMethodID postgres.ColumnInteger
	PaymentName     postgres.ColumnString
	PaymentCurrency postgres.ColumnString
	IsAvailable     postgres.ColumnBool
	CreatedAt       postgres.ColumnTimez
	UpdatedAt       postgres.ColumnTimez

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PaymentMethodTable struct {
	paymentMethodTable

	EXCLUDED paymentMethodTable
}

// AS creates new PaymentMethodTable with assigned alias
func (a PaymentMethodTable) AS(alias string) *PaymentMethodTable {
	return newPaymentMethodTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PaymentMethodTable with assigned schema name
func (a PaymentMethodTable) FromSchema(schemaName string) *PaymentMethodTable {
	return newPaymentMethodTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PaymentMethodTable with assigned table prefix
func (a PaymentMethodTable) WithPrefix(prefix string) *PaymentMethodTable {
	return newPaymentMethodTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PaymentMethodTable with assigned table suffix
func (a PaymentMethodTable) WithSuffix(suffix string) *PaymentMethodTable {
	return newPaymentMethodTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPaymentMethodTable(schemaName, tableName, alias string) *PaymentMethodTable {
	return &PaymentMethodTable{
		paymentMethodTable: newPaymentMethodTableImpl(schemaName, tableName, alias),
		EXCLUDED:           newPaymentMethodTableImpl("", "excluded", ""),
	}
}

func newPaymentMethodTableImpl(schemaName, tableName, alias string) paymentMethodTable {
	var (
		PaymentMethodIDColumn = postgres.IntegerColumn("payment_method_id")
		PaymentNameColumn     = postgres.StringColumn("payment_name")
		PaymentCurrencyColumn = postgres.StringColumn("payment_currency")
		IsAvailableColumn     = postgres.BoolColumn("isAvailable")
		CreatedAtColumn       = postgres.TimezColumn("created_at")
		UpdatedAtColumn       = postgres.TimezColumn("updated_at")
		allColumns            = postgres.ColumnList{PaymentMethodIDColumn, PaymentNameColumn, PaymentCurrencyColumn, IsAvailableColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns        = postgres.ColumnList{PaymentNameColumn, PaymentCurrencyColumn, IsAvailableColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return paymentMethodTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		PaymentMethodID: PaymentMethodIDColumn,
		PaymentName:     PaymentNameColumn,
		PaymentCurrency: PaymentCurrencyColumn,
		IsAvailable:     IsAvailableColumn,
		CreatedAt:       CreatedAtColumn,
		UpdatedAt:       UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
