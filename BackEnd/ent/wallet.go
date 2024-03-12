// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"suwasystem/backend/ent/community"
	"suwasystem/backend/ent/user"
	"suwasystem/backend/ent/wallet"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Wallet is the model entity for the Wallet schema.
type Wallet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Balance holds the value of the "balance" field.
	Balance int `json:"balance,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// CommunityID holds the value of the "community_id" field.
	CommunityID int `json:"community_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WalletQuery when eager-loading is set.
	Edges        WalletEdges `json:"edges"`
	selectValues sql.SelectValues
}

// WalletEdges holds the relations/edges for other nodes in the graph.
type WalletEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Community holds the value of the community edge.
	Community *Community `json:"community,omitempty"`
	// FromTransactions holds the value of the from_transactions edge.
	FromTransactions []*Transaction `json:"from_transactions,omitempty"`
	// ToTransactions holds the value of the to_transactions edge.
	ToTransactions []*Transaction `json:"to_transactions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WalletEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// CommunityOrErr returns the Community value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WalletEdges) CommunityOrErr() (*Community, error) {
	if e.Community != nil {
		return e.Community, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: community.Label}
	}
	return nil, &NotLoadedError{edge: "community"}
}

// FromTransactionsOrErr returns the FromTransactions value or an error if the edge
// was not loaded in eager-loading.
func (e WalletEdges) FromTransactionsOrErr() ([]*Transaction, error) {
	if e.loadedTypes[2] {
		return e.FromTransactions, nil
	}
	return nil, &NotLoadedError{edge: "from_transactions"}
}

// ToTransactionsOrErr returns the ToTransactions value or an error if the edge
// was not loaded in eager-loading.
func (e WalletEdges) ToTransactionsOrErr() ([]*Transaction, error) {
	if e.loadedTypes[3] {
		return e.ToTransactions, nil
	}
	return nil, &NotLoadedError{edge: "to_transactions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Wallet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case wallet.FieldID, wallet.FieldBalance, wallet.FieldCommunityID:
			values[i] = new(sql.NullInt64)
		case wallet.FieldCreatedAt, wallet.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case wallet.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Wallet fields.
func (w *Wallet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wallet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case wallet.FieldBalance:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field balance", values[i])
			} else if value.Valid {
				w.Balance = int(value.Int64)
			}
		case wallet.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				w.CreatedAt = value.Time
			}
		case wallet.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				w.UpdatedAt = value.Time
			}
		case wallet.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				w.UserID = *value
			}
		case wallet.FieldCommunityID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field community_id", values[i])
			} else if value.Valid {
				w.CommunityID = int(value.Int64)
			}
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Wallet.
// This includes values selected through modifiers, order, etc.
func (w *Wallet) Value(name string) (ent.Value, error) {
	return w.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Wallet entity.
func (w *Wallet) QueryUser() *UserQuery {
	return NewWalletClient(w.config).QueryUser(w)
}

// QueryCommunity queries the "community" edge of the Wallet entity.
func (w *Wallet) QueryCommunity() *CommunityQuery {
	return NewWalletClient(w.config).QueryCommunity(w)
}

// QueryFromTransactions queries the "from_transactions" edge of the Wallet entity.
func (w *Wallet) QueryFromTransactions() *TransactionQuery {
	return NewWalletClient(w.config).QueryFromTransactions(w)
}

// QueryToTransactions queries the "to_transactions" edge of the Wallet entity.
func (w *Wallet) QueryToTransactions() *TransactionQuery {
	return NewWalletClient(w.config).QueryToTransactions(w)
}

// Update returns a builder for updating this Wallet.
// Note that you need to call Wallet.Unwrap() before calling this method if this Wallet
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Wallet) Update() *WalletUpdateOne {
	return NewWalletClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Wallet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Wallet) Unwrap() *Wallet {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Wallet is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Wallet) String() string {
	var builder strings.Builder
	builder.WriteString("Wallet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("balance=")
	builder.WriteString(fmt.Sprintf("%v", w.Balance))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(w.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(w.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", w.UserID))
	builder.WriteString(", ")
	builder.WriteString("community_id=")
	builder.WriteString(fmt.Sprintf("%v", w.CommunityID))
	builder.WriteByte(')')
	return builder.String()
}

// Wallets is a parsable slice of Wallet.
type Wallets []*Wallet
