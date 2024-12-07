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
	"github.com/liukeshao/echo-admin/ent/org"
	"github.com/liukeshao/echo-admin/ent/predicate"
)

// OrgUpdate is the builder for updating Org entities.
type OrgUpdate struct {
	config
	hooks    []Hook
	mutation *OrgMutation
}

// Where appends a list predicates to the OrgUpdate builder.
func (ou *OrgUpdate) Where(ps ...predicate.Org) *OrgUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetDeleteTime sets the "delete_time" field.
func (ou *OrgUpdate) SetDeleteTime(i int64) *OrgUpdate {
	ou.mutation.ResetDeleteTime()
	ou.mutation.SetDeleteTime(i)
	return ou
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ou *OrgUpdate) SetNillableDeleteTime(i *int64) *OrgUpdate {
	if i != nil {
		ou.SetDeleteTime(*i)
	}
	return ou
}

// AddDeleteTime adds i to the "delete_time" field.
func (ou *OrgUpdate) AddDeleteTime(i int64) *OrgUpdate {
	ou.mutation.AddDeleteTime(i)
	return ou
}

// SetUpdateTime sets the "update_time" field.
func (ou *OrgUpdate) SetUpdateTime(t time.Time) *OrgUpdate {
	ou.mutation.SetUpdateTime(t)
	return ou
}

// SetName sets the "name" field.
func (ou *OrgUpdate) SetName(s string) *OrgUpdate {
	ou.mutation.SetName(s)
	return ou
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ou *OrgUpdate) SetNillableName(s *string) *OrgUpdate {
	if s != nil {
		ou.SetName(*s)
	}
	return ou
}

// SetDisplayName sets the "display_name" field.
func (ou *OrgUpdate) SetDisplayName(s string) *OrgUpdate {
	ou.mutation.SetDisplayName(s)
	return ou
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ou *OrgUpdate) SetNillableDisplayName(s *string) *OrgUpdate {
	if s != nil {
		ou.SetDisplayName(*s)
	}
	return ou
}

// ClearDisplayName clears the value of the "display_name" field.
func (ou *OrgUpdate) ClearDisplayName() *OrgUpdate {
	ou.mutation.ClearDisplayName()
	return ou
}

// SetDescription sets the "description" field.
func (ou *OrgUpdate) SetDescription(s string) *OrgUpdate {
	ou.mutation.SetDescription(s)
	return ou
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ou *OrgUpdate) SetNillableDescription(s *string) *OrgUpdate {
	if s != nil {
		ou.SetDescription(*s)
	}
	return ou
}

// ClearDescription clears the value of the "description" field.
func (ou *OrgUpdate) ClearDescription() *OrgUpdate {
	ou.mutation.ClearDescription()
	return ou
}

// SetLogo sets the "logo" field.
func (ou *OrgUpdate) SetLogo(s string) *OrgUpdate {
	ou.mutation.SetLogo(s)
	return ou
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (ou *OrgUpdate) SetNillableLogo(s *string) *OrgUpdate {
	if s != nil {
		ou.SetLogo(*s)
	}
	return ou
}

// SetStatus sets the "status" field.
func (ou *OrgUpdate) SetStatus(s string) *OrgUpdate {
	ou.mutation.SetStatus(s)
	return ou
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ou *OrgUpdate) SetNillableStatus(s *string) *OrgUpdate {
	if s != nil {
		ou.SetStatus(*s)
	}
	return ou
}

// SetType sets the "type" field.
func (ou *OrgUpdate) SetType(s string) *OrgUpdate {
	ou.mutation.SetType(s)
	return ou
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ou *OrgUpdate) SetNillableType(s *string) *OrgUpdate {
	if s != nil {
		ou.SetType(*s)
	}
	return ou
}

// Mutation returns the OrgMutation object of the builder.
func (ou *OrgUpdate) Mutation() *OrgMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrgUpdate) Save(ctx context.Context) (int, error) {
	if err := ou.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrgUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrgUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrgUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OrgUpdate) defaults() error {
	if _, ok := ou.mutation.UpdateTime(); !ok {
		if org.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized org.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := org.UpdateDefaultUpdateTime()
		ou.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrgUpdate) check() error {
	if v, ok := ou.mutation.Name(); ok {
		if err := org.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Org.name": %w`, err)}
		}
	}
	if v, ok := ou.mutation.Logo(); ok {
		if err := org.LogoValidator(v); err != nil {
			return &ValidationError{Name: "logo", err: fmt.Errorf(`ent: validator failed for field "Org.logo": %w`, err)}
		}
	}
	if v, ok := ou.mutation.Status(); ok {
		if err := org.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Org.status": %w`, err)}
		}
	}
	if v, ok := ou.mutation.GetType(); ok {
		if err := org.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Org.type": %w`, err)}
		}
	}
	return nil
}

func (ou *OrgUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(org.Table, org.Columns, sqlgraph.NewFieldSpec(org.FieldID, field.TypeInt64))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.DeleteTime(); ok {
		_spec.SetField(org.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := ou.mutation.AddedDeleteTime(); ok {
		_spec.AddField(org.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := ou.mutation.UpdateTime(); ok {
		_spec.SetField(org.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.SetField(org.FieldName, field.TypeString, value)
	}
	if value, ok := ou.mutation.DisplayName(); ok {
		_spec.SetField(org.FieldDisplayName, field.TypeString, value)
	}
	if ou.mutation.DisplayNameCleared() {
		_spec.ClearField(org.FieldDisplayName, field.TypeString)
	}
	if value, ok := ou.mutation.Description(); ok {
		_spec.SetField(org.FieldDescription, field.TypeString, value)
	}
	if ou.mutation.DescriptionCleared() {
		_spec.ClearField(org.FieldDescription, field.TypeString)
	}
	if value, ok := ou.mutation.Logo(); ok {
		_spec.SetField(org.FieldLogo, field.TypeString, value)
	}
	if value, ok := ou.mutation.Status(); ok {
		_spec.SetField(org.FieldStatus, field.TypeString, value)
	}
	if value, ok := ou.mutation.GetType(); ok {
		_spec.SetField(org.FieldType, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{org.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrgUpdateOne is the builder for updating a single Org entity.
type OrgUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrgMutation
}

// SetDeleteTime sets the "delete_time" field.
func (ouo *OrgUpdateOne) SetDeleteTime(i int64) *OrgUpdateOne {
	ouo.mutation.ResetDeleteTime()
	ouo.mutation.SetDeleteTime(i)
	return ouo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ouo *OrgUpdateOne) SetNillableDeleteTime(i *int64) *OrgUpdateOne {
	if i != nil {
		ouo.SetDeleteTime(*i)
	}
	return ouo
}

// AddDeleteTime adds i to the "delete_time" field.
func (ouo *OrgUpdateOne) AddDeleteTime(i int64) *OrgUpdateOne {
	ouo.mutation.AddDeleteTime(i)
	return ouo
}

// SetUpdateTime sets the "update_time" field.
func (ouo *OrgUpdateOne) SetUpdateTime(t time.Time) *OrgUpdateOne {
	ouo.mutation.SetUpdateTime(t)
	return ouo
}

// SetName sets the "name" field.
func (ouo *OrgUpdateOne) SetName(s string) *OrgUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ouo *OrgUpdateOne) SetNillableName(s *string) *OrgUpdateOne {
	if s != nil {
		ouo.SetName(*s)
	}
	return ouo
}

// SetDisplayName sets the "display_name" field.
func (ouo *OrgUpdateOne) SetDisplayName(s string) *OrgUpdateOne {
	ouo.mutation.SetDisplayName(s)
	return ouo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ouo *OrgUpdateOne) SetNillableDisplayName(s *string) *OrgUpdateOne {
	if s != nil {
		ouo.SetDisplayName(*s)
	}
	return ouo
}

// ClearDisplayName clears the value of the "display_name" field.
func (ouo *OrgUpdateOne) ClearDisplayName() *OrgUpdateOne {
	ouo.mutation.ClearDisplayName()
	return ouo
}

// SetDescription sets the "description" field.
func (ouo *OrgUpdateOne) SetDescription(s string) *OrgUpdateOne {
	ouo.mutation.SetDescription(s)
	return ouo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ouo *OrgUpdateOne) SetNillableDescription(s *string) *OrgUpdateOne {
	if s != nil {
		ouo.SetDescription(*s)
	}
	return ouo
}

// ClearDescription clears the value of the "description" field.
func (ouo *OrgUpdateOne) ClearDescription() *OrgUpdateOne {
	ouo.mutation.ClearDescription()
	return ouo
}

// SetLogo sets the "logo" field.
func (ouo *OrgUpdateOne) SetLogo(s string) *OrgUpdateOne {
	ouo.mutation.SetLogo(s)
	return ouo
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (ouo *OrgUpdateOne) SetNillableLogo(s *string) *OrgUpdateOne {
	if s != nil {
		ouo.SetLogo(*s)
	}
	return ouo
}

// SetStatus sets the "status" field.
func (ouo *OrgUpdateOne) SetStatus(s string) *OrgUpdateOne {
	ouo.mutation.SetStatus(s)
	return ouo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ouo *OrgUpdateOne) SetNillableStatus(s *string) *OrgUpdateOne {
	if s != nil {
		ouo.SetStatus(*s)
	}
	return ouo
}

// SetType sets the "type" field.
func (ouo *OrgUpdateOne) SetType(s string) *OrgUpdateOne {
	ouo.mutation.SetType(s)
	return ouo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ouo *OrgUpdateOne) SetNillableType(s *string) *OrgUpdateOne {
	if s != nil {
		ouo.SetType(*s)
	}
	return ouo
}

// Mutation returns the OrgMutation object of the builder.
func (ouo *OrgUpdateOne) Mutation() *OrgMutation {
	return ouo.mutation
}

// Where appends a list predicates to the OrgUpdate builder.
func (ouo *OrgUpdateOne) Where(ps ...predicate.Org) *OrgUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrgUpdateOne) Select(field string, fields ...string) *OrgUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Org entity.
func (ouo *OrgUpdateOne) Save(ctx context.Context) (*Org, error) {
	if err := ouo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrgUpdateOne) SaveX(ctx context.Context) *Org {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrgUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrgUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OrgUpdateOne) defaults() error {
	if _, ok := ouo.mutation.UpdateTime(); !ok {
		if org.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized org.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := org.UpdateDefaultUpdateTime()
		ouo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrgUpdateOne) check() error {
	if v, ok := ouo.mutation.Name(); ok {
		if err := org.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Org.name": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.Logo(); ok {
		if err := org.LogoValidator(v); err != nil {
			return &ValidationError{Name: "logo", err: fmt.Errorf(`ent: validator failed for field "Org.logo": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.Status(); ok {
		if err := org.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Org.status": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.GetType(); ok {
		if err := org.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Org.type": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrgUpdateOne) sqlSave(ctx context.Context) (_node *Org, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(org.Table, org.Columns, sqlgraph.NewFieldSpec(org.FieldID, field.TypeInt64))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Org.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, org.FieldID)
		for _, f := range fields {
			if !org.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != org.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.DeleteTime(); ok {
		_spec.SetField(org.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := ouo.mutation.AddedDeleteTime(); ok {
		_spec.AddField(org.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := ouo.mutation.UpdateTime(); ok {
		_spec.SetField(org.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.SetField(org.FieldName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.DisplayName(); ok {
		_spec.SetField(org.FieldDisplayName, field.TypeString, value)
	}
	if ouo.mutation.DisplayNameCleared() {
		_spec.ClearField(org.FieldDisplayName, field.TypeString)
	}
	if value, ok := ouo.mutation.Description(); ok {
		_spec.SetField(org.FieldDescription, field.TypeString, value)
	}
	if ouo.mutation.DescriptionCleared() {
		_spec.ClearField(org.FieldDescription, field.TypeString)
	}
	if value, ok := ouo.mutation.Logo(); ok {
		_spec.SetField(org.FieldLogo, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Status(); ok {
		_spec.SetField(org.FieldStatus, field.TypeString, value)
	}
	if value, ok := ouo.mutation.GetType(); ok {
		_spec.SetField(org.FieldType, field.TypeString, value)
	}
	_node = &Org{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{org.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}