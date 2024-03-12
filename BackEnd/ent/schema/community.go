package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Community holds the schema definition for the Community entity.
type Community struct {
	ent.Schema
}

// Fields of the Community.
func (Community) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description").
			Optional(),
		field.Float32("rate").
			Default(1.0).
			Positive(),
		field.Float32("tax").
			Default(0.0).
			Range(0, 0.9),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.UUID("owner_user_id", uuid.UUID{}).
			Optional(),
	}
}

// Edges of the Community.
func (Community) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("wallets", Wallet.Type),
		edge.From("owner", User.Type).
			Ref("owner_communities").
			Unique().
			Field("owner_user_id"),
		edge.To("from_transactions", Transaction.Type),
		edge.To("to_transactions", Transaction.Type),
	}
}
