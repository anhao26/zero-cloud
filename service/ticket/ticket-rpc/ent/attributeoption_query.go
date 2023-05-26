// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attributeoption"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/predicate"
)

// AttributeOptionQuery is the builder for querying AttributeOption entities.
type AttributeOptionQuery struct {
	config
	ctx        *QueryContext
	order      []attributeoption.OrderOption
	inters     []Interceptor
	predicates []predicate.AttributeOption
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttributeOptionQuery builder.
func (aoq *AttributeOptionQuery) Where(ps ...predicate.AttributeOption) *AttributeOptionQuery {
	aoq.predicates = append(aoq.predicates, ps...)
	return aoq
}

// Limit the number of records to be returned by this query.
func (aoq *AttributeOptionQuery) Limit(limit int) *AttributeOptionQuery {
	aoq.ctx.Limit = &limit
	return aoq
}

// Offset to start from.
func (aoq *AttributeOptionQuery) Offset(offset int) *AttributeOptionQuery {
	aoq.ctx.Offset = &offset
	return aoq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aoq *AttributeOptionQuery) Unique(unique bool) *AttributeOptionQuery {
	aoq.ctx.Unique = &unique
	return aoq
}

// Order specifies how the records should be ordered.
func (aoq *AttributeOptionQuery) Order(o ...attributeoption.OrderOption) *AttributeOptionQuery {
	aoq.order = append(aoq.order, o...)
	return aoq
}

// First returns the first AttributeOption entity from the query.
// Returns a *NotFoundError when no AttributeOption was found.
func (aoq *AttributeOptionQuery) First(ctx context.Context) (*AttributeOption, error) {
	nodes, err := aoq.Limit(1).All(setContextOp(ctx, aoq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attributeoption.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aoq *AttributeOptionQuery) FirstX(ctx context.Context) *AttributeOption {
	node, err := aoq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AttributeOption ID from the query.
// Returns a *NotFoundError when no AttributeOption ID was found.
func (aoq *AttributeOptionQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = aoq.Limit(1).IDs(setContextOp(ctx, aoq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attributeoption.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aoq *AttributeOptionQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := aoq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AttributeOption entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AttributeOption entity is found.
// Returns a *NotFoundError when no AttributeOption entities are found.
func (aoq *AttributeOptionQuery) Only(ctx context.Context) (*AttributeOption, error) {
	nodes, err := aoq.Limit(2).All(setContextOp(ctx, aoq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attributeoption.Label}
	default:
		return nil, &NotSingularError{attributeoption.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aoq *AttributeOptionQuery) OnlyX(ctx context.Context) *AttributeOption {
	node, err := aoq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AttributeOption ID in the query.
// Returns a *NotSingularError when more than one AttributeOption ID is found.
// Returns a *NotFoundError when no entities are found.
func (aoq *AttributeOptionQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = aoq.Limit(2).IDs(setContextOp(ctx, aoq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attributeoption.Label}
	default:
		err = &NotSingularError{attributeoption.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aoq *AttributeOptionQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := aoq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AttributeOptions.
func (aoq *AttributeOptionQuery) All(ctx context.Context) ([]*AttributeOption, error) {
	ctx = setContextOp(ctx, aoq.ctx, "All")
	if err := aoq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AttributeOption, *AttributeOptionQuery]()
	return withInterceptors[[]*AttributeOption](ctx, aoq, qr, aoq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aoq *AttributeOptionQuery) AllX(ctx context.Context) []*AttributeOption {
	nodes, err := aoq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AttributeOption IDs.
func (aoq *AttributeOptionQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if aoq.ctx.Unique == nil && aoq.path != nil {
		aoq.Unique(true)
	}
	ctx = setContextOp(ctx, aoq.ctx, "IDs")
	if err = aoq.Select(attributeoption.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aoq *AttributeOptionQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := aoq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aoq *AttributeOptionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aoq.ctx, "Count")
	if err := aoq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aoq, querierCount[*AttributeOptionQuery](), aoq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aoq *AttributeOptionQuery) CountX(ctx context.Context) int {
	count, err := aoq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aoq *AttributeOptionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aoq.ctx, "Exist")
	switch _, err := aoq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aoq *AttributeOptionQuery) ExistX(ctx context.Context) bool {
	exist, err := aoq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttributeOptionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aoq *AttributeOptionQuery) Clone() *AttributeOptionQuery {
	if aoq == nil {
		return nil
	}
	return &AttributeOptionQuery{
		config:     aoq.config,
		ctx:        aoq.ctx.Clone(),
		order:      append([]attributeoption.OrderOption{}, aoq.order...),
		inters:     append([]Interceptor{}, aoq.inters...),
		predicates: append([]predicate.AttributeOption{}, aoq.predicates...),
		// clone intermediate query.
		sql:  aoq.sql.Clone(),
		path: aoq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AttributeOption.Query().
//		GroupBy(attributeoption.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aoq *AttributeOptionQuery) GroupBy(field string, fields ...string) *AttributeOptionGroupBy {
	aoq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AttributeOptionGroupBy{build: aoq}
	grbuild.flds = &aoq.ctx.Fields
	grbuild.label = attributeoption.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.AttributeOption.Query().
//		Select(attributeoption.FieldCreatedAt).
//		Scan(ctx, &v)
func (aoq *AttributeOptionQuery) Select(fields ...string) *AttributeOptionSelect {
	aoq.ctx.Fields = append(aoq.ctx.Fields, fields...)
	sbuild := &AttributeOptionSelect{AttributeOptionQuery: aoq}
	sbuild.label = attributeoption.Label
	sbuild.flds, sbuild.scan = &aoq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AttributeOptionSelect configured with the given aggregations.
func (aoq *AttributeOptionQuery) Aggregate(fns ...AggregateFunc) *AttributeOptionSelect {
	return aoq.Select().Aggregate(fns...)
}

func (aoq *AttributeOptionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aoq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aoq); err != nil {
				return err
			}
		}
	}
	for _, f := range aoq.ctx.Fields {
		if !attributeoption.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aoq.path != nil {
		prev, err := aoq.path(ctx)
		if err != nil {
			return err
		}
		aoq.sql = prev
	}
	return nil
}

func (aoq *AttributeOptionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AttributeOption, error) {
	var (
		nodes   = []*AttributeOption{}
		withFKs = aoq.withFKs
		_spec   = aoq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, attributeoption.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AttributeOption).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AttributeOption{config: aoq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aoq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aoq *AttributeOptionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aoq.querySpec()
	_spec.Node.Columns = aoq.ctx.Fields
	if len(aoq.ctx.Fields) > 0 {
		_spec.Unique = aoq.ctx.Unique != nil && *aoq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aoq.driver, _spec)
}

func (aoq *AttributeOptionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(attributeoption.Table, attributeoption.Columns, sqlgraph.NewFieldSpec(attributeoption.FieldID, field.TypeUint64))
	_spec.From = aoq.sql
	if unique := aoq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aoq.path != nil {
		_spec.Unique = true
	}
	if fields := aoq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attributeoption.FieldID)
		for i := range fields {
			if fields[i] != attributeoption.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aoq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aoq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aoq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aoq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aoq *AttributeOptionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aoq.driver.Dialect())
	t1 := builder.Table(attributeoption.Table)
	columns := aoq.ctx.Fields
	if len(columns) == 0 {
		columns = attributeoption.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aoq.sql != nil {
		selector = aoq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aoq.ctx.Unique != nil && *aoq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aoq.predicates {
		p(selector)
	}
	for _, p := range aoq.order {
		p(selector)
	}
	if offset := aoq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aoq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AttributeOptionGroupBy is the group-by builder for AttributeOption entities.
type AttributeOptionGroupBy struct {
	selector
	build *AttributeOptionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aogb *AttributeOptionGroupBy) Aggregate(fns ...AggregateFunc) *AttributeOptionGroupBy {
	aogb.fns = append(aogb.fns, fns...)
	return aogb
}

// Scan applies the selector query and scans the result into the given value.
func (aogb *AttributeOptionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aogb.build.ctx, "GroupBy")
	if err := aogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeOptionQuery, *AttributeOptionGroupBy](ctx, aogb.build, aogb, aogb.build.inters, v)
}

func (aogb *AttributeOptionGroupBy) sqlScan(ctx context.Context, root *AttributeOptionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(aogb.fns))
	for _, fn := range aogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*aogb.flds)+len(aogb.fns))
		for _, f := range *aogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*aogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AttributeOptionSelect is the builder for selecting fields of AttributeOption entities.
type AttributeOptionSelect struct {
	*AttributeOptionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aos *AttributeOptionSelect) Aggregate(fns ...AggregateFunc) *AttributeOptionSelect {
	aos.fns = append(aos.fns, fns...)
	return aos
}

// Scan applies the selector query and scans the result into the given value.
func (aos *AttributeOptionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aos.ctx, "Select")
	if err := aos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeOptionQuery, *AttributeOptionSelect](ctx, aos.AttributeOptionQuery, aos, aos.inters, v)
}

func (aos *AttributeOptionSelect) sqlScan(ctx context.Context, root *AttributeOptionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aos.fns))
	for _, fn := range aos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
