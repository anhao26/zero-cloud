// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attributeset"
)

// AttributeSet is the model entity for the AttributeSet schema.
type AttributeSet struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Entity Id | 实体ID
	EntityID uint64 `json:"entity_id,omitempty"`
	// Attribute Set Name | 属性集名字
	AttributeSetName string `json:"attribute_set_name,omitempty"`
	selectValues     sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AttributeSet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attributeset.FieldID, attributeset.FieldEntityID:
			values[i] = new(sql.NullInt64)
		case attributeset.FieldAttributeSetName:
			values[i] = new(sql.NullString)
		case attributeset.FieldCreatedAt, attributeset.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AttributeSet fields.
func (as *AttributeSet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attributeset.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			as.ID = uint64(value.Int64)
		case attributeset.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				as.CreatedAt = value.Time
			}
		case attributeset.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				as.UpdatedAt = value.Time
			}
		case attributeset.FieldEntityID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field entity_id", values[i])
			} else if value.Valid {
				as.EntityID = uint64(value.Int64)
			}
		case attributeset.FieldAttributeSetName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field attribute_set_name", values[i])
			} else if value.Valid {
				as.AttributeSetName = value.String
			}
		default:
			as.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AttributeSet.
// This includes values selected through modifiers, order, etc.
func (as *AttributeSet) Value(name string) (ent.Value, error) {
	return as.selectValues.Get(name)
}

// Update returns a builder for updating this AttributeSet.
// Note that you need to call AttributeSet.Unwrap() before calling this method if this AttributeSet
// was returned from a transaction, and the transaction was committed or rolled back.
func (as *AttributeSet) Update() *AttributeSetUpdateOne {
	return NewAttributeSetClient(as.config).UpdateOne(as)
}

// Unwrap unwraps the AttributeSet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (as *AttributeSet) Unwrap() *AttributeSet {
	_tx, ok := as.config.driver.(*txDriver)
	if !ok {
		panic("ent: AttributeSet is not a transactional entity")
	}
	as.config.driver = _tx.drv
	return as
}

// String implements the fmt.Stringer.
func (as *AttributeSet) String() string {
	var builder strings.Builder
	builder.WriteString("AttributeSet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", as.ID))
	builder.WriteString("created_at=")
	builder.WriteString(as.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(as.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("entity_id=")
	builder.WriteString(fmt.Sprintf("%v", as.EntityID))
	builder.WriteString(", ")
	builder.WriteString("attribute_set_name=")
	builder.WriteString(as.AttributeSetName)
	builder.WriteByte(')')
	return builder.String()
}

// AttributeSets is a parsable slice of AttributeSet.
type AttributeSets []*AttributeSet
