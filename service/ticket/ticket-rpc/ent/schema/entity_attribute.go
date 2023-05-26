package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type EntityAttribute struct {
	ent.Schema
}

func (EntityAttribute) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("attribute_id").
			Comment("Attribute Id | 属性ID").
			Annotations(entsql.WithComments(true)),
		field.Uint64("entity_id").
			Comment("Entity Id | 实体ID").
			Annotations(entsql.WithComments(true)),
		field.Uint64("attribute_set_id").
			Comment("Attribute Set Id | 属性集表ID").
			Annotations(entsql.WithComments(true)),
		field.Uint64("attribute_group_id").
			Comment("Attribute Group Id | 属性组表ID").
			Annotations(entsql.WithComments(true)),
		field.Uint8("sequence").
			Comment("sequence | 顺序").
			Annotations(entsql.WithComments(true)),
	}
}

func (EntityAttribute) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
	}
}

func (EntityAttribute) Edges() []ent.Edge {
	return nil
}

func (EntityAttribute) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("attribute_id"),
		index.Fields("entity_id"),
		index.Fields("attribute_set_id"),
		index.Fields("attribute_group_id"),
	}
}

func (EntityAttribute) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "entity_attributes"},
	}
}
