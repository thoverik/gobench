// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/gobench-io/gobench/ent/gauge"
	"github.com/gobench-io/gobench/ent/predicate"
)

// GaugeDelete is the builder for deleting a Gauge entity.
type GaugeDelete struct {
	config
	hooks      []Hook
	mutation   *GaugeMutation
	predicates []predicate.Gauge
}

// Where adds a new predicate to the delete builder.
func (gd *GaugeDelete) Where(ps ...predicate.Gauge) *GaugeDelete {
	gd.predicates = append(gd.predicates, ps...)
	return gd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gd *GaugeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gd.hooks) == 0 {
		affected, err = gd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GaugeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gd.mutation = mutation
			affected, err = gd.sqlExec(ctx)
			return affected, err
		})
		for i := len(gd.hooks) - 1; i >= 0; i-- {
			mut = gd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gd *GaugeDelete) ExecX(ctx context.Context) int {
	n, err := gd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gd *GaugeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: gauge.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gauge.FieldID,
			},
		},
	}
	if ps := gd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, gd.driver, _spec)
}

// GaugeDeleteOne is the builder for deleting a single Gauge entity.
type GaugeDeleteOne struct {
	gd *GaugeDelete
}

// Exec executes the deletion query.
func (gdo *GaugeDeleteOne) Exec(ctx context.Context) error {
	n, err := gdo.gd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{gauge.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gdo *GaugeDeleteOne) ExecX(ctx context.Context) {
	gdo.gd.ExecX(ctx)
}