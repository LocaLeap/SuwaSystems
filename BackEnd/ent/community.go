// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"suwasystem/backend/ent/community"
	"suwasystem/backend/ent/user"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Community is the model entity for the Community schema.
type Community struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Rate holds the value of the "rate" field.
	Rate float32 `json:"rate,omitempty"`
	// Tax holds the value of the "tax" field.
	Tax float32 `json:"tax,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// OwnerUserID holds the value of the "owner_user_id" field.
	OwnerUserID uuid.UUID `json:"owner_user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommunityQuery when eager-loading is set.
	Edges        CommunityEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CommunityEdges holds the relations/edges for other nodes in the graph.
type CommunityEdges struct {
	// Wallets holds the value of the wallets edge.
	Wallets []*Wallet `json:"wallets,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// FromTransactions holds the value of the from_transactions edge.
	FromTransactions []*Transaction `json:"from_transactions,omitempty"`
	// ToTransactions holds the value of the to_transactions edge.
	ToTransactions []*Transaction `json:"to_transactions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// WalletsOrErr returns the Wallets value or an error if the edge
// was not loaded in eager-loading.
func (e CommunityEdges) WalletsOrErr() ([]*Wallet, error) {
	if e.loadedTypes[0] {
		return e.Wallets, nil
	}
	return nil, &NotLoadedError{edge: "wallets"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommunityEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// FromTransactionsOrErr returns the FromTransactions value or an error if the edge
// was not loaded in eager-loading.
func (e CommunityEdges) FromTransactionsOrErr() ([]*Transaction, error) {
	if e.loadedTypes[2] {
		return e.FromTransactions, nil
	}
	return nil, &NotLoadedError{edge: "from_transactions"}
}

// ToTransactionsOrErr returns the ToTransactions value or an error if the edge
// was not loaded in eager-loading.
func (e CommunityEdges) ToTransactionsOrErr() ([]*Transaction, error) {
	if e.loadedTypes[3] {
		return e.ToTransactions, nil
	}
	return nil, &NotLoadedError{edge: "to_transactions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Community) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case community.FieldRate, community.FieldTax:
			values[i] = new(sql.NullFloat64)
		case community.FieldID:
			values[i] = new(sql.NullInt64)
		case community.FieldName, community.FieldDescription:
			values[i] = new(sql.NullString)
		case community.FieldCreatedAt, community.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case community.FieldOwnerUserID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Community fields.
func (c *Community) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case community.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case community.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case community.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case community.FieldRate:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field rate", values[i])
			} else if value.Valid {
				c.Rate = float32(value.Float64)
			}
		case community.FieldTax:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field tax", values[i])
			} else if value.Valid {
				c.Tax = float32(value.Float64)
			}
		case community.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case community.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case community.FieldOwnerUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field owner_user_id", values[i])
			} else if value != nil {
				c.OwnerUserID = *value
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Community.
// This includes values selected through modifiers, order, etc.
func (c *Community) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryWallets queries the "wallets" edge of the Community entity.
func (c *Community) QueryWallets() *WalletQuery {
	return NewCommunityClient(c.config).QueryWallets(c)
}

// QueryOwner queries the "owner" edge of the Community entity.
func (c *Community) QueryOwner() *UserQuery {
	return NewCommunityClient(c.config).QueryOwner(c)
}

// QueryFromTransactions queries the "from_transactions" edge of the Community entity.
func (c *Community) QueryFromTransactions() *TransactionQuery {
	return NewCommunityClient(c.config).QueryFromTransactions(c)
}

// QueryToTransactions queries the "to_transactions" edge of the Community entity.
func (c *Community) QueryToTransactions() *TransactionQuery {
	return NewCommunityClient(c.config).QueryToTransactions(c)
}

// Update returns a builder for updating this Community.
// Note that you need to call Community.Unwrap() before calling this method if this Community
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Community) Update() *CommunityUpdateOne {
	return NewCommunityClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Community entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Community) Unwrap() *Community {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Community is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Community) String() string {
	var builder strings.Builder
	builder.WriteString("Community(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("rate=")
	builder.WriteString(fmt.Sprintf("%v", c.Rate))
	builder.WriteString(", ")
	builder.WriteString("tax=")
	builder.WriteString(fmt.Sprintf("%v", c.Tax))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("owner_user_id=")
	builder.WriteString(fmt.Sprintf("%v", c.OwnerUserID))
	builder.WriteByte(')')
	return builder.String()
}

// Communities is a parsable slice of Community.
type Communities []*Community
