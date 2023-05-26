// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attributeoption"
)

// AttributeOption is the model entity for the AttributeOption schema.
type AttributeOption struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Attribute Id | 属性ID
	AttributeID uint64 `json:"attribute_id,omitempty"`
	// Label | 选项名
	Label string `json:"label,omitempty"`
	// value | 选项值
	Value                       uint64 `json:"value,omitempty"`
	attribute_attribute_options *uint64
	selectValues                sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AttributeOption) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attributeoption.FieldID, attributeoption.FieldAttributeID, attributeoption.FieldValue:
			values[i] = new(sql.NullInt64)
		case attributeoption.FieldLabel:
			values[i] = new(sql.NullString)
		case attributeoption.FieldCreatedAt, attributeoption.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case attributeoption.ForeignKeys[0]: // attribute_attribute_options
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AttributeOption fields.
func (ao *AttributeOption) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attributeoption.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ao.ID = uint64(value.Int64)
		case attributeoption.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ao.CreatedAt = value.Time
			}
		case attributeoption.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ao.UpdatedAt = value.Time
			}
		case attributeoption.FieldAttributeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field attribute_id", values[i])
			} else if value.Valid {
				ao.AttributeID = uint64(value.Int64)
			}
		case attributeoption.FieldLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field label", values[i])
			} else if value.Valid {
				ao.Label = value.String
			}
		case attributeoption.FieldValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				ao.Value = uint64(value.Int64)
			}
		case attributeoption.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field attribute_attribute_options", value)
			} else if value.Valid {
				ao.attribute_attribute_options = new(uint64)
				*ao.attribute_attribute_options = uint64(value.Int64)
			}
		default:
			ao.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the AttributeOption.
// This includes values selected through modifiers, order, etc.
func (ao *AttributeOption) GetValue(name string) (ent.Value, error) {
	return ao.selectValues.Get(name)
}

// Update returns a builder for updating this AttributeOption.
// Note that you need to call AttributeOption.Unwrap() before calling this method if this AttributeOption
// was returned from a transaction, and the transaction was committed or rolled back.
func (ao *AttributeOption) Update() *AttributeOptionUpdateOne {
	return NewAttributeOptionClient(ao.config).UpdateOne(ao)
}

// Unwrap unwraps the AttributeOption entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ao *AttributeOption) Unwrap() *AttributeOption {
	_tx, ok := ao.config.driver.(*txDriver)
	if !ok {
		panic("ent: AttributeOption is not a transactional entity")
	}
	ao.config.driver = _tx.drv
	return ao
}

// String implements the fmt.Stringer.
func (ao *AttributeOption) String() string {
	var builder strings.Builder
	builder.WriteString("AttributeOption(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ao.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ao.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ao.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("attribute_id=")
	builder.WriteString(fmt.Sprintf("%v", ao.AttributeID))
	builder.WriteString(", ")
	builder.WriteString("label=")
	builder.WriteString(ao.Label)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(fmt.Sprintf("%v", ao.Value))
	builder.WriteByte(')')
	return builder.String()
}

// AttributeOptions is a parsable slice of AttributeOption.
type AttributeOptions []*AttributeOption
