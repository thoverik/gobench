// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/gobench-io/gobench/ent/counter"
	"github.com/gobench-io/gobench/ent/metric"
	"github.com/gobench-io/gobench/ent/predicate"
)

// CounterUpdate is the builder for updating Counter entities.
type CounterUpdate struct {
	config
	hooks      []Hook
	mutation   *CounterMutation
	predicates []predicate.Counter
}

// Where adds a new predicate for the builder.
func (cu *CounterUpdate) Where(ps ...predicate.Counter) *CounterUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetTime sets the time field.
func (cu *CounterUpdate) SetTime(i int64) *CounterUpdate {
	cu.mutation.ResetTime()
	cu.mutation.SetTime(i)
	return cu
}

// AddTime adds i to time.
func (cu *CounterUpdate) AddTime(i int64) *CounterUpdate {
	cu.mutation.AddTime(i)
	return cu
}

// SetCount sets the count field.
func (cu *CounterUpdate) SetCount(i int64) *CounterUpdate {
	cu.mutation.ResetCount()
	cu.mutation.SetCount(i)
	return cu
}

// AddCount adds i to count.
func (cu *CounterUpdate) AddCount(i int64) *CounterUpdate {
	cu.mutation.AddCount(i)
	return cu
}

// SetMetricID sets the metric edge to Metric by id.
func (cu *CounterUpdate) SetMetricID(id int) *CounterUpdate {
	cu.mutation.SetMetricID(id)
	return cu
}

// SetNillableMetricID sets the metric edge to Metric by id if the given value is not nil.
func (cu *CounterUpdate) SetNillableMetricID(id *int) *CounterUpdate {
	if id != nil {
		cu = cu.SetMetricID(*id)
	}
	return cu
}

// SetMetric sets the metric edge to Metric.
func (cu *CounterUpdate) SetMetric(m *Metric) *CounterUpdate {
	return cu.SetMetricID(m.ID)
}

// ClearMetric clears the metric edge to Metric.
func (cu *CounterUpdate) ClearMetric() *CounterUpdate {
	cu.mutation.ClearMetric()
	return cu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CounterUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CounterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CounterUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CounterUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CounterUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CounterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   counter.Table,
			Columns: counter.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: counter.FieldID,
			},
		},
	}
	if ps := cu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: counter.FieldTime,
		})
	}
	if value, ok := cu.mutation.AddedTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: counter.FieldTime,
		})
	}
	if value, ok := cu.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: counter.FieldCount,
		})
	}
	if value, ok := cu.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: counter.FieldCount,
		})
	}
	if cu.mutation.MetricCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   counter.MetricTable,
			Columns: []string{counter.MetricColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.MetricIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   counter.MetricTable,
			Columns: []string{counter.MetricColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{counter.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CounterUpdateOne is the builder for updating a single Counter entity.
type CounterUpdateOne struct {
	config
	hooks    []Hook
	mutation *CounterMutation
}

// SetTime sets the time field.
func (cuo *CounterUpdateOne) SetTime(i int64) *CounterUpdateOne {
	cuo.mutation.ResetTime()
	cuo.mutation.SetTime(i)
	return cuo
}

// AddTime adds i to time.
func (cuo *CounterUpdateOne) AddTime(i int64) *CounterUpdateOne {
	cuo.mutation.AddTime(i)
	return cuo
}

// SetCount sets the count field.
func (cuo *CounterUpdateOne) SetCount(i int64) *CounterUpdateOne {
	cuo.mutation.ResetCount()
	cuo.mutation.SetCount(i)
	return cuo
}

// AddCount adds i to count.
func (cuo *CounterUpdateOne) AddCount(i int64) *CounterUpdateOne {
	cuo.mutation.AddCount(i)
	return cuo
}

// SetMetricID sets the metric edge to Metric by id.
func (cuo *CounterUpdateOne) SetMetricID(id int) *CounterUpdateOne {
	cuo.mutation.SetMetricID(id)
	return cuo
}

// SetNillableMetricID sets the metric edge to Metric by id if the given value is not nil.
func (cuo *CounterUpdateOne) SetNillableMetricID(id *int) *CounterUpdateOne {
	if id != nil {
		cuo = cuo.SetMetricID(*id)
	}
	return cuo
}

// SetMetric sets the metric edge to Metric.
func (cuo *CounterUpdateOne) SetMetric(m *Metric) *CounterUpdateOne {
	return cuo.SetMetricID(m.ID)
}

// ClearMetric clears the metric edge to Metric.
func (cuo *CounterUpdateOne) ClearMetric() *CounterUpdateOne {
	cuo.mutation.ClearMetric()
	return cuo
}

// Save executes the query and returns the updated entity.
func (cuo *CounterUpdateOne) Save(ctx context.Context) (*Counter, error) {

	var (
		err  error
		node *Counter
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CounterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CounterUpdateOne) SaveX(ctx context.Context) *Counter {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CounterUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CounterUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CounterUpdateOne) sqlSave(ctx context.Context) (c *Counter, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   counter.Table,
			Columns: counter.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: counter.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Counter.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: counter.FieldTime,
		})
	}
	if value, ok := cuo.mutation.AddedTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: counter.FieldTime,
		})
	}
	if value, ok := cuo.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: counter.FieldCount,
		})
	}
	if value, ok := cuo.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: counter.FieldCount,
		})
	}
	if cuo.mutation.MetricCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   counter.MetricTable,
			Columns: []string{counter.MetricColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.MetricIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   counter.MetricTable,
			Columns: []string{counter.MetricColumn},
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
	c = &Counter{config: cuo.config}
	_spec.Assign = c.assignValues
	_spec.ScanValues = c.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{counter.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return c, nil
}
