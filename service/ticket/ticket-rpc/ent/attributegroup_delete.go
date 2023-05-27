// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attributegroup"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/predicate"
)

// AttributeGroupDelete is the builder for deleting a AttributeGroup entity.
type AttributeGroupDelete struct {
	config
	hooks    []Hook
	mutation *AttributeGroupMutation
}

// Where appends a list predicates to the AttributeGroupDelete builder.
func (agd *AttributeGroupDelete) Where(ps ...predicate.AttributeGroup) *AttributeGroupDelete {
	agd.mutation.Where(ps...)
	return agd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (agd *AttributeGroupDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, agd.sqlExec, agd.mutation, agd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (agd *AttributeGroupDelete) ExecX(ctx context.Context) int {
	n, err := agd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (agd *AttributeGroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(attributegroup.Table, sqlgraph.NewFieldSpec(attributegroup.FieldID, field.TypeUint64))
	if ps := agd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, agd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	agd.mutation.done = true
	return affected, err
}

// AttributeGroupDeleteOne is the builder for deleting a single AttributeGroup entity.
type AttributeGroupDeleteOne struct {
	agd *AttributeGroupDelete
}

// Where appends a list predicates to the AttributeGroupDelete builder.
func (agdo *AttributeGroupDeleteOne) Where(ps ...predicate.AttributeGroup) *AttributeGroupDeleteOne {
	agdo.agd.mutation.Where(ps...)
	return agdo
}

// Exec executes the deletion query.
func (agdo *AttributeGroupDeleteOne) Exec(ctx context.Context) error {
	n, err := agdo.agd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{attributegroup.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (agdo *AttributeGroupDeleteOne) ExecX(ctx context.Context) {
	if err := agdo.Exec(ctx); err != nil {
		panic(err)
	}
}