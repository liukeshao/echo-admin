// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/liukeshao/echo-admin/ent/cluster"
)

// ClusterCreate is the builder for creating a Cluster entity.
type ClusterCreate struct {
	config
	mutation *ClusterMutation
	hooks    []Hook
}

// SetDeleteTime sets the "delete_time" field.
func (cc *ClusterCreate) SetDeleteTime(i int64) *ClusterCreate {
	cc.mutation.SetDeleteTime(i)
	return cc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (cc *ClusterCreate) SetNillableDeleteTime(i *int64) *ClusterCreate {
	if i != nil {
		cc.SetDeleteTime(*i)
	}
	return cc
}

// SetCreateTime sets the "create_time" field.
func (cc *ClusterCreate) SetCreateTime(t time.Time) *ClusterCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *ClusterCreate) SetNillableCreateTime(t *time.Time) *ClusterCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *ClusterCreate) SetUpdateTime(t time.Time) *ClusterCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *ClusterCreate) SetNillableUpdateTime(t *time.Time) *ClusterCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetOrgID sets the "org_id" field.
func (cc *ClusterCreate) SetOrgID(u uint64) *ClusterCreate {
	cc.mutation.SetOrgID(u)
	return cc
}

// SetName sets the "name" field.
func (cc *ClusterCreate) SetName(s string) *ClusterCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetDisplayName sets the "display_name" field.
func (cc *ClusterCreate) SetDisplayName(s string) *ClusterCreate {
	cc.mutation.SetDisplayName(s)
	return cc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (cc *ClusterCreate) SetNillableDisplayName(s *string) *ClusterCreate {
	if s != nil {
		cc.SetDisplayName(*s)
	}
	return cc
}

// SetDescription sets the "description" field.
func (cc *ClusterCreate) SetDescription(s string) *ClusterCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *ClusterCreate) SetNillableDescription(s *string) *ClusterCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// SetConfig sets the "config" field.
func (cc *ClusterCreate) SetConfig(s string) *ClusterCreate {
	cc.mutation.SetConfig(s)
	return cc
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (cc *ClusterCreate) SetNillableConfig(s *string) *ClusterCreate {
	if s != nil {
		cc.SetConfig(*s)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ClusterCreate) SetID(u uint64) *ClusterCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ClusterCreate) SetNillableID(u *uint64) *ClusterCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// Mutation returns the ClusterMutation object of the builder.
func (cc *ClusterCreate) Mutation() *ClusterMutation {
	return cc.mutation
}

// Save creates the Cluster in the database.
func (cc *ClusterCreate) Save(ctx context.Context) (*Cluster, error) {
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ClusterCreate) SaveX(ctx context.Context) *Cluster {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ClusterCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ClusterCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ClusterCreate) defaults() error {
	if _, ok := cc.mutation.DeleteTime(); !ok {
		v := cluster.DefaultDeleteTime
		cc.mutation.SetDeleteTime(v)
	}
	if _, ok := cc.mutation.CreateTime(); !ok {
		if cluster.DefaultCreateTime == nil {
			return fmt.Errorf("ent: uninitialized cluster.DefaultCreateTime (forgotten import ent/runtime?)")
		}
		v := cluster.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		if cluster.DefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized cluster.DefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := cluster.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		if cluster.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized cluster.DefaultID (forgotten import ent/runtime?)")
		}
		v := cluster.DefaultID()
		cc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *ClusterCreate) check() error {
	if _, ok := cc.mutation.DeleteTime(); !ok {
		return &ValidationError{Name: "delete_time", err: errors.New(`ent: missing required field "Cluster.delete_time"`)}
	}
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Cluster.create_time"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Cluster.update_time"`)}
	}
	if _, ok := cc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org_id", err: errors.New(`ent: missing required field "Cluster.org_id"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Cluster.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := cluster.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Cluster.name": %w`, err)}
		}
	}
	return nil
}

func (cc *ClusterCreate) sqlSave(ctx context.Context) (*Cluster, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ClusterCreate) createSpec() (*Cluster, *sqlgraph.CreateSpec) {
	var (
		_node = &Cluster{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(cluster.Table, sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeUint64))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.DeleteTime(); ok {
		_spec.SetField(cluster.FieldDeleteTime, field.TypeInt64, value)
		_node.DeleteTime = value
	}
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.SetField(cluster.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.SetField(cluster.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.OrgID(); ok {
		_spec.SetField(cluster.FieldOrgID, field.TypeUint64, value)
		_node.OrgID = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(cluster.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.DisplayName(); ok {
		_spec.SetField(cluster.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(cluster.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cc.mutation.Config(); ok {
		_spec.SetField(cluster.FieldConfig, field.TypeString, value)
		_node.Config = value
	}
	return _node, _spec
}

// ClusterCreateBulk is the builder for creating many Cluster entities in bulk.
type ClusterCreateBulk struct {
	config
	err      error
	builders []*ClusterCreate
}

// Save creates the Cluster entities in the database.
func (ccb *ClusterCreateBulk) Save(ctx context.Context) ([]*Cluster, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Cluster, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClusterMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ClusterCreateBulk) SaveX(ctx context.Context) []*Cluster {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ClusterCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ClusterCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}