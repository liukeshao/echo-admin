// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/liukeshao/echo-admin/ent/predicate"
	"github.com/liukeshao/echo-admin/ent/rolebinding"
)

// RoleBindingDelete is the builder for deleting a RoleBinding entity.
type RoleBindingDelete struct {
	config
	hooks    []Hook
	mutation *RoleBindingMutation
}

// Where appends a list predicates to the RoleBindingDelete builder.
func (rbd *RoleBindingDelete) Where(ps ...predicate.RoleBinding) *RoleBindingDelete {
	rbd.mutation.Where(ps...)
	return rbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rbd *RoleBindingDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rbd.sqlExec, rbd.mutation, rbd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rbd *RoleBindingDelete) ExecX(ctx context.Context) int {
	n, err := rbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rbd *RoleBindingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(rolebinding.Table, sqlgraph.NewFieldSpec(rolebinding.FieldID, field.TypeInt64))
	if ps := rbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rbd.mutation.done = true
	return affected, err
}

// RoleBindingDeleteOne is the builder for deleting a single RoleBinding entity.
type RoleBindingDeleteOne struct {
	rbd *RoleBindingDelete
}

// Where appends a list predicates to the RoleBindingDelete builder.
func (rbdo *RoleBindingDeleteOne) Where(ps ...predicate.RoleBinding) *RoleBindingDeleteOne {
	rbdo.rbd.mutation.Where(ps...)
	return rbdo
}

// Exec executes the deletion query.
func (rbdo *RoleBindingDeleteOne) Exec(ctx context.Context) error {
	n, err := rbdo.rbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rolebinding.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rbdo *RoleBindingDeleteOne) ExecX(ctx context.Context) {
	if err := rbdo.Exec(ctx); err != nil {
		panic(err)
	}
}
