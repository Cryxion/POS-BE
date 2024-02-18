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

var User = newUserTable("public", "user", "")

type userTable struct {
	postgres.Table

	// Columns
	UserID       postgres.ColumnInteger
	Username     postgres.ColumnString
	Email        postgres.ColumnString
	PasswordHash postgres.ColumnString
	FirstName    postgres.ColumnString
	LastName     postgres.ColumnString
	CreatedAt    postgres.ColumnTimestampz
	UpdatedAt    postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type UserTable struct {
	userTable

	EXCLUDED userTable
}

// AS creates new UserTable with assigned alias
func (a UserTable) AS(alias string) *UserTable {
	return newUserTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UserTable with assigned schema name
func (a UserTable) FromSchema(schemaName string) *UserTable {
	return newUserTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UserTable with assigned table prefix
func (a UserTable) WithPrefix(prefix string) *UserTable {
	return newUserTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UserTable with assigned table suffix
func (a UserTable) WithSuffix(suffix string) *UserTable {
	return newUserTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUserTable(schemaName, tableName, alias string) *UserTable {
	return &UserTable{
		userTable: newUserTableImpl(schemaName, tableName, alias),
		EXCLUDED:  newUserTableImpl("", "excluded", ""),
	}
}

func newUserTableImpl(schemaName, tableName, alias string) userTable {
	var (
		UserIDColumn       = postgres.IntegerColumn("user_id")
		UsernameColumn     = postgres.StringColumn("username")
		EmailColumn        = postgres.StringColumn("email")
		PasswordHashColumn = postgres.StringColumn("password_hash")
		FirstNameColumn    = postgres.StringColumn("first_name")
		LastNameColumn     = postgres.StringColumn("last_name")
		CreatedAtColumn    = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn    = postgres.TimestampzColumn("updated_at")
		allColumns         = postgres.ColumnList{UserIDColumn, UsernameColumn, EmailColumn, PasswordHashColumn, FirstNameColumn, LastNameColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns     = postgres.ColumnList{UsernameColumn, EmailColumn, PasswordHashColumn, FirstNameColumn, LastNameColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return userTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:       UserIDColumn,
		Username:     UsernameColumn,
		Email:        EmailColumn,
		PasswordHash: PasswordHashColumn,
		FirstName:    FirstNameColumn,
		LastName:     LastNameColumn,
		CreatedAt:    CreatedAtColumn,
		UpdatedAt:    UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}