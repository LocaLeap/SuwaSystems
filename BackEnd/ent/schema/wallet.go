package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Wallet holds the schema definition for the Wallet entity.
type Wallet struct {
	ent.Schema
}

// Fields of the Wallet.
func (Wallet) Fields() []ent.Field {
	return []ent.Field{
		field.Int("balance").
			Default(0).
			Positive(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.UUID("user_id", uuid.UUID{}).
			Optional(),
		field.Int("community_id").
			Optional(),
	}
}

// Edges of the Wallet.
func (Wallet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("wallets").
			Unique().
			Field("user_id"),
		edge.From("community", Community.Type).
			Ref("wallets").
			Unique().
			Field("community_id"),
		edge.To("from_transactions", Transaction.Type),
		edge.To("to_transactions", Transaction.Type),
	}
}
