// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/student"
	uuid "github.com/gofrs/uuid/v5"
)

// StudentCreate is the builder for creating a Student entity.
type StudentCreate struct {
	config
	mutation *StudentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *StudentCreate) SetCreatedAt(t time.Time) *StudentCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StudentCreate) SetNillableCreatedAt(t *time.Time) *StudentCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *StudentCreate) SetUpdatedAt(t time.Time) *StudentCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *StudentCreate) SetNillableUpdatedAt(t *time.Time) *StudentCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *StudentCreate) SetName(s string) *StudentCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetAge sets the "age" field.
func (sc *StudentCreate) SetAge(i int) *StudentCreate {
	sc.mutation.SetAge(i)
	return sc
}

// SetAgeInt8 sets the "age_int8" field.
func (sc *StudentCreate) SetAgeInt8(i int8) *StudentCreate {
	sc.mutation.SetAgeInt8(i)
	return sc
}

// SetAgeUint8 sets the "age_uint8" field.
func (sc *StudentCreate) SetAgeUint8(u uint8) *StudentCreate {
	sc.mutation.SetAgeUint8(u)
	return sc
}

// SetAgeInt16 sets the "age_int16" field.
func (sc *StudentCreate) SetAgeInt16(i int16) *StudentCreate {
	sc.mutation.SetAgeInt16(i)
	return sc
}

// SetAgeUint16 sets the "age_uint16" field.
func (sc *StudentCreate) SetAgeUint16(u uint16) *StudentCreate {
	sc.mutation.SetAgeUint16(u)
	return sc
}

// SetAgeInt32 sets the "age_int32" field.
func (sc *StudentCreate) SetAgeInt32(i int32) *StudentCreate {
	sc.mutation.SetAgeInt32(i)
	return sc
}

// SetAgeUint32 sets the "age_uint32" field.
func (sc *StudentCreate) SetAgeUint32(u uint32) *StudentCreate {
	sc.mutation.SetAgeUint32(u)
	return sc
}

// SetAgeInt64 sets the "age_int64" field.
func (sc *StudentCreate) SetAgeInt64(i int64) *StudentCreate {
	sc.mutation.SetAgeInt64(i)
	return sc
}

// SetAgeUint64 sets the "age_uint64" field.
func (sc *StudentCreate) SetAgeUint64(u uint64) *StudentCreate {
	sc.mutation.SetAgeUint64(u)
	return sc
}

// SetAgeInt sets the "age_int" field.
func (sc *StudentCreate) SetAgeInt(i int) *StudentCreate {
	sc.mutation.SetAgeInt(i)
	return sc
}

// SetAgeUint sets the "age_uint" field.
func (sc *StudentCreate) SetAgeUint(u uint) *StudentCreate {
	sc.mutation.SetAgeUint(u)
	return sc
}

// SetWeightFloat sets the "weight_float" field.
func (sc *StudentCreate) SetWeightFloat(f float64) *StudentCreate {
	sc.mutation.SetWeightFloat(f)
	return sc
}

// SetWeightFloat32 sets the "weight_float32" field.
func (sc *StudentCreate) SetWeightFloat32(f float32) *StudentCreate {
	sc.mutation.SetWeightFloat32(f)
	return sc
}

// SetClassID sets the "class_id" field.
func (sc *StudentCreate) SetClassID(u uuid.UUID) *StudentCreate {
	sc.mutation.SetClassID(u)
	return sc
}

// SetEnrollAt sets the "enroll_at" field.
func (sc *StudentCreate) SetEnrollAt(t time.Time) *StudentCreate {
	sc.mutation.SetEnrollAt(t)
	return sc
}

// SetStatusBool sets the "status_bool" field.
func (sc *StudentCreate) SetStatusBool(b bool) *StudentCreate {
	sc.mutation.SetStatusBool(b)
	return sc
}

// SetID sets the "id" field.
func (sc *StudentCreate) SetID(u uint64) *StudentCreate {
	sc.mutation.SetID(u)
	return sc
}

// Mutation returns the StudentMutation object of the builder.
func (sc *StudentCreate) Mutation() *StudentMutation {
	return sc.mutation
}

// Save creates the Student in the database.
func (sc *StudentCreate) Save(ctx context.Context) (*Student, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StudentCreate) SaveX(ctx context.Context) *Student {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StudentCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StudentCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StudentCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := student.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := student.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StudentCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Student.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Student.updated_at"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Student.name"`)}
	}
	if _, ok := sc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Student.age"`)}
	}
	if _, ok := sc.mutation.AgeInt8(); !ok {
		return &ValidationError{Name: "age_int8", err: errors.New(`ent: missing required field "Student.age_int8"`)}
	}
	if _, ok := sc.mutation.AgeUint8(); !ok {
		return &ValidationError{Name: "age_uint8", err: errors.New(`ent: missing required field "Student.age_uint8"`)}
	}
	if _, ok := sc.mutation.AgeInt16(); !ok {
		return &ValidationError{Name: "age_int16", err: errors.New(`ent: missing required field "Student.age_int16"`)}
	}
	if _, ok := sc.mutation.AgeUint16(); !ok {
		return &ValidationError{Name: "age_uint16", err: errors.New(`ent: missing required field "Student.age_uint16"`)}
	}
	if _, ok := sc.mutation.AgeInt32(); !ok {
		return &ValidationError{Name: "age_int32", err: errors.New(`ent: missing required field "Student.age_int32"`)}
	}
	if _, ok := sc.mutation.AgeUint32(); !ok {
		return &ValidationError{Name: "age_uint32", err: errors.New(`ent: missing required field "Student.age_uint32"`)}
	}
	if _, ok := sc.mutation.AgeInt64(); !ok {
		return &ValidationError{Name: "age_int64", err: errors.New(`ent: missing required field "Student.age_int64"`)}
	}
	if _, ok := sc.mutation.AgeUint64(); !ok {
		return &ValidationError{Name: "age_uint64", err: errors.New(`ent: missing required field "Student.age_uint64"`)}
	}
	if _, ok := sc.mutation.AgeInt(); !ok {
		return &ValidationError{Name: "age_int", err: errors.New(`ent: missing required field "Student.age_int"`)}
	}
	if _, ok := sc.mutation.AgeUint(); !ok {
		return &ValidationError{Name: "age_uint", err: errors.New(`ent: missing required field "Student.age_uint"`)}
	}
	if _, ok := sc.mutation.WeightFloat(); !ok {
		return &ValidationError{Name: "weight_float", err: errors.New(`ent: missing required field "Student.weight_float"`)}
	}
	if _, ok := sc.mutation.WeightFloat32(); !ok {
		return &ValidationError{Name: "weight_float32", err: errors.New(`ent: missing required field "Student.weight_float32"`)}
	}
	if _, ok := sc.mutation.ClassID(); !ok {
		return &ValidationError{Name: "class_id", err: errors.New(`ent: missing required field "Student.class_id"`)}
	}
	if _, ok := sc.mutation.EnrollAt(); !ok {
		return &ValidationError{Name: "enroll_at", err: errors.New(`ent: missing required field "Student.enroll_at"`)}
	}
	if _, ok := sc.mutation.StatusBool(); !ok {
		return &ValidationError{Name: "status_bool", err: errors.New(`ent: missing required field "Student.status_bool"`)}
	}
	return nil
}

func (sc *StudentCreate) sqlSave(ctx context.Context) (*Student, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StudentCreate) createSpec() (*Student, *sqlgraph.CreateSpec) {
	var (
		_node = &Student{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(student.Table, sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint64))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(student.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(student.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Age(); ok {
		_spec.SetField(student.FieldAge, field.TypeInt, value)
		_node.Age = value
	}
	if value, ok := sc.mutation.AgeInt8(); ok {
		_spec.SetField(student.FieldAgeInt8, field.TypeInt8, value)
		_node.AgeInt8 = value
	}
	if value, ok := sc.mutation.AgeUint8(); ok {
		_spec.SetField(student.FieldAgeUint8, field.TypeUint8, value)
		_node.AgeUint8 = value
	}
	if value, ok := sc.mutation.AgeInt16(); ok {
		_spec.SetField(student.FieldAgeInt16, field.TypeInt16, value)
		_node.AgeInt16 = value
	}
	if value, ok := sc.mutation.AgeUint16(); ok {
		_spec.SetField(student.FieldAgeUint16, field.TypeUint16, value)
		_node.AgeUint16 = value
	}
	if value, ok := sc.mutation.AgeInt32(); ok {
		_spec.SetField(student.FieldAgeInt32, field.TypeInt32, value)
		_node.AgeInt32 = value
	}
	if value, ok := sc.mutation.AgeUint32(); ok {
		_spec.SetField(student.FieldAgeUint32, field.TypeUint32, value)
		_node.AgeUint32 = value
	}
	if value, ok := sc.mutation.AgeInt64(); ok {
		_spec.SetField(student.FieldAgeInt64, field.TypeInt64, value)
		_node.AgeInt64 = value
	}
	if value, ok := sc.mutation.AgeUint64(); ok {
		_spec.SetField(student.FieldAgeUint64, field.TypeUint64, value)
		_node.AgeUint64 = value
	}
	if value, ok := sc.mutation.AgeInt(); ok {
		_spec.SetField(student.FieldAgeInt, field.TypeInt, value)
		_node.AgeInt = value
	}
	if value, ok := sc.mutation.AgeUint(); ok {
		_spec.SetField(student.FieldAgeUint, field.TypeUint, value)
		_node.AgeUint = value
	}
	if value, ok := sc.mutation.WeightFloat(); ok {
		_spec.SetField(student.FieldWeightFloat, field.TypeFloat64, value)
		_node.WeightFloat = value
	}
	if value, ok := sc.mutation.WeightFloat32(); ok {
		_spec.SetField(student.FieldWeightFloat32, field.TypeFloat32, value)
		_node.WeightFloat32 = value
	}
	if value, ok := sc.mutation.ClassID(); ok {
		_spec.SetField(student.FieldClassID, field.TypeUUID, value)
		_node.ClassID = value
	}
	if value, ok := sc.mutation.EnrollAt(); ok {
		_spec.SetField(student.FieldEnrollAt, field.TypeTime, value)
		_node.EnrollAt = value
	}
	if value, ok := sc.mutation.StatusBool(); ok {
		_spec.SetField(student.FieldStatusBool, field.TypeBool, value)
		_node.StatusBool = value
	}
	return _node, _spec
}

// StudentCreateBulk is the builder for creating many Student entities in bulk.
type StudentCreateBulk struct {
	config
	builders []*StudentCreate
}

// Save creates the Student entities in the database.
func (scb *StudentCreateBulk) Save(ctx context.Context) ([]*Student, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Student, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StudentCreateBulk) SaveX(ctx context.Context) []*Student {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StudentCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StudentCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
