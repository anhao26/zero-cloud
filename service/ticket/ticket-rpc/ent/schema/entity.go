package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Entity struct {
	ent.Schema
}

func (Entity) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").
			Comment("API path | API 路径").
			Annotations(entsql.WithComments(true)),
		field.String("description").
			Comment("API description | API 描述").
			Annotations(entsql.WithComments(true)),
		field.String("api_group").
			Comment("API group | API 分组").
			Annotations(entsql.WithComments(true)),
		field.String("method").Default("POST").
			Comment("HTTP method | HTTP 请求类型").
			Annotations(entsql.WithComments(true)),
	}
}

func (Entity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
	}
}

// Edges of the Ticket.
func (Entity) Edges() []ent.Edge {
	return nil
}
