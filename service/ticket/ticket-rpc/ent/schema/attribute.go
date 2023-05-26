package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Attribute struct {
	ent.Schema
}

func (Attribute) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("entity_id").
			Comment("Entity Id | 实体ID").
			Annotations(entsql.WithComments(true)),
		field.String("attribute_code").
			Comment("Attribute Code | 属性标识").
			Annotations(entsql.WithComments(true)),
		field.String("backend_class").
			Comment("Backend Class | 后端类:指定时将用于在属性与数据库结合时向属性添加附加控制").
			Annotations(entsql.WithComments(true)),
		field.String("backend_type").
			Comment("Backend Type | 指定列类型").
			Annotations(entsql.WithComments(true)),
		field.String("backend_table").
			Comment("Backend Table | 后端表:当指定时，它将数据存储到给定的文档").
			Annotations(entsql.WithComments(true)),
		field.String("frontend_class").
			Comment("Frontend Class | 前端类:当在前端使用时，将用于向属性添加附加控制").
			Annotations(entsql.WithComments(true)),
		field.String("frontend_type").
			Comment("Frontend Type | 前端类型:指定 html 字段的类型").
			Annotations(entsql.WithComments(true)),
		field.String("frontend_label").
			Comment("Frontend Label | 前端标签:指定标签").
			Annotations(entsql.WithComments(true)),
		field.String("source_class").
			Comment("Source Class | 指定时将用于填充字段的默认选项，如果 frontend_type 是select").
			Annotations(entsql.WithComments(true)),
		field.Text("default_value").
			Comment("Default Value | 默认值").
			Annotations(entsql.WithComments(true)),
		field.Uint8("is_filterable").Default(0).
			Comment("Is Filterable | 如果启用，属性将包含在分面导航中").
			Annotations(entsql.WithComments(true)),
		field.Uint8("is_searchable").Default(0).
			Comment("Is Searchable | 是否支持搜索").
			Annotations(entsql.WithComments(true)),
		field.Uint8("is_required").Default(0).
			Comment("Is Required | 是否必填 0不必填 1 必填").
			Annotations(entsql.WithComments(true)),
		field.String("required_validate_class").
			Comment("Required Validate Class | 必需的验证类").
			Annotations(entsql.WithComments(true)),
	}
}

func (Attribute) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
	}
}

func (Attribute) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("entities", Entity.Type),
	}
}

func (Attribute) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("entity_id", "attribute_code").
			Unique(),
	}
}

func (Attribute) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Attributes"},
	}
}
