// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/entityattribute"
)

// EntityAttributeCreate is the builder for creating a EntityAttribute entity.
type EntityAttributeCreate struct {
	config
	mutation *EntityAttributeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (eac *EntityAttributeCreate) SetCreatedAt(t time.Time) *EntityAttributeCreate {
	eac.mutation.SetCreatedAt(t)
	return eac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eac *EntityAttributeCreate) SetNillableCreatedAt(t *time.Time) *EntityAttributeCreate {
	if t != nil {
		eac.SetCreatedAt(*t)
	}
	return eac
}

// SetUpdatedAt sets the "updated_at" field.
func (eac *EntityAttributeCreate) SetUpdatedAt(t time.Time) *EntityAttributeCreate {
	eac.mutation.SetUpdatedAt(t)
	return eac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (eac *EntityAttributeCreate) SetNillableUpdatedAt(t *time.Time) *EntityAttributeCreate {
	if t != nil {
		eac.SetUpdatedAt(*t)
	}
	return eac
}

// SetAttributeID sets the "attribute_id" field.
func (eac *EntityAttributeCreate) SetAttributeID(u uint64) *EntityAttributeCreate {
	eac.mutation.SetAttributeID(u)
	return eac
}

// SetEntityID sets the "entity_id" field.
func (eac *EntityAttributeCreate) SetEntityID(u uint64) *EntityAttributeCreate {
	eac.mutation.SetEntityID(u)
	return eac
}

// SetAttributeSetID sets the "attribute_set_id" field.
func (eac *EntityAttributeCreate) SetAttributeSetID(u uint64) *EntityAttributeCreate {
	eac.mutation.SetAttributeSetID(u)
	return eac
}

// SetAttributeGroupID sets the "attribute_group_id" field.
func (eac *EntityAttributeCreate) SetAttributeGroupID(u uint64) *EntityAttributeCreate {
	eac.mutation.SetAttributeGroupID(u)
	return eac
}

// SetSequence sets the "sequence" field.
func (eac *EntityAttributeCreate) SetSequence(u uint8) *EntityAttributeCreate {
	eac.mutation.SetSequence(u)
	return eac
}

// SetID sets the "id" field.
func (eac *EntityAttributeCreate) SetID(u uint64) *EntityAttributeCreate {
	eac.mutation.SetID(u)
	return eac
}

// Mutation returns the EntityAttributeMutation object of the builder.
func (eac *EntityAttributeCreate) Mutation() *EntityAttributeMutation {
	return eac.mutation
}

// Save creates the EntityAttribute in the database.
func (eac *EntityAttributeCreate) Save(ctx context.Context) (*EntityAttribute, error) {
	eac.defaults()
	return withHooks(ctx, eac.sqlSave, eac.mutation, eac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (eac *EntityAttributeCreate) SaveX(ctx context.Context) *EntityAttribute {
	v, err := eac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eac *EntityAttributeCreate) Exec(ctx context.Context) error {
	_, err := eac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eac *EntityAttributeCreate) ExecX(ctx context.Context) {
	if err := eac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eac *EntityAttributeCreate) defaults() {
	if _, ok := eac.mutation.CreatedAt(); !ok {
		v := entityattribute.DefaultCreatedAt()
		eac.mutation.SetCreatedAt(v)
	}
	if _, ok := eac.mutation.UpdatedAt(); !ok {
		v := entityattribute.DefaultUpdatedAt()
		eac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eac *EntityAttributeCreate) check() error {
	if _, ok := eac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "EntityAttribute.created_at"`)}
	}
	if _, ok := eac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "EntityAttribute.updated_at"`)}
	}
	if _, ok := eac.mutation.AttributeID(); !ok {
		return &ValidationError{Name: "attribute_id", err: errors.New(`ent: missing required field "EntityAttribute.attribute_id"`)}
	}
	if _, ok := eac.mutation.EntityID(); !ok {
		return &ValidationError{Name: "entity_id", err: errors.New(`ent: missing required field "EntityAttribute.entity_id"`)}
	}
	if _, ok := eac.mutation.AttributeSetID(); !ok {
		return &ValidationError{Name: "attribute_set_id", err: errors.New(`ent: missing required field "EntityAttribute.attribute_set_id"`)}
	}
	if _, ok := eac.mutation.AttributeGroupID(); !ok {
		return &ValidationError{Name: "attribute_group_id", err: errors.New(`ent: missing required field "EntityAttribute.attribute_group_id"`)}
	}
	if _, ok := eac.mutation.Sequence(); !ok {
		return &ValidationError{Name: "sequence", err: errors.New(`ent: missing required field "EntityAttribute.sequence"`)}
	}
	return nil
}

func (eac *EntityAttributeCreate) sqlSave(ctx context.Context) (*EntityAttribute, error) {
	if err := eac.check(); err != nil {
		return nil, err
	}
	_node, _spec := eac.createSpec()
	if err := sqlgraph.CreateNode(ctx, eac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	eac.mutation.id = &_node.ID
	eac.mutation.done = true
	return _node, nil
}

func (eac *EntityAttributeCreate) createSpec() (*EntityAttribute, *sqlgraph.CreateSpec) {
	var (
		_node = &EntityAttribute{config: eac.config}
		_spec = sqlgraph.NewCreateSpec(entityattribute.Table, sqlgraph.NewFieldSpec(entityattribute.FieldID, field.TypeUint64))
	)
	if id, ok := eac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := eac.mutation.CreatedAt(); ok {
		_spec.SetField(entityattribute.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := eac.mutation.UpdatedAt(); ok {
		_spec.SetField(entityattribute.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := eac.mutation.AttributeID(); ok {
		_spec.SetField(entityattribute.FieldAttributeID, field.TypeUint64, value)
		_node.AttributeID = value
	}
	if value, ok := eac.mutation.EntityID(); ok {
		_spec.SetField(entityattribute.FieldEntityID, field.TypeUint64, value)
		_node.EntityID = value
	}
	if value, ok := eac.mutation.AttributeSetID(); ok {
		_spec.SetField(entityattribute.FieldAttributeSetID, field.TypeUint64, value)
		_node.AttributeSetID = value
	}
	if value, ok := eac.mutation.AttributeGroupID(); ok {
		_spec.SetField(entityattribute.FieldAttributeGroupID, field.TypeUint64, value)
		_node.AttributeGroupID = value
	}
	if value, ok := eac.mutation.Sequence(); ok {
		_spec.SetField(entityattribute.FieldSequence, field.TypeUint8, value)
		_node.Sequence = value
	}
	return _node, _spec
}

// EntityAttributeCreateBulk is the builder for creating many EntityAttribute entities in bulk.
type EntityAttributeCreateBulk struct {
	config
	builders []*EntityAttributeCreate
}

// Save creates the EntityAttribute entities in the database.
func (eacb *EntityAttributeCreateBulk) Save(ctx context.Context) ([]*EntityAttribute, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eacb.builders))
	nodes := make([]*EntityAttribute, len(eacb.builders))
	mutators := make([]Mutator, len(eacb.builders))
	for i := range eacb.builders {
		func(i int, root context.Context) {
			builder := eacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EntityAttributeMutation)
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
					_, err = mutators[i+1].Mutate(root, eacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, eacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eacb *EntityAttributeCreateBulk) SaveX(ctx context.Context) []*EntityAttribute {
	v, err := eacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eacb *EntityAttributeCreateBulk) Exec(ctx context.Context) error {
	_, err := eacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eacb *EntityAttributeCreateBulk) ExecX(ctx context.Context) {
	if err := eacb.Exec(ctx); err != nil {
		panic(err)
	}
}
