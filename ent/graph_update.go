// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/gobench-io/gobench/ent/graph"
	"github.com/gobench-io/gobench/ent/group"
	"github.com/gobench-io/gobench/ent/metric"
	"github.com/gobench-io/gobench/ent/predicate"
)

// GraphUpdate is the builder for updating Graph entities.
type GraphUpdate struct {
	config
	hooks      []Hook
	mutation   *GraphMutation
	predicates []predicate.Graph
}

// Where adds a new predicate for the builder.
func (gu *GraphUpdate) Where(ps ...predicate.Graph) *GraphUpdate {
	gu.predicates = append(gu.predicates, ps...)
	return gu
}

// SetUnit sets the unit field.
func (gu *GraphUpdate) SetUnit(s string) *GraphUpdate {
	gu.mutation.SetUnit(s)
	return gu
}

// SetGroupID sets the group edge to Group by id.
func (gu *GraphUpdate) SetGroupID(id int) *GraphUpdate {
	gu.mutation.SetGroupID(id)
	return gu
}

// SetNillableGroupID sets the group edge to Group by id if the given value is not nil.
func (gu *GraphUpdate) SetNillableGroupID(id *int) *GraphUpdate {
	if id != nil {
		gu = gu.SetGroupID(*id)
	}
	return gu
}

// SetGroup sets the group edge to Group.
func (gu *GraphUpdate) SetGroup(g *Group) *GraphUpdate {
	return gu.SetGroupID(g.ID)
}

// AddMetricIDs adds the metrics edge to Metric by ids.
func (gu *GraphUpdate) AddMetricIDs(ids ...int) *GraphUpdate {
	gu.mutation.AddMetricIDs(ids...)
	return gu
}

// AddMetrics adds the metrics edges to Metric.
func (gu *GraphUpdate) AddMetrics(m ...*Metric) *GraphUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gu.AddMetricIDs(ids...)
}

// ClearGroup clears the group edge to Group.
func (gu *GraphUpdate) ClearGroup() *GraphUpdate {
	gu.mutation.ClearGroup()
	return gu
}

// RemoveMetricIDs removes the metrics edge to Metric by ids.
func (gu *GraphUpdate) RemoveMetricIDs(ids ...int) *GraphUpdate {
	gu.mutation.RemoveMetricIDs(ids...)
	return gu
}

// RemoveMetrics removes metrics edges to Metric.
func (gu *GraphUpdate) RemoveMetrics(m ...*Metric) *GraphUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gu.RemoveMetricIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (gu *GraphUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GraphMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GraphUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GraphUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GraphUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GraphUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   graph.Table,
			Columns: graph.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: graph.FieldID,
			},
		},
	}
	if ps := gu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Unit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: graph.FieldUnit,
		})
	}
	if gu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   graph.GroupTable,
			Columns: []string{graph.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   graph.GroupTable,
			Columns: []string{graph.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := gu.mutation.RemovedMetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   graph.MetricsTable,
			Columns: []string{graph.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.MetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   graph.MetricsTable,
			Columns: []string{graph.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{graph.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GraphUpdateOne is the builder for updating a single Graph entity.
type GraphUpdateOne struct {
	config
	hooks    []Hook
	mutation *GraphMutation
}

// SetUnit sets the unit field.
func (guo *GraphUpdateOne) SetUnit(s string) *GraphUpdateOne {
	guo.mutation.SetUnit(s)
	return guo
}

// SetGroupID sets the group edge to Group by id.
func (guo *GraphUpdateOne) SetGroupID(id int) *GraphUpdateOne {
	guo.mutation.SetGroupID(id)
	return guo
}

// SetNillableGroupID sets the group edge to Group by id if the given value is not nil.
func (guo *GraphUpdateOne) SetNillableGroupID(id *int) *GraphUpdateOne {
	if id != nil {
		guo = guo.SetGroupID(*id)
	}
	return guo
}

// SetGroup sets the group edge to Group.
func (guo *GraphUpdateOne) SetGroup(g *Group) *GraphUpdateOne {
	return guo.SetGroupID(g.ID)
}

// AddMetricIDs adds the metrics edge to Metric by ids.
func (guo *GraphUpdateOne) AddMetricIDs(ids ...int) *GraphUpdateOne {
	guo.mutation.AddMetricIDs(ids...)
	return guo
}

// AddMetrics adds the metrics edges to Metric.
func (guo *GraphUpdateOne) AddMetrics(m ...*Metric) *GraphUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return guo.AddMetricIDs(ids...)
}

// ClearGroup clears the group edge to Group.
func (guo *GraphUpdateOne) ClearGroup() *GraphUpdateOne {
	guo.mutation.ClearGroup()
	return guo
}

// RemoveMetricIDs removes the metrics edge to Metric by ids.
func (guo *GraphUpdateOne) RemoveMetricIDs(ids ...int) *GraphUpdateOne {
	guo.mutation.RemoveMetricIDs(ids...)
	return guo
}

// RemoveMetrics removes metrics edges to Metric.
func (guo *GraphUpdateOne) RemoveMetrics(m ...*Metric) *GraphUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return guo.RemoveMetricIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (guo *GraphUpdateOne) Save(ctx context.Context) (*Graph, error) {

	var (
		err  error
		node *Graph
	)
	if len(guo.hooks) == 0 {
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GraphMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GraphUpdateOne) SaveX(ctx context.Context) *Graph {
	gr, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return gr
}

// Exec executes the query on the entity.
func (guo *GraphUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GraphUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GraphUpdateOne) sqlSave(ctx context.Context) (gr *Graph, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   graph.Table,
			Columns: graph.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: graph.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Graph.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := guo.mutation.Unit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: graph.FieldUnit,
		})
	}
	if guo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   graph.GroupTable,
			Columns: []string{graph.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   graph.GroupTable,
			Columns: []string{graph.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := guo.mutation.RemovedMetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   graph.MetricsTable,
			Columns: []string{graph.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.MetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   graph.MetricsTable,
			Columns: []string{graph.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	gr = &Graph{config: guo.config}
	_spec.Assign = gr.assignValues
	_spec.ScanValues = gr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{graph.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return gr, nil
}