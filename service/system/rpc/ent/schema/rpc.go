package schema

import "entgo.io/ent"

// Rpc holds the schema definition for the Rpc entity.
type Rpc struct {
	ent.Schema
}

// Fields of the Rpc.
func (Rpc) Fields() []ent.Field {
	return nil
}

// Edges of the Rpc.
func (Rpc) Edges() []ent.Edge {
	return nil
}
