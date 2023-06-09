// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attributeoption"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/predicate"
)

// AttributeOptionDelete is the builder for deleting a AttributeOption entity.
type AttributeOptionDelete struct {
	config
	hooks    []Hook
	mutation *AttributeOptionMutation
}

// Where appends a list predicates to the AttributeOptionDelete builder.
func (aod *AttributeOptionDelete) Where(ps ...predicate.AttributeOption) *AttributeOptionDelete {
	aod.mutation.Where(ps...)
	return aod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aod *AttributeOptionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, aod.sqlExec, aod.mutation, aod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (aod *AttributeOptionDelete) ExecX(ctx context.Context) int {
	n, err := aod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aod *AttributeOptionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(attributeoption.Table, sqlgraph.NewFieldSpec(attributeoption.FieldID, field.TypeUint64))
	if ps := aod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, aod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	aod.mutation.done = true
	return affected, err
}

// AttributeOptionDeleteOne is the builder for deleting a single AttributeOption entity.
type AttributeOptionDeleteOne struct {
	aod *AttributeOptionDelete
}

// Where appends a list predicates to the AttributeOptionDelete builder.
func (aodo *AttributeOptionDeleteOne) Where(ps ...predicate.AttributeOption) *AttributeOptionDeleteOne {
	aodo.aod.mutation.Where(ps...)
	return aodo
}

// Exec executes the deletion query.
func (aodo *AttributeOptionDeleteOne) Exec(ctx context.Context) error {
	n, err := aodo.aod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{attributeoption.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aodo *AttributeOptionDeleteOne) ExecX(ctx context.Context) {
	if err := aodo.Exec(ctx); err != nil {
		panic(err)
	}
}
