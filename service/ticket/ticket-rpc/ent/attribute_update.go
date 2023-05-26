// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attribute"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attributeoption"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/entity"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/predicate"
)

// AttributeUpdate is the builder for updating Attribute entities.
type AttributeUpdate struct {
	config
	hooks    []Hook
	mutation *AttributeMutation
}

// Where appends a list predicates to the AttributeUpdate builder.
func (au *AttributeUpdate) Where(ps ...predicate.Attribute) *AttributeUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AttributeUpdate) SetUpdatedAt(t time.Time) *AttributeUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetEntityID sets the "entity_id" field.
func (au *AttributeUpdate) SetEntityID(u uint64) *AttributeUpdate {
	au.mutation.ResetEntityID()
	au.mutation.SetEntityID(u)
	return au
}

// AddEntityID adds u to the "entity_id" field.
func (au *AttributeUpdate) AddEntityID(u int64) *AttributeUpdate {
	au.mutation.AddEntityID(u)
	return au
}

// SetAttributeCode sets the "attribute_code" field.
func (au *AttributeUpdate) SetAttributeCode(s string) *AttributeUpdate {
	au.mutation.SetAttributeCode(s)
	return au
}

// SetBackendClass sets the "backend_class" field.
func (au *AttributeUpdate) SetBackendClass(s string) *AttributeUpdate {
	au.mutation.SetBackendClass(s)
	return au
}

// SetBackendType sets the "backend_type" field.
func (au *AttributeUpdate) SetBackendType(s string) *AttributeUpdate {
	au.mutation.SetBackendType(s)
	return au
}

// SetBackendTable sets the "backend_table" field.
func (au *AttributeUpdate) SetBackendTable(s string) *AttributeUpdate {
	au.mutation.SetBackendTable(s)
	return au
}

// SetFrontendClass sets the "frontend_class" field.
func (au *AttributeUpdate) SetFrontendClass(s string) *AttributeUpdate {
	au.mutation.SetFrontendClass(s)
	return au
}

// SetFrontendType sets the "frontend_type" field.
func (au *AttributeUpdate) SetFrontendType(s string) *AttributeUpdate {
	au.mutation.SetFrontendType(s)
	return au
}

// SetFrontendLabel sets the "frontend_label" field.
func (au *AttributeUpdate) SetFrontendLabel(s string) *AttributeUpdate {
	au.mutation.SetFrontendLabel(s)
	return au
}

// SetSourceClass sets the "source_class" field.
func (au *AttributeUpdate) SetSourceClass(s string) *AttributeUpdate {
	au.mutation.SetSourceClass(s)
	return au
}

// SetDefaultValue sets the "default_value" field.
func (au *AttributeUpdate) SetDefaultValue(s string) *AttributeUpdate {
	au.mutation.SetDefaultValue(s)
	return au
}

// SetIsFilterable sets the "is_filterable" field.
func (au *AttributeUpdate) SetIsFilterable(u uint8) *AttributeUpdate {
	au.mutation.ResetIsFilterable()
	au.mutation.SetIsFilterable(u)
	return au
}

// SetNillableIsFilterable sets the "is_filterable" field if the given value is not nil.
func (au *AttributeUpdate) SetNillableIsFilterable(u *uint8) *AttributeUpdate {
	if u != nil {
		au.SetIsFilterable(*u)
	}
	return au
}

// AddIsFilterable adds u to the "is_filterable" field.
func (au *AttributeUpdate) AddIsFilterable(u int8) *AttributeUpdate {
	au.mutation.AddIsFilterable(u)
	return au
}

// SetIsSearchable sets the "is_searchable" field.
func (au *AttributeUpdate) SetIsSearchable(u uint8) *AttributeUpdate {
	au.mutation.ResetIsSearchable()
	au.mutation.SetIsSearchable(u)
	return au
}

// SetNillableIsSearchable sets the "is_searchable" field if the given value is not nil.
func (au *AttributeUpdate) SetNillableIsSearchable(u *uint8) *AttributeUpdate {
	if u != nil {
		au.SetIsSearchable(*u)
	}
	return au
}

// AddIsSearchable adds u to the "is_searchable" field.
func (au *AttributeUpdate) AddIsSearchable(u int8) *AttributeUpdate {
	au.mutation.AddIsSearchable(u)
	return au
}

// SetIsRequired sets the "is_required" field.
func (au *AttributeUpdate) SetIsRequired(u uint8) *AttributeUpdate {
	au.mutation.ResetIsRequired()
	au.mutation.SetIsRequired(u)
	return au
}

// SetNillableIsRequired sets the "is_required" field if the given value is not nil.
func (au *AttributeUpdate) SetNillableIsRequired(u *uint8) *AttributeUpdate {
	if u != nil {
		au.SetIsRequired(*u)
	}
	return au
}

// AddIsRequired adds u to the "is_required" field.
func (au *AttributeUpdate) AddIsRequired(u int8) *AttributeUpdate {
	au.mutation.AddIsRequired(u)
	return au
}

// SetRequiredValidateClass sets the "required_validate_class" field.
func (au *AttributeUpdate) SetRequiredValidateClass(s string) *AttributeUpdate {
	au.mutation.SetRequiredValidateClass(s)
	return au
}

// AddEntityIDs adds the "entities" edge to the Entity entity by IDs.
func (au *AttributeUpdate) AddEntityIDs(ids ...uint64) *AttributeUpdate {
	au.mutation.AddEntityIDs(ids...)
	return au
}

// AddEntities adds the "entities" edges to the Entity entity.
func (au *AttributeUpdate) AddEntities(e ...*Entity) *AttributeUpdate {
	ids := make([]uint64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return au.AddEntityIDs(ids...)
}

// AddAttributeOptionIDs adds the "attribute_options" edge to the AttributeOption entity by IDs.
func (au *AttributeUpdate) AddAttributeOptionIDs(ids ...uint64) *AttributeUpdate {
	au.mutation.AddAttributeOptionIDs(ids...)
	return au
}

// AddAttributeOptions adds the "attribute_options" edges to the AttributeOption entity.
func (au *AttributeUpdate) AddAttributeOptions(a ...*AttributeOption) *AttributeUpdate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddAttributeOptionIDs(ids...)
}

// Mutation returns the AttributeMutation object of the builder.
func (au *AttributeUpdate) Mutation() *AttributeMutation {
	return au.mutation
}

// ClearEntities clears all "entities" edges to the Entity entity.
func (au *AttributeUpdate) ClearEntities() *AttributeUpdate {
	au.mutation.ClearEntities()
	return au
}

// RemoveEntityIDs removes the "entities" edge to Entity entities by IDs.
func (au *AttributeUpdate) RemoveEntityIDs(ids ...uint64) *AttributeUpdate {
	au.mutation.RemoveEntityIDs(ids...)
	return au
}

// RemoveEntities removes "entities" edges to Entity entities.
func (au *AttributeUpdate) RemoveEntities(e ...*Entity) *AttributeUpdate {
	ids := make([]uint64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return au.RemoveEntityIDs(ids...)
}

// ClearAttributeOptions clears all "attribute_options" edges to the AttributeOption entity.
func (au *AttributeUpdate) ClearAttributeOptions() *AttributeUpdate {
	au.mutation.ClearAttributeOptions()
	return au
}

// RemoveAttributeOptionIDs removes the "attribute_options" edge to AttributeOption entities by IDs.
func (au *AttributeUpdate) RemoveAttributeOptionIDs(ids ...uint64) *AttributeUpdate {
	au.mutation.RemoveAttributeOptionIDs(ids...)
	return au
}

// RemoveAttributeOptions removes "attribute_options" edges to AttributeOption entities.
func (au *AttributeUpdate) RemoveAttributeOptions(a ...*AttributeOption) *AttributeUpdate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveAttributeOptionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AttributeUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AttributeUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AttributeUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AttributeUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AttributeUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := attribute.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

func (au *AttributeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(attribute.Table, attribute.Columns, sqlgraph.NewFieldSpec(attribute.FieldID, field.TypeUint64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(attribute.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.EntityID(); ok {
		_spec.SetField(attribute.FieldEntityID, field.TypeUint64, value)
	}
	if value, ok := au.mutation.AddedEntityID(); ok {
		_spec.AddField(attribute.FieldEntityID, field.TypeUint64, value)
	}
	if value, ok := au.mutation.AttributeCode(); ok {
		_spec.SetField(attribute.FieldAttributeCode, field.TypeString, value)
	}
	if value, ok := au.mutation.BackendClass(); ok {
		_spec.SetField(attribute.FieldBackendClass, field.TypeString, value)
	}
	if value, ok := au.mutation.BackendType(); ok {
		_spec.SetField(attribute.FieldBackendType, field.TypeString, value)
	}
	if value, ok := au.mutation.BackendTable(); ok {
		_spec.SetField(attribute.FieldBackendTable, field.TypeString, value)
	}
	if value, ok := au.mutation.FrontendClass(); ok {
		_spec.SetField(attribute.FieldFrontendClass, field.TypeString, value)
	}
	if value, ok := au.mutation.FrontendType(); ok {
		_spec.SetField(attribute.FieldFrontendType, field.TypeString, value)
	}
	if value, ok := au.mutation.FrontendLabel(); ok {
		_spec.SetField(attribute.FieldFrontendLabel, field.TypeString, value)
	}
	if value, ok := au.mutation.SourceClass(); ok {
		_spec.SetField(attribute.FieldSourceClass, field.TypeString, value)
	}
	if value, ok := au.mutation.DefaultValue(); ok {
		_spec.SetField(attribute.FieldDefaultValue, field.TypeString, value)
	}
	if value, ok := au.mutation.IsFilterable(); ok {
		_spec.SetField(attribute.FieldIsFilterable, field.TypeUint8, value)
	}
	if value, ok := au.mutation.AddedIsFilterable(); ok {
		_spec.AddField(attribute.FieldIsFilterable, field.TypeUint8, value)
	}
	if value, ok := au.mutation.IsSearchable(); ok {
		_spec.SetField(attribute.FieldIsSearchable, field.TypeUint8, value)
	}
	if value, ok := au.mutation.AddedIsSearchable(); ok {
		_spec.AddField(attribute.FieldIsSearchable, field.TypeUint8, value)
	}
	if value, ok := au.mutation.IsRequired(); ok {
		_spec.SetField(attribute.FieldIsRequired, field.TypeUint8, value)
	}
	if value, ok := au.mutation.AddedIsRequired(); ok {
		_spec.AddField(attribute.FieldIsRequired, field.TypeUint8, value)
	}
	if value, ok := au.mutation.RequiredValidateClass(); ok {
		_spec.SetField(attribute.FieldRequiredValidateClass, field.TypeString, value)
	}
	if au.mutation.EntitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.EntitiesTable,
			Columns: []string{attribute.EntitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entity.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedEntitiesIDs(); len(nodes) > 0 && !au.mutation.EntitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.EntitiesTable,
			Columns: []string{attribute.EntitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entity.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.EntitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.EntitiesTable,
			Columns: []string{attribute.EntitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entity.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.AttributeOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AttributeOptionsTable,
			Columns: []string{attribute.AttributeOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attributeoption.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedAttributeOptionsIDs(); len(nodes) > 0 && !au.mutation.AttributeOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AttributeOptionsTable,
			Columns: []string{attribute.AttributeOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attributeoption.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AttributeOptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AttributeOptionsTable,
			Columns: []string{attribute.AttributeOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attributeoption.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attribute.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AttributeUpdateOne is the builder for updating a single Attribute entity.
type AttributeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttributeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AttributeUpdateOne) SetUpdatedAt(t time.Time) *AttributeUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetEntityID sets the "entity_id" field.
func (auo *AttributeUpdateOne) SetEntityID(u uint64) *AttributeUpdateOne {
	auo.mutation.ResetEntityID()
	auo.mutation.SetEntityID(u)
	return auo
}

// AddEntityID adds u to the "entity_id" field.
func (auo *AttributeUpdateOne) AddEntityID(u int64) *AttributeUpdateOne {
	auo.mutation.AddEntityID(u)
	return auo
}

// SetAttributeCode sets the "attribute_code" field.
func (auo *AttributeUpdateOne) SetAttributeCode(s string) *AttributeUpdateOne {
	auo.mutation.SetAttributeCode(s)
	return auo
}

// SetBackendClass sets the "backend_class" field.
func (auo *AttributeUpdateOne) SetBackendClass(s string) *AttributeUpdateOne {
	auo.mutation.SetBackendClass(s)
	return auo
}

// SetBackendType sets the "backend_type" field.
func (auo *AttributeUpdateOne) SetBackendType(s string) *AttributeUpdateOne {
	auo.mutation.SetBackendType(s)
	return auo
}

// SetBackendTable sets the "backend_table" field.
func (auo *AttributeUpdateOne) SetBackendTable(s string) *AttributeUpdateOne {
	auo.mutation.SetBackendTable(s)
	return auo
}

// SetFrontendClass sets the "frontend_class" field.
func (auo *AttributeUpdateOne) SetFrontendClass(s string) *AttributeUpdateOne {
	auo.mutation.SetFrontendClass(s)
	return auo
}

// SetFrontendType sets the "frontend_type" field.
func (auo *AttributeUpdateOne) SetFrontendType(s string) *AttributeUpdateOne {
	auo.mutation.SetFrontendType(s)
	return auo
}

// SetFrontendLabel sets the "frontend_label" field.
func (auo *AttributeUpdateOne) SetFrontendLabel(s string) *AttributeUpdateOne {
	auo.mutation.SetFrontendLabel(s)
	return auo
}

// SetSourceClass sets the "source_class" field.
func (auo *AttributeUpdateOne) SetSourceClass(s string) *AttributeUpdateOne {
	auo.mutation.SetSourceClass(s)
	return auo
}

// SetDefaultValue sets the "default_value" field.
func (auo *AttributeUpdateOne) SetDefaultValue(s string) *AttributeUpdateOne {
	auo.mutation.SetDefaultValue(s)
	return auo
}

// SetIsFilterable sets the "is_filterable" field.
func (auo *AttributeUpdateOne) SetIsFilterable(u uint8) *AttributeUpdateOne {
	auo.mutation.ResetIsFilterable()
	auo.mutation.SetIsFilterable(u)
	return auo
}

// SetNillableIsFilterable sets the "is_filterable" field if the given value is not nil.
func (auo *AttributeUpdateOne) SetNillableIsFilterable(u *uint8) *AttributeUpdateOne {
	if u != nil {
		auo.SetIsFilterable(*u)
	}
	return auo
}

// AddIsFilterable adds u to the "is_filterable" field.
func (auo *AttributeUpdateOne) AddIsFilterable(u int8) *AttributeUpdateOne {
	auo.mutation.AddIsFilterable(u)
	return auo
}

// SetIsSearchable sets the "is_searchable" field.
func (auo *AttributeUpdateOne) SetIsSearchable(u uint8) *AttributeUpdateOne {
	auo.mutation.ResetIsSearchable()
	auo.mutation.SetIsSearchable(u)
	return auo
}

// SetNillableIsSearchable sets the "is_searchable" field if the given value is not nil.
func (auo *AttributeUpdateOne) SetNillableIsSearchable(u *uint8) *AttributeUpdateOne {
	if u != nil {
		auo.SetIsSearchable(*u)
	}
	return auo
}

// AddIsSearchable adds u to the "is_searchable" field.
func (auo *AttributeUpdateOne) AddIsSearchable(u int8) *AttributeUpdateOne {
	auo.mutation.AddIsSearchable(u)
	return auo
}

// SetIsRequired sets the "is_required" field.
func (auo *AttributeUpdateOne) SetIsRequired(u uint8) *AttributeUpdateOne {
	auo.mutation.ResetIsRequired()
	auo.mutation.SetIsRequired(u)
	return auo
}

// SetNillableIsRequired sets the "is_required" field if the given value is not nil.
func (auo *AttributeUpdateOne) SetNillableIsRequired(u *uint8) *AttributeUpdateOne {
	if u != nil {
		auo.SetIsRequired(*u)
	}
	return auo
}

// AddIsRequired adds u to the "is_required" field.
func (auo *AttributeUpdateOne) AddIsRequired(u int8) *AttributeUpdateOne {
	auo.mutation.AddIsRequired(u)
	return auo
}

// SetRequiredValidateClass sets the "required_validate_class" field.
func (auo *AttributeUpdateOne) SetRequiredValidateClass(s string) *AttributeUpdateOne {
	auo.mutation.SetRequiredValidateClass(s)
	return auo
}

// AddEntityIDs adds the "entities" edge to the Entity entity by IDs.
func (auo *AttributeUpdateOne) AddEntityIDs(ids ...uint64) *AttributeUpdateOne {
	auo.mutation.AddEntityIDs(ids...)
	return auo
}

// AddEntities adds the "entities" edges to the Entity entity.
func (auo *AttributeUpdateOne) AddEntities(e ...*Entity) *AttributeUpdateOne {
	ids := make([]uint64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return auo.AddEntityIDs(ids...)
}

// AddAttributeOptionIDs adds the "attribute_options" edge to the AttributeOption entity by IDs.
func (auo *AttributeUpdateOne) AddAttributeOptionIDs(ids ...uint64) *AttributeUpdateOne {
	auo.mutation.AddAttributeOptionIDs(ids...)
	return auo
}

// AddAttributeOptions adds the "attribute_options" edges to the AttributeOption entity.
func (auo *AttributeUpdateOne) AddAttributeOptions(a ...*AttributeOption) *AttributeUpdateOne {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddAttributeOptionIDs(ids...)
}

// Mutation returns the AttributeMutation object of the builder.
func (auo *AttributeUpdateOne) Mutation() *AttributeMutation {
	return auo.mutation
}

// ClearEntities clears all "entities" edges to the Entity entity.
func (auo *AttributeUpdateOne) ClearEntities() *AttributeUpdateOne {
	auo.mutation.ClearEntities()
	return auo
}

// RemoveEntityIDs removes the "entities" edge to Entity entities by IDs.
func (auo *AttributeUpdateOne) RemoveEntityIDs(ids ...uint64) *AttributeUpdateOne {
	auo.mutation.RemoveEntityIDs(ids...)
	return auo
}

// RemoveEntities removes "entities" edges to Entity entities.
func (auo *AttributeUpdateOne) RemoveEntities(e ...*Entity) *AttributeUpdateOne {
	ids := make([]uint64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return auo.RemoveEntityIDs(ids...)
}

// ClearAttributeOptions clears all "attribute_options" edges to the AttributeOption entity.
func (auo *AttributeUpdateOne) ClearAttributeOptions() *AttributeUpdateOne {
	auo.mutation.ClearAttributeOptions()
	return auo
}

// RemoveAttributeOptionIDs removes the "attribute_options" edge to AttributeOption entities by IDs.
func (auo *AttributeUpdateOne) RemoveAttributeOptionIDs(ids ...uint64) *AttributeUpdateOne {
	auo.mutation.RemoveAttributeOptionIDs(ids...)
	return auo
}

// RemoveAttributeOptions removes "attribute_options" edges to AttributeOption entities.
func (auo *AttributeUpdateOne) RemoveAttributeOptions(a ...*AttributeOption) *AttributeUpdateOne {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveAttributeOptionIDs(ids...)
}

// Where appends a list predicates to the AttributeUpdate builder.
func (auo *AttributeUpdateOne) Where(ps ...predicate.Attribute) *AttributeUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AttributeUpdateOne) Select(field string, fields ...string) *AttributeUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Attribute entity.
func (auo *AttributeUpdateOne) Save(ctx context.Context) (*Attribute, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AttributeUpdateOne) SaveX(ctx context.Context) *Attribute {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AttributeUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AttributeUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AttributeUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := attribute.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

func (auo *AttributeUpdateOne) sqlSave(ctx context.Context) (_node *Attribute, err error) {
	_spec := sqlgraph.NewUpdateSpec(attribute.Table, attribute.Columns, sqlgraph.NewFieldSpec(attribute.FieldID, field.TypeUint64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Attribute.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attribute.FieldID)
		for _, f := range fields {
			if !attribute.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != attribute.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(attribute.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.EntityID(); ok {
		_spec.SetField(attribute.FieldEntityID, field.TypeUint64, value)
	}
	if value, ok := auo.mutation.AddedEntityID(); ok {
		_spec.AddField(attribute.FieldEntityID, field.TypeUint64, value)
	}
	if value, ok := auo.mutation.AttributeCode(); ok {
		_spec.SetField(attribute.FieldAttributeCode, field.TypeString, value)
	}
	if value, ok := auo.mutation.BackendClass(); ok {
		_spec.SetField(attribute.FieldBackendClass, field.TypeString, value)
	}
	if value, ok := auo.mutation.BackendType(); ok {
		_spec.SetField(attribute.FieldBackendType, field.TypeString, value)
	}
	if value, ok := auo.mutation.BackendTable(); ok {
		_spec.SetField(attribute.FieldBackendTable, field.TypeString, value)
	}
	if value, ok := auo.mutation.FrontendClass(); ok {
		_spec.SetField(attribute.FieldFrontendClass, field.TypeString, value)
	}
	if value, ok := auo.mutation.FrontendType(); ok {
		_spec.SetField(attribute.FieldFrontendType, field.TypeString, value)
	}
	if value, ok := auo.mutation.FrontendLabel(); ok {
		_spec.SetField(attribute.FieldFrontendLabel, field.TypeString, value)
	}
	if value, ok := auo.mutation.SourceClass(); ok {
		_spec.SetField(attribute.FieldSourceClass, field.TypeString, value)
	}
	if value, ok := auo.mutation.DefaultValue(); ok {
		_spec.SetField(attribute.FieldDefaultValue, field.TypeString, value)
	}
	if value, ok := auo.mutation.IsFilterable(); ok {
		_spec.SetField(attribute.FieldIsFilterable, field.TypeUint8, value)
	}
	if value, ok := auo.mutation.AddedIsFilterable(); ok {
		_spec.AddField(attribute.FieldIsFilterable, field.TypeUint8, value)
	}
	if value, ok := auo.mutation.IsSearchable(); ok {
		_spec.SetField(attribute.FieldIsSearchable, field.TypeUint8, value)
	}
	if value, ok := auo.mutation.AddedIsSearchable(); ok {
		_spec.AddField(attribute.FieldIsSearchable, field.TypeUint8, value)
	}
	if value, ok := auo.mutation.IsRequired(); ok {
		_spec.SetField(attribute.FieldIsRequired, field.TypeUint8, value)
	}
	if value, ok := auo.mutation.AddedIsRequired(); ok {
		_spec.AddField(attribute.FieldIsRequired, field.TypeUint8, value)
	}
	if value, ok := auo.mutation.RequiredValidateClass(); ok {
		_spec.SetField(attribute.FieldRequiredValidateClass, field.TypeString, value)
	}
	if auo.mutation.EntitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.EntitiesTable,
			Columns: []string{attribute.EntitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entity.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedEntitiesIDs(); len(nodes) > 0 && !auo.mutation.EntitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.EntitiesTable,
			Columns: []string{attribute.EntitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entity.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.EntitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.EntitiesTable,
			Columns: []string{attribute.EntitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entity.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.AttributeOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AttributeOptionsTable,
			Columns: []string{attribute.AttributeOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attributeoption.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedAttributeOptionsIDs(); len(nodes) > 0 && !auo.mutation.AttributeOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AttributeOptionsTable,
			Columns: []string{attribute.AttributeOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attributeoption.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AttributeOptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AttributeOptionsTable,
			Columns: []string{attribute.AttributeOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attributeoption.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Attribute{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attribute.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
