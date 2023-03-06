package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/hay-kot/homebox/backend/internal/data/ent/schema/mixins"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the Home.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(255).
			NotEmpty(),
		field.Enum("currency").
			Default("usd").
			Values("usd", "eur", "gbp", "jpy", "zar", "aud", "nok", "sek", "dkk", "inr", "rmb", "bgn", "chf", "pln", "try", "ron"),
	}
}

// Edges of the Home.
func (Group) Edges() []ent.Edge {
	owned := func(name string, t any) ent.Edge {
		return edge.To(name, t).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			})
	}

	return []ent.Edge{
		owned("users", User.Type),
		owned("locations", Location.Type),
		owned("items", Item.Type),
		owned("labels", Label.Type),
		owned("documents", Document.Type),
		owned("invitation_tokens", GroupInvitationToken.Type),
		owned("notifiers", Notifier.Type),
	}
}

// GroupMixin when embedded in an ent.Schema, adds a reference to
// the Group entity.
type GroupMixin struct {
	ref string
	mixin.Schema
}

func (g GroupMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).
			Ref(g.ref).
			Unique().
			Required(),
	}
}
