// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/liukeshao/echo-admin/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetDeleteTime sets the "delete_time" field.
func (uc *UserCreate) SetDeleteTime(i int64) *UserCreate {
	uc.mutation.SetDeleteTime(i)
	return uc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeleteTime(i *int64) *UserCreate {
	if i != nil {
		uc.SetDeleteTime(*i)
	}
	return uc
}

// SetCreateTime sets the "create_time" field.
func (uc *UserCreate) SetCreateTime(t time.Time) *UserCreate {
	uc.mutation.SetCreateTime(t)
	return uc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreateTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreateTime(*t)
	}
	return uc
}

// SetUpdateTime sets the "update_time" field.
func (uc *UserCreate) SetUpdateTime(t time.Time) *UserCreate {
	uc.mutation.SetUpdateTime(t)
	return uc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdateTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdateTime(*t)
	}
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uc *UserCreate) SetNillableName(s *string) *UserCreate {
	if s != nil {
		uc.SetName(*s)
	}
	return uc
}

// SetDisplayName sets the "display_name" field.
func (uc *UserCreate) SetDisplayName(s string) *UserCreate {
	uc.mutation.SetDisplayName(s)
	return uc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (uc *UserCreate) SetNillableDisplayName(s *string) *UserCreate {
	if s != nil {
		uc.SetDisplayName(*s)
	}
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uc *UserCreate) SetNillableEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetEmail(*s)
	}
	return uc
}

// SetPhone sets the "phone" field.
func (uc *UserCreate) SetPhone(s string) *UserCreate {
	uc.mutation.SetPhone(s)
	return uc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uc *UserCreate) SetNillablePhone(s *string) *UserCreate {
	if s != nil {
		uc.SetPhone(*s)
	}
	return uc
}

// SetGender sets the "gender" field.
func (uc *UserCreate) SetGender(s string) *UserCreate {
	uc.mutation.SetGender(s)
	return uc
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (uc *UserCreate) SetNillableGender(s *string) *UserCreate {
	if s != nil {
		uc.SetGender(*s)
	}
	return uc
}

// SetAvatar sets the "avatar" field.
func (uc *UserCreate) SetAvatar(s string) *UserCreate {
	uc.mutation.SetAvatar(s)
	return uc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uc *UserCreate) SetNillableAvatar(s *string) *UserCreate {
	if s != nil {
		uc.SetAvatar(*s)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uc *UserCreate) SetNillablePassword(s *string) *UserCreate {
	if s != nil {
		uc.SetPassword(*s)
	}
	return uc
}

// SetNonLocked sets the "non_locked" field.
func (uc *UserCreate) SetNonLocked(i int) *UserCreate {
	uc.mutation.SetNonLocked(i)
	return uc
}

// SetNillableNonLocked sets the "non_locked" field if the given value is not nil.
func (uc *UserCreate) SetNillableNonLocked(i *int) *UserCreate {
	if i != nil {
		uc.SetNonLocked(*i)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(s string) *UserCreate {
	uc.mutation.SetID(s)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(s *string) *UserCreate {
	if s != nil {
		uc.SetID(*s)
	}
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if err := uc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() error {
	if _, ok := uc.mutation.DeleteTime(); !ok {
		v := user.DefaultDeleteTime
		uc.mutation.SetDeleteTime(v)
	}
	if _, ok := uc.mutation.CreateTime(); !ok {
		if user.DefaultCreateTime == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultCreateTime (forgotten import ent/runtime?)")
		}
		v := user.DefaultCreateTime()
		uc.mutation.SetCreateTime(v)
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		if user.DefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := user.DefaultUpdateTime()
		uc.mutation.SetUpdateTime(v)
	}
	if _, ok := uc.mutation.Gender(); !ok {
		v := user.DefaultGender
		uc.mutation.SetGender(v)
	}
	if _, ok := uc.mutation.NonLocked(); !ok {
		v := user.DefaultNonLocked
		uc.mutation.SetNonLocked(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		if user.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultID (forgotten import ent/runtime?)")
		}
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.DeleteTime(); !ok {
		return &ValidationError{Name: "delete_time", err: errors.New(`ent: missing required field "User.delete_time"`)}
	}
	if _, ok := uc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "User.create_time"`)}
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "User.update_time"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected User.ID type: %T", _spec.ID.Value)
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.DeleteTime(); ok {
		_spec.SetField(user.FieldDeleteTime, field.TypeInt64, value)
		_node.DeleteTime = value
	}
	if value, ok := uc.mutation.CreateTime(); ok {
		_spec.SetField(user.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := uc.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.DisplayName(); ok {
		_spec.SetField(user.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := uc.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeString, value)
		_node.Gender = value
	}
	if value, ok := uc.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.NonLocked(); ok {
		_spec.SetField(user.FieldNonLocked, field.TypeInt, value)
		_node.NonLocked = value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
