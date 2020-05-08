// Code generated by entc, DO NOT EDIT.

package metric

const (
	// Label holds the string label denoting the metric type in the database.
	Label = "metric"
	// FieldID holds the string denoting the id field in the database.
	FieldID    = "id"    // FieldTitle holds the string denoting the title vertex property in the database.
	FieldTitle = "title" // FieldType holds the string denoting the type vertex property in the database.
	FieldType  = "type"

	// EdgeGraph holds the string denoting the graph edge name in mutations.
	EdgeGraph = "graph"
	// EdgeHistograms holds the string denoting the histograms edge name in mutations.
	EdgeHistograms = "histograms"
	// EdgeCounters holds the string denoting the counters edge name in mutations.
	EdgeCounters = "counters"
	// EdgeGauges holds the string denoting the gauges edge name in mutations.
	EdgeGauges = "gauges"

	// Table holds the table name of the metric in the database.
	Table = "metrics"
	// GraphTable is the table the holds the graph relation/edge.
	GraphTable = "metrics"
	// GraphInverseTable is the table name for the Graph entity.
	// It exists in this package in order to avoid circular dependency with the "graph" package.
	GraphInverseTable = "graphs"
	// GraphColumn is the table column denoting the graph relation/edge.
	GraphColumn = "graph_metrics"
	// HistogramsTable is the table the holds the histograms relation/edge.
	HistogramsTable = "histograms"
	// HistogramsInverseTable is the table name for the Histogram entity.
	// It exists in this package in order to avoid circular dependency with the "histogram" package.
	HistogramsInverseTable = "histograms"
	// HistogramsColumn is the table column denoting the histograms relation/edge.
	HistogramsColumn = "metric_histograms"
	// CountersTable is the table the holds the counters relation/edge.
	CountersTable = "counters"
	// CountersInverseTable is the table name for the Counter entity.
	// It exists in this package in order to avoid circular dependency with the "counter" package.
	CountersInverseTable = "counters"
	// CountersColumn is the table column denoting the counters relation/edge.
	CountersColumn = "metric_counters"
	// GaugesTable is the table the holds the gauges relation/edge.
	GaugesTable = "gauges"
	// GaugesInverseTable is the table name for the Gauge entity.
	// It exists in this package in order to avoid circular dependency with the "gauge" package.
	GaugesInverseTable = "gauges"
	// GaugesColumn is the table column denoting the gauges relation/edge.
	GaugesColumn = "metric_gauges"
)

// Columns holds all SQL columns for metric fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Metric type.
var ForeignKeys = []string{
	"graph_metrics",
}