// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/liukeshao/echo-admin/ent/predicate"
	"github.com/liukeshao/echo-admin/ent/rolebinding"
)

// RoleBindingQuery is the builder for querying RoleBinding entities.
type RoleBindingQuery struct {
	config
	ctx        *QueryContext
	order      []rolebinding.OrderOption
	inters     []Interceptor
	predicates []predicate.RoleBinding
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoleBindingQuery builder.
func (rbq *RoleBindingQuery) Where(ps ...predicate.RoleBinding) *RoleBindingQuery {
	rbq.predicates = append(rbq.predicates, ps...)
	return rbq
}

// Limit the number of records to be returned by this query.
func (rbq *RoleBindingQuery) Limit(limit int) *RoleBindingQuery {
	rbq.ctx.Limit = &limit
	return rbq
}

// Offset to start from.
func (rbq *RoleBindingQuery) Offset(offset int) *RoleBindingQuery {
	rbq.ctx.Offset = &offset
	return rbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rbq *RoleBindingQuery) Unique(unique bool) *RoleBindingQuery {
	rbq.ctx.Unique = &unique
	return rbq
}

// Order specifies how the records should be ordered.
func (rbq *RoleBindingQuery) Order(o ...rolebinding.OrderOption) *RoleBindingQuery {
	rbq.order = append(rbq.order, o...)
	return rbq
}

// First returns the first RoleBinding entity from the query.
// Returns a *NotFoundError when no RoleBinding was found.
func (rbq *RoleBindingQuery) First(ctx context.Context) (*RoleBinding, error) {
	nodes, err := rbq.Limit(1).All(setContextOp(ctx, rbq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rolebinding.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rbq *RoleBindingQuery) FirstX(ctx context.Context) *RoleBinding {
	node, err := rbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RoleBinding ID from the query.
// Returns a *NotFoundError when no RoleBinding ID was found.
func (rbq *RoleBindingQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rbq.Limit(1).IDs(setContextOp(ctx, rbq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rolebinding.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rbq *RoleBindingQuery) FirstIDX(ctx context.Context) string {
	id, err := rbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RoleBinding entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RoleBinding entity is found.
// Returns a *NotFoundError when no RoleBinding entities are found.
func (rbq *RoleBindingQuery) Only(ctx context.Context) (*RoleBinding, error) {
	nodes, err := rbq.Limit(2).All(setContextOp(ctx, rbq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rolebinding.Label}
	default:
		return nil, &NotSingularError{rolebinding.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rbq *RoleBindingQuery) OnlyX(ctx context.Context) *RoleBinding {
	node, err := rbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RoleBinding ID in the query.
// Returns a *NotSingularError when more than one RoleBinding ID is found.
// Returns a *NotFoundError when no entities are found.
func (rbq *RoleBindingQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rbq.Limit(2).IDs(setContextOp(ctx, rbq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rolebinding.Label}
	default:
		err = &NotSingularError{rolebinding.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rbq *RoleBindingQuery) OnlyIDX(ctx context.Context) string {
	id, err := rbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RoleBindings.
func (rbq *RoleBindingQuery) All(ctx context.Context) ([]*RoleBinding, error) {
	ctx = setContextOp(ctx, rbq.ctx, ent.OpQueryAll)
	if err := rbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RoleBinding, *RoleBindingQuery]()
	return withInterceptors[[]*RoleBinding](ctx, rbq, qr, rbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rbq *RoleBindingQuery) AllX(ctx context.Context) []*RoleBinding {
	nodes, err := rbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RoleBinding IDs.
func (rbq *RoleBindingQuery) IDs(ctx context.Context) (ids []string, err error) {
	if rbq.ctx.Unique == nil && rbq.path != nil {
		rbq.Unique(true)
	}
	ctx = setContextOp(ctx, rbq.ctx, ent.OpQueryIDs)
	if err = rbq.Select(rolebinding.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rbq *RoleBindingQuery) IDsX(ctx context.Context) []string {
	ids, err := rbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rbq *RoleBindingQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rbq.ctx, ent.OpQueryCount)
	if err := rbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rbq, querierCount[*RoleBindingQuery](), rbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rbq *RoleBindingQuery) CountX(ctx context.Context) int {
	count, err := rbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rbq *RoleBindingQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rbq.ctx, ent.OpQueryExist)
	switch _, err := rbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rbq *RoleBindingQuery) ExistX(ctx context.Context) bool {
	exist, err := rbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoleBindingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rbq *RoleBindingQuery) Clone() *RoleBindingQuery {
	if rbq == nil {
		return nil
	}
	return &RoleBindingQuery{
		config:     rbq.config,
		ctx:        rbq.ctx.Clone(),
		order:      append([]rolebinding.OrderOption{}, rbq.order...),
		inters:     append([]Interceptor{}, rbq.inters...),
		predicates: append([]predicate.RoleBinding{}, rbq.predicates...),
		// clone intermediate query.
		sql:  rbq.sql.Clone(),
		path: rbq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DeleteTime int64 `json:"delete_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RoleBinding.Query().
//		GroupBy(rolebinding.FieldDeleteTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rbq *RoleBindingQuery) GroupBy(field string, fields ...string) *RoleBindingGroupBy {
	rbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RoleBindingGroupBy{build: rbq}
	grbuild.flds = &rbq.ctx.Fields
	grbuild.label = rolebinding.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DeleteTime int64 `json:"delete_time,omitempty"`
//	}
//
//	client.RoleBinding.Query().
//		Select(rolebinding.FieldDeleteTime).
//		Scan(ctx, &v)
func (rbq *RoleBindingQuery) Select(fields ...string) *RoleBindingSelect {
	rbq.ctx.Fields = append(rbq.ctx.Fields, fields...)
	sbuild := &RoleBindingSelect{RoleBindingQuery: rbq}
	sbuild.label = rolebinding.Label
	sbuild.flds, sbuild.scan = &rbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RoleBindingSelect configured with the given aggregations.
func (rbq *RoleBindingQuery) Aggregate(fns ...AggregateFunc) *RoleBindingSelect {
	return rbq.Select().Aggregate(fns...)
}

func (rbq *RoleBindingQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rbq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rbq); err != nil {
				return err
			}
		}
	}
	for _, f := range rbq.ctx.Fields {
		if !rolebinding.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rbq.path != nil {
		prev, err := rbq.path(ctx)
		if err != nil {
			return err
		}
		rbq.sql = prev
	}
	return nil
}

func (rbq *RoleBindingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RoleBinding, error) {
	var (
		nodes = []*RoleBinding{}
		_spec = rbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RoleBinding).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RoleBinding{config: rbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (rbq *RoleBindingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rbq.querySpec()
	_spec.Node.Columns = rbq.ctx.Fields
	if len(rbq.ctx.Fields) > 0 {
		_spec.Unique = rbq.ctx.Unique != nil && *rbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rbq.driver, _spec)
}

func (rbq *RoleBindingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(rolebinding.Table, rolebinding.Columns, sqlgraph.NewFieldSpec(rolebinding.FieldID, field.TypeString))
	_spec.From = rbq.sql
	if unique := rbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rbq.path != nil {
		_spec.Unique = true
	}
	if fields := rbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rolebinding.FieldID)
		for i := range fields {
			if fields[i] != rolebinding.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rbq *RoleBindingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rbq.driver.Dialect())
	t1 := builder.Table(rolebinding.Table)
	columns := rbq.ctx.Fields
	if len(columns) == 0 {
		columns = rolebinding.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rbq.sql != nil {
		selector = rbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rbq.ctx.Unique != nil && *rbq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rbq.predicates {
		p(selector)
	}
	for _, p := range rbq.order {
		p(selector)
	}
	if offset := rbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RoleBindingGroupBy is the group-by builder for RoleBinding entities.
type RoleBindingGroupBy struct {
	selector
	build *RoleBindingQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rbgb *RoleBindingGroupBy) Aggregate(fns ...AggregateFunc) *RoleBindingGroupBy {
	rbgb.fns = append(rbgb.fns, fns...)
	return rbgb
}

// Scan applies the selector query and scans the result into the given value.
func (rbgb *RoleBindingGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rbgb.build.ctx, ent.OpQueryGroupBy)
	if err := rbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoleBindingQuery, *RoleBindingGroupBy](ctx, rbgb.build, rbgb, rbgb.build.inters, v)
}

func (rbgb *RoleBindingGroupBy) sqlScan(ctx context.Context, root *RoleBindingQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rbgb.fns))
	for _, fn := range rbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rbgb.flds)+len(rbgb.fns))
		for _, f := range *rbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RoleBindingSelect is the builder for selecting fields of RoleBinding entities.
type RoleBindingSelect struct {
	*RoleBindingQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rbs *RoleBindingSelect) Aggregate(fns ...AggregateFunc) *RoleBindingSelect {
	rbs.fns = append(rbs.fns, fns...)
	return rbs
}

// Scan applies the selector query and scans the result into the given value.
func (rbs *RoleBindingSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rbs.ctx, ent.OpQuerySelect)
	if err := rbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoleBindingQuery, *RoleBindingSelect](ctx, rbs.RoleBindingQuery, rbs, rbs.inters, v)
}

func (rbs *RoleBindingSelect) sqlScan(ctx context.Context, root *RoleBindingQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rbs.fns))
	for _, fn := range rbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
