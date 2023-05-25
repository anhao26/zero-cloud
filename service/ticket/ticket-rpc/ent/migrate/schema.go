// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AttributesColumns holds the columns for the "Attributes" table.
	AttributesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "entity_id", Type: field.TypeUint64, Comment: "Entity Id | 实体ID"},
		{Name: "attribute_code", Type: field.TypeString, Comment: "Attribute Code | 属性标识"},
		{Name: "backend_class", Type: field.TypeString, Comment: "Backend Class | 后端类:指定时将用于在属性与数据库结合时向属性添加附加控制"},
		{Name: "backend_type", Type: field.TypeString, Comment: "Backend Type | 指定列类型"},
		{Name: "backend_table", Type: field.TypeString, Comment: "Backend Table | 后端表:当指定时，它将数据存储到给定的文档"},
		{Name: "frontend_class", Type: field.TypeString, Comment: "Frontend Class | 前端类:当在前端使用时，将用于向属性添加附加控制"},
		{Name: "frontend_type", Type: field.TypeString, Comment: "Frontend Type | 前端类型:指定 html 字段的类型"},
		{Name: "frontend_label", Type: field.TypeString, Comment: "Frontend Label | 前端标签:指定标签"},
		{Name: "source_class", Type: field.TypeString, Comment: "Source Class | 指定时将用于填充字段的默认选项，如果 frontend_type 是select"},
		{Name: "default_value", Type: field.TypeString, Size: 2147483647, Comment: "Default Value | 默认值"},
		{Name: "is_filterable", Type: field.TypeUint8, Comment: "Is Filterable | 如果启用，属性将包含在分面导航中", Default: 0},
		{Name: "is_searchable", Type: field.TypeUint8, Comment: "Is Searchable | 是否支持搜索", Default: 0},
		{Name: "is_required", Type: field.TypeUint8, Comment: "Is Required | 是否必填 0不必填 1 必填", Default: 0},
		{Name: "required_validate_class", Type: field.TypeString, Comment: "Required Validate Class | 必需的验证类"},
	}
	// AttributesTable holds the schema information for the "Attributes" table.
	AttributesTable = &schema.Table{
		Name:       "Attributes",
		Columns:    AttributesColumns,
		PrimaryKey: []*schema.Column{AttributesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "attribute_entity_id_attribute_code",
				Unique:  true,
				Columns: []*schema.Column{AttributesColumns[3], AttributesColumns[4]},
			},
		},
	}
	// EntitiesColumns holds the columns for the "entities" table.
	EntitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "entity_code", Type: field.TypeString, Unique: true, Comment: "Entity Code | 实体标识"},
		{Name: "entity_class", Type: field.TypeString, Comment: "Entity Class | 实体类"},
		{Name: "entity_table", Type: field.TypeString, Comment: "Entity Table | 实体表名称"},
		{Name: "default_attribute_set_id", Type: field.TypeUint64, Comment: "Default Attribute Set Id | 默认属性SET ID"},
		{Name: "additional_attribute_table", Type: field.TypeString, Comment: "Additional Attribute Table | 附加属性表"},
		{Name: "is_flat_enabled", Type: field.TypeUint32, Comment: "Is Flat Enabled | 是否设置一张表存储", Default: 0},
	}
	// EntitiesTable holds the schema information for the "entities" table.
	EntitiesTable = &schema.Table{
		Name:       "entities",
		Columns:    EntitiesColumns,
		PrimaryKey: []*schema.Column{EntitiesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AttributesTable,
		EntitiesTable,
	}
)

func init() {
	AttributesTable.Annotation = &entsql.Annotation{
		Table: "Attributes",
	}
	EntitiesTable.Annotation = &entsql.Annotation{
		Table: "entities",
	}
}
