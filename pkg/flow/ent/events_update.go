// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/events"
	"github.com/direktiv/direktiv/pkg/flow/ent/eventswait"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// EventsUpdate is the builder for updating Events entities.
type EventsUpdate struct {
	config
	hooks    []Hook
	mutation *EventsMutation
}

// Where appends a list predicates to the EventsUpdate builder.
func (eu *EventsUpdate) Where(ps ...predicate.Events) *EventsUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetEvents sets the "events" field.
func (eu *EventsUpdate) SetEvents(m []map[string]interface{}) *EventsUpdate {
	eu.mutation.SetEvents(m)
	return eu
}

// SetCorrelations sets the "correlations" field.
func (eu *EventsUpdate) SetCorrelations(s []string) *EventsUpdate {
	eu.mutation.SetCorrelations(s)
	return eu
}

// SetSignature sets the "signature" field.
func (eu *EventsUpdate) SetSignature(b []byte) *EventsUpdate {
	eu.mutation.SetSignature(b)
	return eu
}

// ClearSignature clears the value of the "signature" field.
func (eu *EventsUpdate) ClearSignature() *EventsUpdate {
	eu.mutation.ClearSignature()
	return eu
}

// SetCount sets the "count" field.
func (eu *EventsUpdate) SetCount(i int) *EventsUpdate {
	eu.mutation.ResetCount()
	eu.mutation.SetCount(i)
	return eu
}

// AddCount adds i to the "count" field.
func (eu *EventsUpdate) AddCount(i int) *EventsUpdate {
	eu.mutation.AddCount(i)
	return eu
}

// SetUpdatedAt sets the "updated_at" field.
func (eu *EventsUpdate) SetUpdatedAt(t time.Time) *EventsUpdate {
	eu.mutation.SetUpdatedAt(t)
	return eu
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (eu *EventsUpdate) SetWorkflowID(id uuid.UUID) *EventsUpdate {
	eu.mutation.SetWorkflowID(id)
	return eu
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (eu *EventsUpdate) SetWorkflow(w *Workflow) *EventsUpdate {
	return eu.SetWorkflowID(w.ID)
}

// AddWfeventswaitIDs adds the "wfeventswait" edge to the EventsWait entity by IDs.
func (eu *EventsUpdate) AddWfeventswaitIDs(ids ...uuid.UUID) *EventsUpdate {
	eu.mutation.AddWfeventswaitIDs(ids...)
	return eu
}

// AddWfeventswait adds the "wfeventswait" edges to the EventsWait entity.
func (eu *EventsUpdate) AddWfeventswait(e ...*EventsWait) *EventsUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return eu.AddWfeventswaitIDs(ids...)
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (eu *EventsUpdate) SetInstanceID(id uuid.UUID) *EventsUpdate {
	eu.mutation.SetInstanceID(id)
	return eu
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (eu *EventsUpdate) SetNillableInstanceID(id *uuid.UUID) *EventsUpdate {
	if id != nil {
		eu = eu.SetInstanceID(*id)
	}
	return eu
}

// SetInstance sets the "instance" edge to the Instance entity.
func (eu *EventsUpdate) SetInstance(i *Instance) *EventsUpdate {
	return eu.SetInstanceID(i.ID)
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (eu *EventsUpdate) SetNamespaceID(id uuid.UUID) *EventsUpdate {
	eu.mutation.SetNamespaceID(id)
	return eu
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (eu *EventsUpdate) SetNamespace(n *Namespace) *EventsUpdate {
	return eu.SetNamespaceID(n.ID)
}

// Mutation returns the EventsMutation object of the builder.
func (eu *EventsUpdate) Mutation() *EventsMutation {
	return eu.mutation
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (eu *EventsUpdate) ClearWorkflow() *EventsUpdate {
	eu.mutation.ClearWorkflow()
	return eu
}

// ClearWfeventswait clears all "wfeventswait" edges to the EventsWait entity.
func (eu *EventsUpdate) ClearWfeventswait() *EventsUpdate {
	eu.mutation.ClearWfeventswait()
	return eu
}

// RemoveWfeventswaitIDs removes the "wfeventswait" edge to EventsWait entities by IDs.
func (eu *EventsUpdate) RemoveWfeventswaitIDs(ids ...uuid.UUID) *EventsUpdate {
	eu.mutation.RemoveWfeventswaitIDs(ids...)
	return eu
}

// RemoveWfeventswait removes "wfeventswait" edges to EventsWait entities.
func (eu *EventsUpdate) RemoveWfeventswait(e ...*EventsWait) *EventsUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return eu.RemoveWfeventswaitIDs(ids...)
}

// ClearInstance clears the "instance" edge to the Instance entity.
func (eu *EventsUpdate) ClearInstance() *EventsUpdate {
	eu.mutation.ClearInstance()
	return eu
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (eu *EventsUpdate) ClearNamespace() *EventsUpdate {
	eu.mutation.ClearNamespace()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EventsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	eu.defaults()
	if len(eu.hooks) == 0 {
		if err = eu.check(); err != nil {
			return 0, err
		}
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eu.check(); err != nil {
				return 0, err
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventsUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventsUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventsUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *EventsUpdate) defaults() {
	if _, ok := eu.mutation.UpdatedAt(); !ok {
		v := events.UpdateDefaultUpdatedAt()
		eu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EventsUpdate) check() error {
	if _, ok := eu.mutation.WorkflowID(); eu.mutation.WorkflowCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Events.workflow"`)
	}
	if _, ok := eu.mutation.NamespaceID(); eu.mutation.NamespaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Events.namespace"`)
	}
	return nil
}

func (eu *EventsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   events.Table,
			Columns: events.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: events.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Events(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: events.FieldEvents,
		})
	}
	if value, ok := eu.mutation.Correlations(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: events.FieldCorrelations,
		})
	}
	if value, ok := eu.mutation.Signature(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: events.FieldSignature,
		})
	}
	if eu.mutation.SignatureCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: events.FieldSignature,
		})
	}
	if value, ok := eu.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: events.FieldCount,
		})
	}
	if value, ok := eu.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: events.FieldCount,
		})
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: events.FieldUpdatedAt,
		})
	}
	if eu.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.WorkflowTable,
			Columns: []string{events.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.WorkflowTable,
			Columns: []string{events.WorkflowColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.WfeventswaitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   events.WfeventswaitTable,
			Columns: []string{events.WfeventswaitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: eventswait.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedWfeventswaitIDs(); len(nodes) > 0 && !eu.mutation.WfeventswaitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   events.WfeventswaitTable,
			Columns: []string{events.WfeventswaitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: eventswait.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.WfeventswaitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   events.WfeventswaitTable,
			Columns: []string{events.WfeventswaitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: eventswait.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.InstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.InstanceTable,
			Columns: []string{events.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.InstanceTable,
			Columns: []string{events.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.NamespaceTable,
			Columns: []string{events.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.NamespaceTable,
			Columns: []string{events.NamespaceColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{events.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// EventsUpdateOne is the builder for updating a single Events entity.
type EventsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventsMutation
}

// SetEvents sets the "events" field.
func (euo *EventsUpdateOne) SetEvents(m []map[string]interface{}) *EventsUpdateOne {
	euo.mutation.SetEvents(m)
	return euo
}

// SetCorrelations sets the "correlations" field.
func (euo *EventsUpdateOne) SetCorrelations(s []string) *EventsUpdateOne {
	euo.mutation.SetCorrelations(s)
	return euo
}

// SetSignature sets the "signature" field.
func (euo *EventsUpdateOne) SetSignature(b []byte) *EventsUpdateOne {
	euo.mutation.SetSignature(b)
	return euo
}

// ClearSignature clears the value of the "signature" field.
func (euo *EventsUpdateOne) ClearSignature() *EventsUpdateOne {
	euo.mutation.ClearSignature()
	return euo
}

// SetCount sets the "count" field.
func (euo *EventsUpdateOne) SetCount(i int) *EventsUpdateOne {
	euo.mutation.ResetCount()
	euo.mutation.SetCount(i)
	return euo
}

// AddCount adds i to the "count" field.
func (euo *EventsUpdateOne) AddCount(i int) *EventsUpdateOne {
	euo.mutation.AddCount(i)
	return euo
}

// SetUpdatedAt sets the "updated_at" field.
func (euo *EventsUpdateOne) SetUpdatedAt(t time.Time) *EventsUpdateOne {
	euo.mutation.SetUpdatedAt(t)
	return euo
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (euo *EventsUpdateOne) SetWorkflowID(id uuid.UUID) *EventsUpdateOne {
	euo.mutation.SetWorkflowID(id)
	return euo
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (euo *EventsUpdateOne) SetWorkflow(w *Workflow) *EventsUpdateOne {
	return euo.SetWorkflowID(w.ID)
}

// AddWfeventswaitIDs adds the "wfeventswait" edge to the EventsWait entity by IDs.
func (euo *EventsUpdateOne) AddWfeventswaitIDs(ids ...uuid.UUID) *EventsUpdateOne {
	euo.mutation.AddWfeventswaitIDs(ids...)
	return euo
}

// AddWfeventswait adds the "wfeventswait" edges to the EventsWait entity.
func (euo *EventsUpdateOne) AddWfeventswait(e ...*EventsWait) *EventsUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return euo.AddWfeventswaitIDs(ids...)
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (euo *EventsUpdateOne) SetInstanceID(id uuid.UUID) *EventsUpdateOne {
	euo.mutation.SetInstanceID(id)
	return euo
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (euo *EventsUpdateOne) SetNillableInstanceID(id *uuid.UUID) *EventsUpdateOne {
	if id != nil {
		euo = euo.SetInstanceID(*id)
	}
	return euo
}

// SetInstance sets the "instance" edge to the Instance entity.
func (euo *EventsUpdateOne) SetInstance(i *Instance) *EventsUpdateOne {
	return euo.SetInstanceID(i.ID)
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (euo *EventsUpdateOne) SetNamespaceID(id uuid.UUID) *EventsUpdateOne {
	euo.mutation.SetNamespaceID(id)
	return euo
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (euo *EventsUpdateOne) SetNamespace(n *Namespace) *EventsUpdateOne {
	return euo.SetNamespaceID(n.ID)
}

// Mutation returns the EventsMutation object of the builder.
func (euo *EventsUpdateOne) Mutation() *EventsMutation {
	return euo.mutation
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (euo *EventsUpdateOne) ClearWorkflow() *EventsUpdateOne {
	euo.mutation.ClearWorkflow()
	return euo
}

// ClearWfeventswait clears all "wfeventswait" edges to the EventsWait entity.
func (euo *EventsUpdateOne) ClearWfeventswait() *EventsUpdateOne {
	euo.mutation.ClearWfeventswait()
	return euo
}

// RemoveWfeventswaitIDs removes the "wfeventswait" edge to EventsWait entities by IDs.
func (euo *EventsUpdateOne) RemoveWfeventswaitIDs(ids ...uuid.UUID) *EventsUpdateOne {
	euo.mutation.RemoveWfeventswaitIDs(ids...)
	return euo
}

// RemoveWfeventswait removes "wfeventswait" edges to EventsWait entities.
func (euo *EventsUpdateOne) RemoveWfeventswait(e ...*EventsWait) *EventsUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return euo.RemoveWfeventswaitIDs(ids...)
}

// ClearInstance clears the "instance" edge to the Instance entity.
func (euo *EventsUpdateOne) ClearInstance() *EventsUpdateOne {
	euo.mutation.ClearInstance()
	return euo
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (euo *EventsUpdateOne) ClearNamespace() *EventsUpdateOne {
	euo.mutation.ClearNamespace()
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EventsUpdateOne) Select(field string, fields ...string) *EventsUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Events entity.
func (euo *EventsUpdateOne) Save(ctx context.Context) (*Events, error) {
	var (
		err  error
		node *Events
	)
	euo.defaults()
	if len(euo.hooks) == 0 {
		if err = euo.check(); err != nil {
			return nil, err
		}
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = euo.check(); err != nil {
				return nil, err
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventsUpdateOne) SaveX(ctx context.Context) *Events {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventsUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventsUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *EventsUpdateOne) defaults() {
	if _, ok := euo.mutation.UpdatedAt(); !ok {
		v := events.UpdateDefaultUpdatedAt()
		euo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EventsUpdateOne) check() error {
	if _, ok := euo.mutation.WorkflowID(); euo.mutation.WorkflowCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Events.workflow"`)
	}
	if _, ok := euo.mutation.NamespaceID(); euo.mutation.NamespaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Events.namespace"`)
	}
	return nil
}

func (euo *EventsUpdateOne) sqlSave(ctx context.Context) (_node *Events, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   events.Table,
			Columns: events.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: events.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Events.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, events.FieldID)
		for _, f := range fields {
			if !events.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != events.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Events(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: events.FieldEvents,
		})
	}
	if value, ok := euo.mutation.Correlations(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: events.FieldCorrelations,
		})
	}
	if value, ok := euo.mutation.Signature(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: events.FieldSignature,
		})
	}
	if euo.mutation.SignatureCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: events.FieldSignature,
		})
	}
	if value, ok := euo.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: events.FieldCount,
		})
	}
	if value, ok := euo.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: events.FieldCount,
		})
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: events.FieldUpdatedAt,
		})
	}
	if euo.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.WorkflowTable,
			Columns: []string{events.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.WorkflowTable,
			Columns: []string{events.WorkflowColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.WfeventswaitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   events.WfeventswaitTable,
			Columns: []string{events.WfeventswaitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: eventswait.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedWfeventswaitIDs(); len(nodes) > 0 && !euo.mutation.WfeventswaitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   events.WfeventswaitTable,
			Columns: []string{events.WfeventswaitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: eventswait.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.WfeventswaitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   events.WfeventswaitTable,
			Columns: []string{events.WfeventswaitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: eventswait.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.InstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.InstanceTable,
			Columns: []string{events.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.InstanceTable,
			Columns: []string{events.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.NamespaceTable,
			Columns: []string{events.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.NamespaceTable,
			Columns: []string{events.NamespaceColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Events{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{events.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
