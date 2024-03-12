// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeWallets holds the string denoting the wallets edge name in mutations.
	EdgeWallets = "wallets"
	// EdgeOwnerCommunities holds the string denoting the owner_communities edge name in mutations.
	EdgeOwnerCommunities = "owner_communities"
	// Table holds the table name of the user in the database.
	Table = "users"
	// WalletsTable is the table that holds the wallets relation/edge.
	WalletsTable = "wallets"
	// WalletsInverseTable is the table name for the Wallet entity.
	// It exists in this package in order to avoid circular dependency with the "wallet" package.
	WalletsInverseTable = "wallets"
	// WalletsColumn is the table column denoting the wallets relation/edge.
	WalletsColumn = "user_id"
	// OwnerCommunitiesTable is the table that holds the owner_communities relation/edge.
	OwnerCommunitiesTable = "communities"
	// OwnerCommunitiesInverseTable is the table name for the Community entity.
	// It exists in this package in order to avoid circular dependency with the "community" package.
	OwnerCommunitiesInverseTable = "communities"
	// OwnerCommunitiesColumn is the table column denoting the owner_communities relation/edge.
	OwnerCommunitiesColumn = "owner_user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldEmail,
	FieldPasswordHash,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeCard   Type = "CARD"
	TypeMobile Type = "MOBILE"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeCard, TypeMobile:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPasswordHash orders the results by the password_hash field.
func ByPasswordHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswordHash, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByWalletsCount orders the results by wallets count.
func ByWalletsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWalletsStep(), opts...)
	}
}

// ByWallets orders the results by wallets terms.
func ByWallets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWalletsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOwnerCommunitiesCount orders the results by owner_communities count.
func ByOwnerCommunitiesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOwnerCommunitiesStep(), opts...)
	}
}

// ByOwnerCommunities orders the results by owner_communities terms.
func ByOwnerCommunities(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerCommunitiesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newWalletsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WalletsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WalletsTable, WalletsColumn),
	)
}
func newOwnerCommunitiesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerCommunitiesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OwnerCommunitiesTable, OwnerCommunitiesColumn),
	)
}
