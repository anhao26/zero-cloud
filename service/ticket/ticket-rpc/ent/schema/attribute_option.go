package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type AttributeOption struct {
	ent.Schema
}

func (AttributeOption) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("attribute_id").
			Comment("Attribute Id | 属性ID").
			Annotations(entsql.WithComments(true)),
		field.String("label").
			Comment("Label | 选项名").
			Annotations(entsql.WithComments(true)),
		field.Uint64("value").
			Comment("value | 选项值").
			Annotations(entsql.WithComments(true)),
	}
}

func (AttributeOption) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
	}
}

func (AttributeOption) Edges() []ent.Edge {
	return nil
}

func (AttributeOption) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("attribute_id"),
	}
}

func (AttributeOption) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "attribute_options"},
	}
}
