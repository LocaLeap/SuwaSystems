package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Int("amount").
			Positive(),
		field.Int("tax").
			Positive().
			Default(0),
		field.Int("total_amount").
			Positive(),
		field.Time("transaction_date").
			Default(time.Now),
		field.Int("from_community_id").
			Optional(),
		field.Int("to_community_id").
			Optional(),
		field.Int("from_wallet_id").
			Optional(),
		field.Int("to_wallet_id").
			Optional(),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("from_community", Community.Type).
			Ref("from_transactions").
			Unique().
			Field("from_community_id"),
		edge.From("to_community", Community.Type).
			Ref("to_transactions").
			Unique().
			Field("to_community_id"),
		edge.From("from_wallet", Wallet.Type).
			Ref("from_transactions").
			Unique().
			Field("from_wallet_id"),
		edge.From("to_wallet", Wallet.Type).
			Ref("to_transactions").
			Unique().
			Field("to_wallet_id"),
	}
}
