// Code generated by ent, DO NOT EDIT.

package transaction

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the transaction type in the database.
	Label = "transaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldTax holds the string denoting the tax field in the database.
	FieldTax = "tax"
	// FieldTotalAmount holds the string denoting the total_amount field in the database.
	FieldTotalAmount = "total_amount"
	// FieldTransactionDate holds the string denoting the transaction_date field in the database.
	FieldTransactionDate = "transaction_date"
	// FieldFromCommunityID holds the string denoting the from_community_id field in the database.
	FieldFromCommunityID = "from_community_id"
	// FieldToCommunityID holds the string denoting the to_community_id field in the database.
	FieldToCommunityID = "to_community_id"
	// FieldFromWalletID holds the string denoting the from_wallet_id field in the database.
	FieldFromWalletID = "from_wallet_id"
	// FieldToWalletID holds the string denoting the to_wallet_id field in the database.
	FieldToWalletID = "to_wallet_id"
	// EdgeFromCommunity holds the string denoting the from_community edge name in mutations.
	EdgeFromCommunity = "from_community"
	// EdgeToCommunity holds the string denoting the to_community edge name in mutations.
	EdgeToCommunity = "to_community"
	// EdgeFromWallet holds the string denoting the from_wallet edge name in mutations.
	EdgeFromWallet = "from_wallet"
	// EdgeToWallet holds the string denoting the to_wallet edge name in mutations.
	EdgeToWallet = "to_wallet"
	// Table holds the table name of the transaction in the database.
	Table = "transactions"
	// FromCommunityTable is the table that holds the from_community relation/edge.
	FromCommunityTable = "transactions"
	// FromCommunityInverseTable is the table name for the Community entity.
	// It exists in this package in order to avoid circular dependency with the "community" package.
	FromCommunityInverseTable = "communities"
	// FromCommunityColumn is the table column denoting the from_community relation/edge.
	FromCommunityColumn = "from_community_id"
	// ToCommunityTable is the table that holds the to_community relation/edge.
	ToCommunityTable = "transactions"
	// ToCommunityInverseTable is the table name for the Community entity.
	// It exists in this package in order to avoid circular dependency with the "community" package.
	ToCommunityInverseTable = "communities"
	// ToCommunityColumn is the table column denoting the to_community relation/edge.
	ToCommunityColumn = "to_community_id"
	// FromWalletTable is the table that holds the from_wallet relation/edge.
	FromWalletTable = "transactions"
	// FromWalletInverseTable is the table name for the Wallet entity.
	// It exists in this package in order to avoid circular dependency with the "wallet" package.
	FromWalletInverseTable = "wallets"
	// FromWalletColumn is the table column denoting the from_wallet relation/edge.
	FromWalletColumn = "from_wallet_id"
	// ToWalletTable is the table that holds the to_wallet relation/edge.
	ToWalletTable = "transactions"
	// ToWalletInverseTable is the table name for the Wallet entity.
	// It exists in this package in order to avoid circular dependency with the "wallet" package.
	ToWalletInverseTable = "wallets"
	// ToWalletColumn is the table column denoting the to_wallet relation/edge.
	ToWalletColumn = "to_wallet_id"
)

// Columns holds all SQL columns for transaction fields.
var Columns = []string{
	FieldID,
	FieldAmount,
	FieldTax,
	FieldTotalAmount,
	FieldTransactionDate,
	FieldFromCommunityID,
	FieldToCommunityID,
	FieldFromWalletID,
	FieldToWalletID,
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
	// AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	AmountValidator func(int) error
	// DefaultTax holds the default value on creation for the "tax" field.
	DefaultTax int
	// TaxValidator is a validator for the "tax" field. It is called by the builders before save.
	TaxValidator func(int) error
	// TotalAmountValidator is a validator for the "total_amount" field. It is called by the builders before save.
	TotalAmountValidator func(int) error
	// DefaultTransactionDate holds the default value on creation for the "transaction_date" field.
	DefaultTransactionDate func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Transaction queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByTax orders the results by the tax field.
func ByTax(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTax, opts...).ToFunc()
}

// ByTotalAmount orders the results by the total_amount field.
func ByTotalAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalAmount, opts...).ToFunc()
}

// ByTransactionDate orders the results by the transaction_date field.
func ByTransactionDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransactionDate, opts...).ToFunc()
}

// ByFromCommunityID orders the results by the from_community_id field.
func ByFromCommunityID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFromCommunityID, opts...).ToFunc()
}

// ByToCommunityID orders the results by the to_community_id field.
func ByToCommunityID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToCommunityID, opts...).ToFunc()
}

// ByFromWalletID orders the results by the from_wallet_id field.
func ByFromWalletID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFromWalletID, opts...).ToFunc()
}

// ByToWalletID orders the results by the to_wallet_id field.
func ByToWalletID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToWalletID, opts...).ToFunc()
}

// ByFromCommunityField orders the results by from_community field.
func ByFromCommunityField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFromCommunityStep(), sql.OrderByField(field, opts...))
	}
}

// ByToCommunityField orders the results by to_community field.
func ByToCommunityField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newToCommunityStep(), sql.OrderByField(field, opts...))
	}
}

// ByFromWalletField orders the results by from_wallet field.
func ByFromWalletField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFromWalletStep(), sql.OrderByField(field, opts...))
	}
}

// ByToWalletField orders the results by to_wallet field.
func ByToWalletField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newToWalletStep(), sql.OrderByField(field, opts...))
	}
}
func newFromCommunityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FromCommunityInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, FromCommunityTable, FromCommunityColumn),
	)
}
func newToCommunityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ToCommunityInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ToCommunityTable, ToCommunityColumn),
	)
}
func newFromWalletStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FromWalletInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, FromWalletTable, FromWalletColumn),
	)
}
func newToWalletStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ToWalletInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ToWalletTable, ToWalletColumn),
	)
}
