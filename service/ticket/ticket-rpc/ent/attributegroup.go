// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attributegroup"
)

// AttributeGroup is the model entity for the AttributeGroup schema.
type AttributeGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Attribute Set Id | 属性集ID
	AttributeSetID uint64 `json:"attribute_set_id,omitempty"`
	// Attribute Group Name | 属性组名
	AttributeGroupName string `json:"attribute_group_name,omitempty"`
	// sequence | 顺序
	Sequence     uint8 `json:"sequence,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AttributeGroup) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attributegroup.FieldID, attributegroup.FieldAttributeSetID, attributegroup.FieldSequence:
			values[i] = new(sql.NullInt64)
		case attributegroup.FieldAttributeGroupName:
			values[i] = new(sql.NullString)
		case attributegroup.FieldCreatedAt, attributegroup.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AttributeGroup fields.
func (ag *AttributeGroup) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attributegroup.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ag.ID = uint64(value.Int64)
		case attributegroup.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ag.CreatedAt = value.Time
			}
		case attributegroup.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ag.UpdatedAt = value.Time
			}
		case attributegroup.FieldAttributeSetID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field attribute_set_id", values[i])
			} else if value.Valid {
				ag.AttributeSetID = uint64(value.Int64)
			}
		case attributegroup.FieldAttributeGroupName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field attribute_group_name", values[i])
			} else if value.Valid {
				ag.AttributeGroupName = value.String
			}
		case attributegroup.FieldSequence:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sequence", values[i])
			} else if value.Valid {
				ag.Sequence = uint8(value.Int64)
			}
		default:
			ag.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AttributeGroup.
// This includes values selected through modifiers, order, etc.
func (ag *AttributeGroup) Value(name string) (ent.Value, error) {
	return ag.selectValues.Get(name)
}

// Update returns a builder for updating this AttributeGroup.
// Note that you need to call AttributeGroup.Unwrap() before calling this method if this AttributeGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (ag *AttributeGroup) Update() *AttributeGroupUpdateOne {
	return NewAttributeGroupClient(ag.config).UpdateOne(ag)
}

// Unwrap unwraps the AttributeGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ag *AttributeGroup) Unwrap() *AttributeGroup {
	_tx, ok := ag.config.driver.(*txDriver)
	if !ok {
		panic("ent: AttributeGroup is not a transactional entity")
	}
	ag.config.driver = _tx.drv
	return ag
}

// String implements the fmt.Stringer.
func (ag *AttributeGroup) String() string {
	var builder strings.Builder
	builder.WriteString("AttributeGroup(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ag.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ag.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ag.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("attribute_set_id=")
	builder.WriteString(fmt.Sprintf("%v", ag.AttributeSetID))
	builder.WriteString(", ")
	builder.WriteString("attribute_group_name=")
	builder.WriteString(ag.AttributeGroupName)
	builder.WriteString(", ")
	builder.WriteString("sequence=")
	builder.WriteString(fmt.Sprintf("%v", ag.Sequence))
	builder.WriteByte(')')
	return builder.String()
}

// AttributeGroups is a parsable slice of AttributeGroup.
type AttributeGroups []*AttributeGroup