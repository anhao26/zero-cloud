package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type AttributeGroup struct {
	ent.Schema
}

func (AttributeGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("attribute_set_id").
			Comment("Attribute Set Id | 属性集ID").
			Annotations(entsql.WithComments(true)),
		field.String("attribute_group_name").
			Comment("Attribute Group Name | 属性组名").
			Annotations(entsql.WithComments(true)),
		field.Uint8("sequence").
			Comment("sequence | 顺序").
			Annotations(entsql.WithComments(true)),
	}
}

func (AttributeGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
	}
}

func (AttributeGroup) Edges() []ent.Edge {
	return nil
}

func (AttributeGroup) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("attribute_set_id", "attribute_group_name").
			Unique(),
	}
}

func (AttributeGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "attribute_groups"},
	}
}
