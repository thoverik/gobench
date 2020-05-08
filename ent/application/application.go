// Code generated by entc, DO NOT EDIT.

package application

import (
	"time"
)

const (
	// Label holds the string label denoting the application type in the database.
	Label = "application"
	// FieldID holds the string denoting the id field in the database.
	FieldID         = "id"         // FieldName holds the string denoting the name vertex property in the database.
	FieldName       = "name"       // FieldStatus holds the string denoting the status vertex property in the database.
	FieldStatus     = "status"     // FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt  = "created_at" // FieldFinishedAt holds the string denoting the finished_at vertex property in the database.
	FieldFinishedAt = "finished_at"

	// Table holds the table name of the application in the database.
	Table = "applications"
)

// Columns holds all SQL columns for application fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldStatus,
	FieldCreatedAt,
	FieldFinishedAt,
}

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
)