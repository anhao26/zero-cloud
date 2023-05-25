package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Entity struct {
	ent.Schema
}

func (Entity) Fields() []ent.Field {
	return []ent.Field{
		field.String("entity_code").
			Comment("Entity Code | 实体标识").
			Annotations(entsql.WithComments(true)).
			Unique(),
		field.String("entity_class").
			Comment("Entity Class | 实体类").
			Annotations(entsql.WithComments(true)),
		field.String("entity_table").
			Comment("Entity Table | 实体表名称").
			Annotations(entsql.WithComments(true)),
		field.Uint64("default_attribute_set_id").
			Comment("Default Attribute Set Id | 默认属性SET ID").
			Annotations(entsql.WithComments(true)),
		field.String("additional_attribute_table").
			Comment("Additional Attribute Table | 附加属性表").
			Annotations(entsql.WithComments(true)),
		field.Uint32("is_flat_enabled").Default(0).
			Comment("Is Flat Enabled | 是否设置一张表存储").
			Annotations(entsql.WithComments(true)),
	}
}

func (Entity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
	}
}

func (Entity) Edges() []ent.Edge {
	return nil
}

func (Entity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "entities"},
	}
}
