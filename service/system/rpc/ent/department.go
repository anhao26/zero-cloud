// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/anhao26/zero-cloud/service/system/rpc/ent/department"
)

// Department is the model entity for the Department schema.
type Department struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// status 1 normal 2 ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// Sort number | 排序编号
	Sort uint32 `json:"sort,omitempty"`
	// Department name | 部门名称
	Name string `json:"name,omitempty"`
	// Parents' IDs | 父级列表
	Ancestors string `json:"ancestors,omitempty"`
	// Department leader | 部门负责人
	Leader string `json:"leader,omitempty"`
	// Leader's phone number | 负责人电话
	Phone string `json:"phone,omitempty"`
	// Leader's email | 部门负责人电子邮箱
	Email string `json:"email,omitempty"`
	// Remark | 备注
	Remark string `json:"remark,omitempty"`
	// Parent department ID | 父级部门ID
	ParentID uint64 `json:"parent_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DepartmentQuery when eager-loading is set.
	Edges        DepartmentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DepartmentEdges holds the relations/edges for other nodes in the graph.
type DepartmentEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Department `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Department `json:"children,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DepartmentEdges) ParentOrErr() (*Department, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: department.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) ChildrenOrErr() ([]*Department, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Department) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case department.FieldID, department.FieldStatus, department.FieldSort, department.FieldParentID:
			values[i] = new(sql.NullInt64)
		case department.FieldName, department.FieldAncestors, department.FieldLeader, department.FieldPhone, department.FieldEmail, department.FieldRemark:
			values[i] = new(sql.NullString)
		case department.FieldCreatedAt, department.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Department fields.
func (d *Department) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case department.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = uint64(value.Int64)
		case department.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case department.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Time
			}
		case department.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				d.Status = uint8(value.Int64)
			}
		case department.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				d.Sort = uint32(value.Int64)
			}
		case department.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case department.FieldAncestors:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ancestors", values[i])
			} else if value.Valid {
				d.Ancestors = value.String
			}
		case department.FieldLeader:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field leader", values[i])
			} else if value.Valid {
				d.Leader = value.String
			}
		case department.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				d.Phone = value.String
			}
		case department.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				d.Email = value.String
			}
		case department.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				d.Remark = value.String
			}
		case department.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				d.ParentID = uint64(value.Int64)
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Department.
// This includes values selected through modifiers, order, etc.
func (d *Department) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryParent queries the "parent" edge of the Department entity.
func (d *Department) QueryParent() *DepartmentQuery {
	return NewDepartmentClient(d.config).QueryParent(d)
}

// QueryChildren queries the "children" edge of the Department entity.
func (d *Department) QueryChildren() *DepartmentQuery {
	return NewDepartmentClient(d.config).QueryChildren(d)
}

// QueryUsers queries the "users" edge of the Department entity.
func (d *Department) QueryUsers() *UserQuery {
	return NewDepartmentClient(d.config).QueryUsers(d)
}

// Update returns a builder for updating this Department.
// Note that you need to call Department.Unwrap() before calling this method if this Department
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Department) Update() *DepartmentUpdateOne {
	return NewDepartmentClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Department entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Department) Unwrap() *Department {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Department is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Department) String() string {
	var builder strings.Builder
	builder.WriteString("Department(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", d.Status))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", d.Sort))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("ancestors=")
	builder.WriteString(d.Ancestors)
	builder.WriteString(", ")
	builder.WriteString("leader=")
	builder.WriteString(d.Leader)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(d.Phone)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(d.Email)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(d.Remark)
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", d.ParentID))
	builder.WriteByte(')')
	return builder.String()
}

// Departments is a parsable slice of Department.
type Departments []*Department
