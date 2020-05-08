// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/gobench-io/gobench/ent/graph"
	"github.com/gobench-io/gobench/ent/metric"
)

// Metric is the model entity for the Metric schema.
type Metric struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title"`
	// Type holds the value of the "type" field.
	Type string `json:"type"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MetricQuery when eager-loading is set.
	Edges         MetricEdges `json:"edges"`
	graph_metrics *int
}

// MetricEdges holds the relations/edges for other nodes in the graph.
type MetricEdges struct {
	// Graph holds the value of the graph edge.
	Graph *Graph `json:"-"`
	// Histograms holds the value of the histograms edge.
	Histograms []*Histogram
	// Counters holds the value of the counters edge.
	Counters []*Counter
	// Gauges holds the value of the gauges edge.
	Gauges []*Gauge
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// GraphOrErr returns the Graph value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MetricEdges) GraphOrErr() (*Graph, error) {
	if e.loadedTypes[0] {
		if e.Graph == nil {
			// The edge graph was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: graph.Label}
		}
		return e.Graph, nil
	}
	return nil, &NotLoadedError{edge: "graph"}
}

// HistogramsOrErr returns the Histograms value or an error if the edge
// was not loaded in eager-loading.
func (e MetricEdges) HistogramsOrErr() ([]*Histogram, error) {
	if e.loadedTypes[1] {
		return e.Histograms, nil
	}
	return nil, &NotLoadedError{edge: "histograms"}
}

// CountersOrErr returns the Counters value or an error if the edge
// was not loaded in eager-loading.
func (e MetricEdges) CountersOrErr() ([]*Counter, error) {
	if e.loadedTypes[2] {
		return e.Counters, nil
	}
	return nil, &NotLoadedError{edge: "counters"}
}

// GaugesOrErr returns the Gauges value or an error if the edge
// was not loaded in eager-loading.
func (e MetricEdges) GaugesOrErr() ([]*Gauge, error) {
	if e.loadedTypes[3] {
		return e.Gauges, nil
	}
	return nil, &NotLoadedError{edge: "gauges"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Metric) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // title
		&sql.NullString{}, // type
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Metric) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // graph_metrics
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Metric fields.
func (m *Metric) assignValues(values ...interface{}) error {
	if m, n := len(values), len(metric.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field title", values[0])
	} else if value.Valid {
		m.Title = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[1])
	} else if value.Valid {
		m.Type = value.String
	}
	values = values[2:]
	if len(values) == len(metric.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field graph_metrics", value)
		} else if value.Valid {
			m.graph_metrics = new(int)
			*m.graph_metrics = int(value.Int64)
		}
	}
	return nil
}

// QueryGraph queries the graph edge of the Metric.
func (m *Metric) QueryGraph() *GraphQuery {
	return (&MetricClient{config: m.config}).QueryGraph(m)
}

// QueryHistograms queries the histograms edge of the Metric.
func (m *Metric) QueryHistograms() *HistogramQuery {
	return (&MetricClient{config: m.config}).QueryHistograms(m)
}

// QueryCounters queries the counters edge of the Metric.
func (m *Metric) QueryCounters() *CounterQuery {
	return (&MetricClient{config: m.config}).QueryCounters(m)
}

// QueryGauges queries the gauges edge of the Metric.
func (m *Metric) QueryGauges() *GaugeQuery {
	return (&MetricClient{config: m.config}).QueryGauges(m)
}

// Update returns a builder for updating this Metric.
// Note that, you need to call Metric.Unwrap() before calling this method, if this Metric
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Metric) Update() *MetricUpdateOne {
	return (&MetricClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Metric) Unwrap() *Metric {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Metric is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Metric) String() string {
	var builder strings.Builder
	builder.WriteString("Metric(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", title=")
	builder.WriteString(m.Title)
	builder.WriteString(", type=")
	builder.WriteString(m.Type)
	builder.WriteByte(')')
	return builder.String()
}

// Metrics is a parsable slice of Metric.
type Metrics []*Metric

func (m Metrics) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
