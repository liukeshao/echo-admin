// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/liukeshao/echo-admin/ent/rolebinding"
)

// RoleBindingCreate is the builder for creating a RoleBinding entity.
type RoleBindingCreate struct {
	config
	mutation *RoleBindingMutation
	hooks    []Hook
}

// SetDeleteTime sets the "delete_time" field.
func (rbc *RoleBindingCreate) SetDeleteTime(i int64) *RoleBindingCreate {
	rbc.mutation.SetDeleteTime(i)
	return rbc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (rbc *RoleBindingCreate) SetNillableDeleteTime(i *int64) *RoleBindingCreate {
	if i != nil {
		rbc.SetDeleteTime(*i)
	}
	return rbc
}

// SetCreateTime sets the "create_time" field.
func (rbc *RoleBindingCreate) SetCreateTime(t time.Time) *RoleBindingCreate {
	rbc.mutation.SetCreateTime(t)
	return rbc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rbc *RoleBindingCreate) SetNillableCreateTime(t *time.Time) *RoleBindingCreate {
	if t != nil {
		rbc.SetCreateTime(*t)
	}
	return rbc
}

// SetUpdateTime sets the "update_time" field.
func (rbc *RoleBindingCreate) SetUpdateTime(t time.Time) *RoleBindingCreate {
	rbc.mutation.SetUpdateTime(t)
	return rbc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (rbc *RoleBindingCreate) SetNillableUpdateTime(t *time.Time) *RoleBindingCreate {
	if t != nil {
		rbc.SetUpdateTime(*t)
	}
	return rbc
}

// SetName sets the "name" field.
func (rbc *RoleBindingCreate) SetName(s string) *RoleBindingCreate {
	rbc.mutation.SetName(s)
	return rbc
}

// SetDisplayName sets the "display_name" field.
func (rbc *RoleBindingCreate) SetDisplayName(s string) *RoleBindingCreate {
	rbc.mutation.SetDisplayName(s)
	return rbc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (rbc *RoleBindingCreate) SetNillableDisplayName(s *string) *RoleBindingCreate {
	if s != nil {
		rbc.SetDisplayName(*s)
	}
	return rbc
}

// SetRoleName sets the "role_name" field.
func (rbc *RoleBindingCreate) SetRoleName(s string) *RoleBindingCreate {
	rbc.mutation.SetRoleName(s)
	return rbc
}

// SetUserID sets the "user_id" field.
func (rbc *RoleBindingCreate) SetUserID(s string) *RoleBindingCreate {
	rbc.mutation.SetUserID(s)
	return rbc
}

// SetDeletedTime sets the "deleted_time" field.
func (rbc *RoleBindingCreate) SetDeletedTime(i int64) *RoleBindingCreate {
	rbc.mutation.SetDeletedTime(i)
	return rbc
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (rbc *RoleBindingCreate) SetNillableDeletedTime(i *int64) *RoleBindingCreate {
	if i != nil {
		rbc.SetDeletedTime(*i)
	}
	return rbc
}

// SetID sets the "id" field.
func (rbc *RoleBindingCreate) SetID(s string) *RoleBindingCreate {
	rbc.mutation.SetID(s)
	return rbc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rbc *RoleBindingCreate) SetNillableID(s *string) *RoleBindingCreate {
	if s != nil {
		rbc.SetID(*s)
	}
	return rbc
}

// Mutation returns the RoleBindingMutation object of the builder.
func (rbc *RoleBindingCreate) Mutation() *RoleBindingMutation {
	return rbc.mutation
}

// Save creates the RoleBinding in the database.
func (rbc *RoleBindingCreate) Save(ctx context.Context) (*RoleBinding, error) {
	if err := rbc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, rbc.sqlSave, rbc.mutation, rbc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rbc *RoleBindingCreate) SaveX(ctx context.Context) *RoleBinding {
	v, err := rbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rbc *RoleBindingCreate) Exec(ctx context.Context) error {
	_, err := rbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rbc *RoleBindingCreate) ExecX(ctx context.Context) {
	if err := rbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rbc *RoleBindingCreate) defaults() error {
	if _, ok := rbc.mutation.DeleteTime(); !ok {
		v := rolebinding.DefaultDeleteTime
		rbc.mutation.SetDeleteTime(v)
	}
	if _, ok := rbc.mutation.CreateTime(); !ok {
		if rolebinding.DefaultCreateTime == nil {
			return fmt.Errorf("ent: uninitialized rolebinding.DefaultCreateTime (forgotten import ent/runtime?)")
		}
		v := rolebinding.DefaultCreateTime()
		rbc.mutation.SetCreateTime(v)
	}
	if _, ok := rbc.mutation.UpdateTime(); !ok {
		if rolebinding.DefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized rolebinding.DefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := rolebinding.DefaultUpdateTime()
		rbc.mutation.SetUpdateTime(v)
	}
	if _, ok := rbc.mutation.ID(); !ok {
		if rolebinding.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized rolebinding.DefaultID (forgotten import ent/runtime?)")
		}
		v := rolebinding.DefaultID()
		rbc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rbc *RoleBindingCreate) check() error {
	if _, ok := rbc.mutation.DeleteTime(); !ok {
		return &ValidationError{Name: "delete_time", err: errors.New(`ent: missing required field "RoleBinding.delete_time"`)}
	}
	if _, ok := rbc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "RoleBinding.create_time"`)}
	}
	if _, ok := rbc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "RoleBinding.update_time"`)}
	}
	if _, ok := rbc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "RoleBinding.name"`)}
	}
	if _, ok := rbc.mutation.RoleName(); !ok {
		return &ValidationError{Name: "role_name", err: errors.New(`ent: missing required field "RoleBinding.role_name"`)}
	}
	if _, ok := rbc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "RoleBinding.user_id"`)}
	}
	return nil
}

func (rbc *RoleBindingCreate) sqlSave(ctx context.Context) (*RoleBinding, error) {
	if err := rbc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected RoleBinding.ID type: %T", _spec.ID.Value)
		}
	}
	rbc.mutation.id = &_node.ID
	rbc.mutation.done = true
	return _node, nil
}

func (rbc *RoleBindingCreate) createSpec() (*RoleBinding, *sqlgraph.CreateSpec) {
	var (
		_node = &RoleBinding{config: rbc.config}
		_spec = sqlgraph.NewCreateSpec(rolebinding.Table, sqlgraph.NewFieldSpec(rolebinding.FieldID, field.TypeString))
	)
	if id, ok := rbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rbc.mutation.DeleteTime(); ok {
		_spec.SetField(rolebinding.FieldDeleteTime, field.TypeInt64, value)
		_node.DeleteTime = value
	}
	if value, ok := rbc.mutation.CreateTime(); ok {
		_spec.SetField(rolebinding.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := rbc.mutation.UpdateTime(); ok {
		_spec.SetField(rolebinding.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := rbc.mutation.Name(); ok {
		_spec.SetField(rolebinding.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rbc.mutation.DisplayName(); ok {
		_spec.SetField(rolebinding.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := rbc.mutation.RoleName(); ok {
		_spec.SetField(rolebinding.FieldRoleName, field.TypeString, value)
		_node.RoleName = value
	}
	if value, ok := rbc.mutation.UserID(); ok {
		_spec.SetField(rolebinding.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := rbc.mutation.DeletedTime(); ok {
		_spec.SetField(rolebinding.FieldDeletedTime, field.TypeInt64, value)
		_node.DeletedTime = value
	}
	return _node, _spec
}

// RoleBindingCreateBulk is the builder for creating many RoleBinding entities in bulk.
type RoleBindingCreateBulk struct {
	config
	err      error
	builders []*RoleBindingCreate
}

// Save creates the RoleBinding entities in the database.
func (rbcb *RoleBindingCreateBulk) Save(ctx context.Context) ([]*RoleBinding, error) {
	if rbcb.err != nil {
		return nil, rbcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rbcb.builders))
	nodes := make([]*RoleBinding, len(rbcb.builders))
	mutators := make([]Mutator, len(rbcb.builders))
	for i := range rbcb.builders {
		func(i int, root context.Context) {
			builder := rbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleBindingMutation)
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
					_, err = mutators[i+1].Mutate(root, rbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rbcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, rbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rbcb *RoleBindingCreateBulk) SaveX(ctx context.Context) []*RoleBinding {
	v, err := rbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rbcb *RoleBindingCreateBulk) Exec(ctx context.Context) error {
	_, err := rbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rbcb *RoleBindingCreateBulk) ExecX(ctx context.Context) {
	if err := rbcb.Exec(ctx); err != nil {
		panic(err)
	}
}
