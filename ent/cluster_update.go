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
	"github.com/liukeshao/echo-admin/ent/cluster"
	"github.com/liukeshao/echo-admin/ent/predicate"
)

// ClusterUpdate is the builder for updating Cluster entities.
type ClusterUpdate struct {
	config
	hooks    []Hook
	mutation *ClusterMutation
}

// Where appends a list predicates to the ClusterUpdate builder.
func (cu *ClusterUpdate) Where(ps ...predicate.Cluster) *ClusterUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetDeleteTime sets the "delete_time" field.
func (cu *ClusterUpdate) SetDeleteTime(i int64) *ClusterUpdate {
	cu.mutation.ResetDeleteTime()
	cu.mutation.SetDeleteTime(i)
	return cu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (cu *ClusterUpdate) SetNillableDeleteTime(i *int64) *ClusterUpdate {
	if i != nil {
		cu.SetDeleteTime(*i)
	}
	return cu
}

// AddDeleteTime adds i to the "delete_time" field.
func (cu *ClusterUpdate) AddDeleteTime(i int64) *ClusterUpdate {
	cu.mutation.AddDeleteTime(i)
	return cu
}

// SetUpdateTime sets the "update_time" field.
func (cu *ClusterUpdate) SetUpdateTime(t time.Time) *ClusterUpdate {
	cu.mutation.SetUpdateTime(t)
	return cu
}

// SetOrgID sets the "org_id" field.
func (cu *ClusterUpdate) SetOrgID(u uint64) *ClusterUpdate {
	cu.mutation.ResetOrgID()
	cu.mutation.SetOrgID(u)
	return cu
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (cu *ClusterUpdate) SetNillableOrgID(u *uint64) *ClusterUpdate {
	if u != nil {
		cu.SetOrgID(*u)
	}
	return cu
}

// AddOrgID adds u to the "org_id" field.
func (cu *ClusterUpdate) AddOrgID(u int64) *ClusterUpdate {
	cu.mutation.AddOrgID(u)
	return cu
}

// SetName sets the "name" field.
func (cu *ClusterUpdate) SetName(s string) *ClusterUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *ClusterUpdate) SetNillableName(s *string) *ClusterUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetDisplayName sets the "display_name" field.
func (cu *ClusterUpdate) SetDisplayName(s string) *ClusterUpdate {
	cu.mutation.SetDisplayName(s)
	return cu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (cu *ClusterUpdate) SetNillableDisplayName(s *string) *ClusterUpdate {
	if s != nil {
		cu.SetDisplayName(*s)
	}
	return cu
}

// ClearDisplayName clears the value of the "display_name" field.
func (cu *ClusterUpdate) ClearDisplayName() *ClusterUpdate {
	cu.mutation.ClearDisplayName()
	return cu
}

// SetDescription sets the "description" field.
func (cu *ClusterUpdate) SetDescription(s string) *ClusterUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *ClusterUpdate) SetNillableDescription(s *string) *ClusterUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// ClearDescription clears the value of the "description" field.
func (cu *ClusterUpdate) ClearDescription() *ClusterUpdate {
	cu.mutation.ClearDescription()
	return cu
}

// SetConfig sets the "config" field.
func (cu *ClusterUpdate) SetConfig(s string) *ClusterUpdate {
	cu.mutation.SetConfig(s)
	return cu
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (cu *ClusterUpdate) SetNillableConfig(s *string) *ClusterUpdate {
	if s != nil {
		cu.SetConfig(*s)
	}
	return cu
}

// ClearConfig clears the value of the "config" field.
func (cu *ClusterUpdate) ClearConfig() *ClusterUpdate {
	cu.mutation.ClearConfig()
	return cu
}

// Mutation returns the ClusterMutation object of the builder.
func (cu *ClusterUpdate) Mutation() *ClusterMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ClusterUpdate) Save(ctx context.Context) (int, error) {
	if err := cu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ClusterUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ClusterUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ClusterUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ClusterUpdate) defaults() error {
	if _, ok := cu.mutation.UpdateTime(); !ok {
		if cluster.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized cluster.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := cluster.UpdateDefaultUpdateTime()
		cu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cu *ClusterUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := cluster.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Cluster.name": %w`, err)}
		}
	}
	return nil
}

func (cu *ClusterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(cluster.Table, cluster.Columns, sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeUint64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.DeleteTime(); ok {
		_spec.SetField(cluster.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedDeleteTime(); ok {
		_spec.AddField(cluster.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.UpdateTime(); ok {
		_spec.SetField(cluster.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cu.mutation.OrgID(); ok {
		_spec.SetField(cluster.FieldOrgID, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.AddedOrgID(); ok {
		_spec.AddField(cluster.FieldOrgID, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(cluster.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.DisplayName(); ok {
		_spec.SetField(cluster.FieldDisplayName, field.TypeString, value)
	}
	if cu.mutation.DisplayNameCleared() {
		_spec.ClearField(cluster.FieldDisplayName, field.TypeString)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(cluster.FieldDescription, field.TypeString, value)
	}
	if cu.mutation.DescriptionCleared() {
		_spec.ClearField(cluster.FieldDescription, field.TypeString)
	}
	if value, ok := cu.mutation.Config(); ok {
		_spec.SetField(cluster.FieldConfig, field.TypeString, value)
	}
	if cu.mutation.ConfigCleared() {
		_spec.ClearField(cluster.FieldConfig, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cluster.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ClusterUpdateOne is the builder for updating a single Cluster entity.
type ClusterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClusterMutation
}

// SetDeleteTime sets the "delete_time" field.
func (cuo *ClusterUpdateOne) SetDeleteTime(i int64) *ClusterUpdateOne {
	cuo.mutation.ResetDeleteTime()
	cuo.mutation.SetDeleteTime(i)
	return cuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (cuo *ClusterUpdateOne) SetNillableDeleteTime(i *int64) *ClusterUpdateOne {
	if i != nil {
		cuo.SetDeleteTime(*i)
	}
	return cuo
}

// AddDeleteTime adds i to the "delete_time" field.
func (cuo *ClusterUpdateOne) AddDeleteTime(i int64) *ClusterUpdateOne {
	cuo.mutation.AddDeleteTime(i)
	return cuo
}

// SetUpdateTime sets the "update_time" field.
func (cuo *ClusterUpdateOne) SetUpdateTime(t time.Time) *ClusterUpdateOne {
	cuo.mutation.SetUpdateTime(t)
	return cuo
}

// SetOrgID sets the "org_id" field.
func (cuo *ClusterUpdateOne) SetOrgID(u uint64) *ClusterUpdateOne {
	cuo.mutation.ResetOrgID()
	cuo.mutation.SetOrgID(u)
	return cuo
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (cuo *ClusterUpdateOne) SetNillableOrgID(u *uint64) *ClusterUpdateOne {
	if u != nil {
		cuo.SetOrgID(*u)
	}
	return cuo
}

// AddOrgID adds u to the "org_id" field.
func (cuo *ClusterUpdateOne) AddOrgID(u int64) *ClusterUpdateOne {
	cuo.mutation.AddOrgID(u)
	return cuo
}

// SetName sets the "name" field.
func (cuo *ClusterUpdateOne) SetName(s string) *ClusterUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *ClusterUpdateOne) SetNillableName(s *string) *ClusterUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetDisplayName sets the "display_name" field.
func (cuo *ClusterUpdateOne) SetDisplayName(s string) *ClusterUpdateOne {
	cuo.mutation.SetDisplayName(s)
	return cuo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (cuo *ClusterUpdateOne) SetNillableDisplayName(s *string) *ClusterUpdateOne {
	if s != nil {
		cuo.SetDisplayName(*s)
	}
	return cuo
}

// ClearDisplayName clears the value of the "display_name" field.
func (cuo *ClusterUpdateOne) ClearDisplayName() *ClusterUpdateOne {
	cuo.mutation.ClearDisplayName()
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *ClusterUpdateOne) SetDescription(s string) *ClusterUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *ClusterUpdateOne) SetNillableDescription(s *string) *ClusterUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// ClearDescription clears the value of the "description" field.
func (cuo *ClusterUpdateOne) ClearDescription() *ClusterUpdateOne {
	cuo.mutation.ClearDescription()
	return cuo
}

// SetConfig sets the "config" field.
func (cuo *ClusterUpdateOne) SetConfig(s string) *ClusterUpdateOne {
	cuo.mutation.SetConfig(s)
	return cuo
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (cuo *ClusterUpdateOne) SetNillableConfig(s *string) *ClusterUpdateOne {
	if s != nil {
		cuo.SetConfig(*s)
	}
	return cuo
}

// ClearConfig clears the value of the "config" field.
func (cuo *ClusterUpdateOne) ClearConfig() *ClusterUpdateOne {
	cuo.mutation.ClearConfig()
	return cuo
}

// Mutation returns the ClusterMutation object of the builder.
func (cuo *ClusterUpdateOne) Mutation() *ClusterMutation {
	return cuo.mutation
}

// Where appends a list predicates to the ClusterUpdate builder.
func (cuo *ClusterUpdateOne) Where(ps ...predicate.Cluster) *ClusterUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ClusterUpdateOne) Select(field string, fields ...string) *ClusterUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cluster entity.
func (cuo *ClusterUpdateOne) Save(ctx context.Context) (*Cluster, error) {
	if err := cuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ClusterUpdateOne) SaveX(ctx context.Context) *Cluster {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ClusterUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ClusterUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ClusterUpdateOne) defaults() error {
	if _, ok := cuo.mutation.UpdateTime(); !ok {
		if cluster.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized cluster.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := cluster.UpdateDefaultUpdateTime()
		cuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ClusterUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := cluster.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Cluster.name": %w`, err)}
		}
	}
	return nil
}

func (cuo *ClusterUpdateOne) sqlSave(ctx context.Context) (_node *Cluster, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(cluster.Table, cluster.Columns, sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeUint64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cluster.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cluster.FieldID)
		for _, f := range fields {
			if !cluster.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cluster.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.DeleteTime(); ok {
		_spec.SetField(cluster.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedDeleteTime(); ok {
		_spec.AddField(cluster.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.UpdateTime(); ok {
		_spec.SetField(cluster.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.OrgID(); ok {
		_spec.SetField(cluster.FieldOrgID, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.AddedOrgID(); ok {
		_spec.AddField(cluster.FieldOrgID, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(cluster.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.DisplayName(); ok {
		_spec.SetField(cluster.FieldDisplayName, field.TypeString, value)
	}
	if cuo.mutation.DisplayNameCleared() {
		_spec.ClearField(cluster.FieldDisplayName, field.TypeString)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(cluster.FieldDescription, field.TypeString, value)
	}
	if cuo.mutation.DescriptionCleared() {
		_spec.ClearField(cluster.FieldDescription, field.TypeString)
	}
	if value, ok := cuo.mutation.Config(); ok {
		_spec.SetField(cluster.FieldConfig, field.TypeString, value)
	}
	if cuo.mutation.ConfigCleared() {
		_spec.ClearField(cluster.FieldConfig, field.TypeString)
	}
	_node = &Cluster{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cluster.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
