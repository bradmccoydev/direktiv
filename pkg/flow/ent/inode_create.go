// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/inode"
	"github.com/direktiv/direktiv/pkg/flow/ent/mirror"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// InodeCreate is the builder for creating a Inode entity.
type InodeCreate struct {
	config
	mutation *InodeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ic *InodeCreate) SetCreatedAt(t time.Time) *InodeCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *InodeCreate) SetNillableCreatedAt(t *time.Time) *InodeCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *InodeCreate) SetUpdatedAt(t time.Time) *InodeCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *InodeCreate) SetNillableUpdatedAt(t *time.Time) *InodeCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetName sets the "name" field.
func (ic *InodeCreate) SetName(s string) *InodeCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ic *InodeCreate) SetNillableName(s *string) *InodeCreate {
	if s != nil {
		ic.SetName(*s)
	}
	return ic
}

// SetType sets the "type" field.
func (ic *InodeCreate) SetType(s string) *InodeCreate {
	ic.mutation.SetType(s)
	return ic
}

// SetAttributes sets the "attributes" field.
func (ic *InodeCreate) SetAttributes(s []string) *InodeCreate {
	ic.mutation.SetAttributes(s)
	return ic
}

// SetExtendedType sets the "extended_type" field.
func (ic *InodeCreate) SetExtendedType(s string) *InodeCreate {
	ic.mutation.SetExtendedType(s)
	return ic
}

// SetNillableExtendedType sets the "extended_type" field if the given value is not nil.
func (ic *InodeCreate) SetNillableExtendedType(s *string) *InodeCreate {
	if s != nil {
		ic.SetExtendedType(*s)
	}
	return ic
}

// SetReadOnly sets the "readOnly" field.
func (ic *InodeCreate) SetReadOnly(b bool) *InodeCreate {
	ic.mutation.SetReadOnly(b)
	return ic
}

// SetNillableReadOnly sets the "readOnly" field if the given value is not nil.
func (ic *InodeCreate) SetNillableReadOnly(b *bool) *InodeCreate {
	if b != nil {
		ic.SetReadOnly(*b)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *InodeCreate) SetID(u uuid.UUID) *InodeCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *InodeCreate) SetNillableID(u *uuid.UUID) *InodeCreate {
	if u != nil {
		ic.SetID(*u)
	}
	return ic
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (ic *InodeCreate) SetNamespaceID(id uuid.UUID) *InodeCreate {
	ic.mutation.SetNamespaceID(id)
	return ic
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (ic *InodeCreate) SetNamespace(n *Namespace) *InodeCreate {
	return ic.SetNamespaceID(n.ID)
}

// AddChildIDs adds the "children" edge to the Inode entity by IDs.
func (ic *InodeCreate) AddChildIDs(ids ...uuid.UUID) *InodeCreate {
	ic.mutation.AddChildIDs(ids...)
	return ic
}

// AddChildren adds the "children" edges to the Inode entity.
func (ic *InodeCreate) AddChildren(i ...*Inode) *InodeCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ic.AddChildIDs(ids...)
}

// SetParentID sets the "parent" edge to the Inode entity by ID.
func (ic *InodeCreate) SetParentID(id uuid.UUID) *InodeCreate {
	ic.mutation.SetParentID(id)
	return ic
}

// SetNillableParentID sets the "parent" edge to the Inode entity by ID if the given value is not nil.
func (ic *InodeCreate) SetNillableParentID(id *uuid.UUID) *InodeCreate {
	if id != nil {
		ic = ic.SetParentID(*id)
	}
	return ic
}

// SetParent sets the "parent" edge to the Inode entity.
func (ic *InodeCreate) SetParent(i *Inode) *InodeCreate {
	return ic.SetParentID(i.ID)
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (ic *InodeCreate) SetWorkflowID(id uuid.UUID) *InodeCreate {
	ic.mutation.SetWorkflowID(id)
	return ic
}

// SetNillableWorkflowID sets the "workflow" edge to the Workflow entity by ID if the given value is not nil.
func (ic *InodeCreate) SetNillableWorkflowID(id *uuid.UUID) *InodeCreate {
	if id != nil {
		ic = ic.SetWorkflowID(*id)
	}
	return ic
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (ic *InodeCreate) SetWorkflow(w *Workflow) *InodeCreate {
	return ic.SetWorkflowID(w.ID)
}

// SetMirrorID sets the "mirror" edge to the Mirror entity by ID.
func (ic *InodeCreate) SetMirrorID(id uuid.UUID) *InodeCreate {
	ic.mutation.SetMirrorID(id)
	return ic
}

// SetNillableMirrorID sets the "mirror" edge to the Mirror entity by ID if the given value is not nil.
func (ic *InodeCreate) SetNillableMirrorID(id *uuid.UUID) *InodeCreate {
	if id != nil {
		ic = ic.SetMirrorID(*id)
	}
	return ic
}

// SetMirror sets the "mirror" edge to the Mirror entity.
func (ic *InodeCreate) SetMirror(m *Mirror) *InodeCreate {
	return ic.SetMirrorID(m.ID)
}

// Mutation returns the InodeMutation object of the builder.
func (ic *InodeCreate) Mutation() *InodeMutation {
	return ic.mutation
}

// Save creates the Inode in the database.
func (ic *InodeCreate) Save(ctx context.Context) (*Inode, error) {
	var (
		err  error
		node *Inode
	)
	ic.defaults()
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InodeCreate) SaveX(ctx context.Context) *Inode {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InodeCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InodeCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *InodeCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := inode.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := inode.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
	if _, ok := ic.mutation.ReadOnly(); !ok {
		v := inode.DefaultReadOnly
		ic.mutation.SetReadOnly(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := inode.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InodeCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Inode.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Inode.updated_at"`)}
	}
	if v, ok := ic.mutation.Name(); ok {
		if err := inode.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Inode.name": %w`, err)}
		}
	}
	if _, ok := ic.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Inode.type"`)}
	}
	if _, ok := ic.mutation.NamespaceID(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`ent: missing required edge "Inode.namespace"`)}
	}
	return nil
}

func (ic *InodeCreate) sqlSave(ctx context.Context) (*Inode, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ic *InodeCreate) createSpec() (*Inode, *sqlgraph.CreateSpec) {
	var (
		_node = &Inode{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: inode.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: inode.FieldID,
			},
		}
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: inode.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: inode.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inode.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ic.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inode.FieldType,
		})
		_node.Type = value
	}
	if value, ok := ic.mutation.Attributes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: inode.FieldAttributes,
		})
		_node.Attributes = value
	}
	if value, ok := ic.mutation.ExtendedType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inode.FieldExtendedType,
		})
		_node.ExtendedType = value
	}
	if value, ok := ic.mutation.ReadOnly(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: inode.FieldReadOnly,
		})
		_node.ReadOnly = value
	}
	if nodes := ic.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inode.NamespaceTable,
			Columns: []string{inode.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.namespace_inodes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inode.ChildrenTable,
			Columns: []string{inode.ChildrenColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: inode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inode.ParentTable,
			Columns: []string{inode.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: inode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.inode_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   inode.WorkflowTable,
			Columns: []string{inode.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.MirrorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   inode.MirrorTable,
			Columns: []string{inode.MirrorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mirror.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InodeCreateBulk is the builder for creating many Inode entities in bulk.
type InodeCreateBulk struct {
	config
	builders []*InodeCreate
}

// Save creates the Inode entities in the database.
func (icb *InodeCreateBulk) Save(ctx context.Context) ([]*Inode, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Inode, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InodeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InodeCreateBulk) SaveX(ctx context.Context) []*Inode {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *InodeCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *InodeCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
