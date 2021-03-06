// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/thoverik/gobench/ent/counter"
	"github.com/thoverik/gobench/ent/metric"
)

// CounterCreate is the builder for creating a Counter entity.
type CounterCreate struct {
	config
	mutation *CounterMutation
	hooks    []Hook
}

// SetTime sets the time field.
func (cc *CounterCreate) SetTime(i int64) *CounterCreate {
	cc.mutation.SetTime(i)
	return cc
}

// SetCount sets the count field.
func (cc *CounterCreate) SetCount(i int64) *CounterCreate {
	cc.mutation.SetCount(i)
	return cc
}

// SetMetricID sets the metric edge to Metric by id.
func (cc *CounterCreate) SetMetricID(id int) *CounterCreate {
	cc.mutation.SetMetricID(id)
	return cc
}

// SetNillableMetricID sets the metric edge to Metric by id if the given value is not nil.
func (cc *CounterCreate) SetNillableMetricID(id *int) *CounterCreate {
	if id != nil {
		cc = cc.SetMetricID(*id)
	}
	return cc
}

// SetMetric sets the metric edge to Metric.
func (cc *CounterCreate) SetMetric(m *Metric) *CounterCreate {
	return cc.SetMetricID(m.ID)
}

// Save creates the Counter in the database.
func (cc *CounterCreate) Save(ctx context.Context) (*Counter, error) {
	if _, ok := cc.mutation.Time(); !ok {
		return nil, errors.New("ent: missing required field \"time\"")
	}
	if _, ok := cc.mutation.Count(); !ok {
		return nil, errors.New("ent: missing required field \"count\"")
	}
	var (
		err  error
		node *Counter
	)
	if len(cc.hooks) == 0 {
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CounterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CounterCreate) SaveX(ctx context.Context) *Counter {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CounterCreate) sqlSave(ctx context.Context) (*Counter, error) {
	var (
		c     = &Counter{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: counter.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: counter.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: counter.FieldTime,
		})
		c.Time = value
	}
	if value, ok := cc.mutation.Count(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: counter.FieldCount,
		})
		c.Count = value
	}
	if nodes := cc.mutation.MetricIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	c.ID = int(id)
	return c, nil
}
