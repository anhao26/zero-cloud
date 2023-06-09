// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attribute"
)

// Attribute is the model entity for the Attribute schema.
type Attribute struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Entity Id | 实体ID
	EntityID uint64 `json:"entity_id,omitempty"`
	// Attribute Code | 属性标识
	AttributeCode string `json:"attribute_code,omitempty"`
	// Backend Class | 后端类:指定时将用于在属性与数据库结合时向属性添加附加控制
	BackendClass string `json:"backend_class,omitempty"`
	// Backend Type | 指定列类型
	BackendType string `json:"backend_type,omitempty"`
	// Backend Table | 后端表:当指定时，它将数据存储到给定的文档
	BackendTable string `json:"backend_table,omitempty"`
	// Frontend Class | 前端类:当在前端使用时，将用于向属性添加附加控制
	FrontendClass string `json:"frontend_class,omitempty"`
	// Frontend Type | 前端类型:指定 html 字段的类型
	FrontendType string `json:"frontend_type,omitempty"`
	// Frontend Label | 前端标签:指定标签
	FrontendLabel string `json:"frontend_label,omitempty"`
	// Source Class | 指定时将用于填充字段的默认选项，如果 frontend_type 是select
	SourceClass string `json:"source_class,omitempty"`
	// Default Value | 默认值
	DefaultValue string `json:"default_value,omitempty"`
	// Is Filterable | 如果启用，属性将包含在分面导航中
	IsFilterable uint8 `json:"is_filterable,omitempty"`
	// Is Searchable | 是否支持搜索
	IsSearchable uint8 `json:"is_searchable,omitempty"`
	// Is Required | 是否必填 0不必填 1 必填
	IsRequired uint8 `json:"is_required,omitempty"`
	// Required Validate Class | 必需的验证类
	RequiredValidateClass string `json:"required_validate_class,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AttributeQuery when eager-loading is set.
	Edges        AttributeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AttributeEdges holds the relations/edges for other nodes in the graph.
type AttributeEdges struct {
	// OptionID holds the value of the option_id edge.
	OptionID []*AttributeOption `json:"option_id,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OptionIDOrErr returns the OptionID value or an error if the edge
// was not loaded in eager-loading.
func (e AttributeEdges) OptionIDOrErr() ([]*AttributeOption, error) {
	if e.loadedTypes[0] {
		return e.OptionID, nil
	}
	return nil, &NotLoadedError{edge: "option_id"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Attribute) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attribute.FieldID, attribute.FieldEntityID, attribute.FieldIsFilterable, attribute.FieldIsSearchable, attribute.FieldIsRequired:
			values[i] = new(sql.NullInt64)
		case attribute.FieldAttributeCode, attribute.FieldBackendClass, attribute.FieldBackendType, attribute.FieldBackendTable, attribute.FieldFrontendClass, attribute.FieldFrontendType, attribute.FieldFrontendLabel, attribute.FieldSourceClass, attribute.FieldDefaultValue, attribute.FieldRequiredValidateClass:
			values[i] = new(sql.NullString)
		case attribute.FieldCreatedAt, attribute.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Attribute fields.
func (a *Attribute) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attribute.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint64(value.Int64)
		case attribute.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case attribute.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case attribute.FieldEntityID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field entity_id", values[i])
			} else if value.Valid {
				a.EntityID = uint64(value.Int64)
			}
		case attribute.FieldAttributeCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field attribute_code", values[i])
			} else if value.Valid {
				a.AttributeCode = value.String
			}
		case attribute.FieldBackendClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field backend_class", values[i])
			} else if value.Valid {
				a.BackendClass = value.String
			}
		case attribute.FieldBackendType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field backend_type", values[i])
			} else if value.Valid {
				a.BackendType = value.String
			}
		case attribute.FieldBackendTable:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field backend_table", values[i])
			} else if value.Valid {
				a.BackendTable = value.String
			}
		case attribute.FieldFrontendClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field frontend_class", values[i])
			} else if value.Valid {
				a.FrontendClass = value.String
			}
		case attribute.FieldFrontendType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field frontend_type", values[i])
			} else if value.Valid {
				a.FrontendType = value.String
			}
		case attribute.FieldFrontendLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field frontend_label", values[i])
			} else if value.Valid {
				a.FrontendLabel = value.String
			}
		case attribute.FieldSourceClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source_class", values[i])
			} else if value.Valid {
				a.SourceClass = value.String
			}
		case attribute.FieldDefaultValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field default_value", values[i])
			} else if value.Valid {
				a.DefaultValue = value.String
			}
		case attribute.FieldIsFilterable:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_filterable", values[i])
			} else if value.Valid {
				a.IsFilterable = uint8(value.Int64)
			}
		case attribute.FieldIsSearchable:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_searchable", values[i])
			} else if value.Valid {
				a.IsSearchable = uint8(value.Int64)
			}
		case attribute.FieldIsRequired:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_required", values[i])
			} else if value.Valid {
				a.IsRequired = uint8(value.Int64)
			}
		case attribute.FieldRequiredValidateClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field required_validate_class", values[i])
			} else if value.Valid {
				a.RequiredValidateClass = value.String
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Attribute.
// This includes values selected through modifiers, order, etc.
func (a *Attribute) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryOptionID queries the "option_id" edge of the Attribute entity.
func (a *Attribute) QueryOptionID() *AttributeOptionQuery {
	return NewAttributeClient(a.config).QueryOptionID(a)
}

// Update returns a builder for updating this Attribute.
// Note that you need to call Attribute.Unwrap() before calling this method if this Attribute
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Attribute) Update() *AttributeUpdateOne {
	return NewAttributeClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Attribute entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Attribute) Unwrap() *Attribute {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Attribute is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Attribute) String() string {
	var builder strings.Builder
	builder.WriteString("Attribute(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("entity_id=")
	builder.WriteString(fmt.Sprintf("%v", a.EntityID))
	builder.WriteString(", ")
	builder.WriteString("attribute_code=")
	builder.WriteString(a.AttributeCode)
	builder.WriteString(", ")
	builder.WriteString("backend_class=")
	builder.WriteString(a.BackendClass)
	builder.WriteString(", ")
	builder.WriteString("backend_type=")
	builder.WriteString(a.BackendType)
	builder.WriteString(", ")
	builder.WriteString("backend_table=")
	builder.WriteString(a.BackendTable)
	builder.WriteString(", ")
	builder.WriteString("frontend_class=")
	builder.WriteString(a.FrontendClass)
	builder.WriteString(", ")
	builder.WriteString("frontend_type=")
	builder.WriteString(a.FrontendType)
	builder.WriteString(", ")
	builder.WriteString("frontend_label=")
	builder.WriteString(a.FrontendLabel)
	builder.WriteString(", ")
	builder.WriteString("source_class=")
	builder.WriteString(a.SourceClass)
	builder.WriteString(", ")
	builder.WriteString("default_value=")
	builder.WriteString(a.DefaultValue)
	builder.WriteString(", ")
	builder.WriteString("is_filterable=")
	builder.WriteString(fmt.Sprintf("%v", a.IsFilterable))
	builder.WriteString(", ")
	builder.WriteString("is_searchable=")
	builder.WriteString(fmt.Sprintf("%v", a.IsSearchable))
	builder.WriteString(", ")
	builder.WriteString("is_required=")
	builder.WriteString(fmt.Sprintf("%v", a.IsRequired))
	builder.WriteString(", ")
	builder.WriteString("required_validate_class=")
	builder.WriteString(a.RequiredValidateClass)
	builder.WriteByte(')')
	return builder.String()
}

// Attributes is a parsable slice of Attribute.
type Attributes []*Attribute
