package schema

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/kcarretto/realm/tavern/internal/namegen"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Beacon holds the schema definition for the Beacon entity.
type Beacon struct {
	ent.Schema
}

// Fields of the Beacon.
func (Beacon) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique().
			Immutable().
			DefaultFunc(namegen.GetRandomName).
			Comment("A human readable identifier for the beacon."),
		field.String("principal").
			Optional().
			NotEmpty().
			Annotations(
				entgql.Skip(entgql.SkipMutationUpdateInput),
			).
			Comment("The identity the beacon is authenticated as (e.g. 'root')"),
		field.String("identifier").
			DefaultFunc(newRandomIdentifier).
			NotEmpty().
			Unique().
			Annotations(
				entgql.Skip(entgql.SkipMutationUpdateInput),
			).
			Comment("Unique identifier for the beacon. Unique to each instance of the beacon."),
		field.String("agent_identifier").
			Optional().
			NotEmpty().
			Annotations(
				entgql.Skip(entgql.SkipMutationUpdateInput),
			).
			Comment("Identifies the agent that the beacon is running as (e.g. 'imix')."),
		field.Time("last_seen_at").
			Optional().
			Annotations(
				entgql.OrderField("LAST_SEEN_AT"),
				entgql.Skip(entgql.SkipMutationUpdateInput),
			).
			Comment("Timestamp of when a task was last claimed or updated for the beacon."),
	}
}

// Edges of the Beacon.
func (Beacon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("host", Host.Type).
			Required().
			Unique().
			Comment("Host this beacon is running on."),
		edge.From("tasks", Task.Type).
			Annotations(
				entgql.Skip(entgql.SkipMutationUpdateInput),
			).
			Ref("beacon").
			Comment("Tasks that have been assigned to the beacon."),
	}
}

// Annotations describes additional information for the ent.
func (Beacon) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(
			entgql.MutationUpdate(),
		),
	}
}

func newRandomIdentifier() string {
	buf := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(fmt.Errorf("failed to generate random identifier: %w", err))
	}
	return base64.StdEncoding.EncodeToString(buf)
}
