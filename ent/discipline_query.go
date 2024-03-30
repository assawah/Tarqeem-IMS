// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tarqeem/ims/ent/discipline"
	"github.com/tarqeem/ims/ent/predicate"
)

// DisciplineQuery is the builder for querying Discipline entities.
type DisciplineQuery struct {
	config
	ctx        *QueryContext
	order      []discipline.OrderOption
	inters     []Interceptor
	predicates []predicate.Discipline
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DisciplineQuery builder.
func (dq *DisciplineQuery) Where(ps ...predicate.Discipline) *DisciplineQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DisciplineQuery) Limit(limit int) *DisciplineQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DisciplineQuery) Offset(offset int) *DisciplineQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DisciplineQuery) Unique(unique bool) *DisciplineQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DisciplineQuery) Order(o ...discipline.OrderOption) *DisciplineQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// First returns the first Discipline entity from the query.
// Returns a *NotFoundError when no Discipline was found.
func (dq *DisciplineQuery) First(ctx context.Context) (*Discipline, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{discipline.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DisciplineQuery) FirstX(ctx context.Context) *Discipline {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Discipline ID from the query.
// Returns a *NotFoundError when no Discipline ID was found.
func (dq *DisciplineQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{discipline.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DisciplineQuery) FirstIDX(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Discipline entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Discipline entity is found.
// Returns a *NotFoundError when no Discipline entities are found.
func (dq *DisciplineQuery) Only(ctx context.Context) (*Discipline, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{discipline.Label}
	default:
		return nil, &NotSingularError{discipline.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DisciplineQuery) OnlyX(ctx context.Context) *Discipline {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Discipline ID in the query.
// Returns a *NotSingularError when more than one Discipline ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DisciplineQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{discipline.Label}
	default:
		err = &NotSingularError{discipline.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DisciplineQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Disciplines.
func (dq *DisciplineQuery) All(ctx context.Context) ([]*Discipline, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Discipline, *DisciplineQuery]()
	return withInterceptors[[]*Discipline](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DisciplineQuery) AllX(ctx context.Context) []*Discipline {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Discipline IDs.
func (dq *DisciplineQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err = dq.Select(discipline.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DisciplineQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DisciplineQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DisciplineQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DisciplineQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DisciplineQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DisciplineQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DisciplineQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DisciplineQuery) Clone() *DisciplineQuery {
	if dq == nil {
		return nil
	}
	return &DisciplineQuery{
		config:     dq.config,
		ctx:        dq.ctx.Clone(),
		order:      append([]discipline.OrderOption{}, dq.order...),
		inters:     append([]Interceptor{}, dq.inters...),
		predicates: append([]predicate.Discipline{}, dq.predicates...),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Discipline.Query().
//		GroupBy(discipline.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DisciplineQuery) GroupBy(field string, fields ...string) *DisciplineGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DisciplineGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = discipline.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Discipline.Query().
//		Select(discipline.FieldName).
//		Scan(ctx, &v)
func (dq *DisciplineQuery) Select(fields ...string) *DisciplineSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DisciplineSelect{DisciplineQuery: dq}
	sbuild.label = discipline.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DisciplineSelect configured with the given aggregations.
func (dq *DisciplineQuery) Aggregate(fns ...AggregateFunc) *DisciplineSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DisciplineQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !discipline.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DisciplineQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Discipline, error) {
	var (
		nodes = []*Discipline{}
		_spec = dq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Discipline).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Discipline{config: dq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dq *DisciplineQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DisciplineQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(discipline.Table, discipline.Columns, sqlgraph.NewFieldSpec(discipline.FieldID, field.TypeInt))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, discipline.FieldID)
		for i := range fields {
			if fields[i] != discipline.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DisciplineQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(discipline.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = discipline.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DisciplineGroupBy is the group-by builder for Discipline entities.
type DisciplineGroupBy struct {
	selector
	build *DisciplineQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DisciplineGroupBy) Aggregate(fns ...AggregateFunc) *DisciplineGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DisciplineGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DisciplineQuery, *DisciplineGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DisciplineGroupBy) sqlScan(ctx context.Context, root *DisciplineQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DisciplineSelect is the builder for selecting fields of Discipline entities.
type DisciplineSelect struct {
	*DisciplineQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DisciplineSelect) Aggregate(fns ...AggregateFunc) *DisciplineSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DisciplineSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DisciplineQuery, *DisciplineSelect](ctx, ds.DisciplineQuery, ds, ds.inters, v)
}

func (ds *DisciplineSelect) sqlScan(ctx context.Context, root *DisciplineQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}