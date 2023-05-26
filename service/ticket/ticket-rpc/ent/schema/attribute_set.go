package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type AttributeSet struct {
	ent.Schema
}

func (AttributeSet) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("entity_id").
			Comment("Entity Id | 实体ID").
			Annotations(entsql.WithComments(true)),
		field.String("attribute_set_name").
			Comment("Attribute Set Name | 属性集名字").
			Annotations(entsql.WithComments(true)),
	}
}

func (AttributeSet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
	}
}

func (AttributeSet) Edges() []ent.Edge {
	return nil
}

func (AttributeSet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("entity_id", "attribute_set_name").
			Unique(),
	}
}

func (AttributeSet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "attribute_sets"},
	}
}
