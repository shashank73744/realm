// Code generated by ent, DO NOT EDIT.

package beacon

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the beacon type in the database.
	Label = "beacon"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPrincipal holds the string denoting the principal field in the database.
	FieldPrincipal = "principal"
	// FieldIdentifier holds the string denoting the identifier field in the database.
	FieldIdentifier = "identifier"
	// FieldAgentIdentifier holds the string denoting the agent_identifier field in the database.
	FieldAgentIdentifier = "agent_identifier"
	// FieldLastSeenAt holds the string denoting the last_seen_at field in the database.
	FieldLastSeenAt = "last_seen_at"
	// EdgeHost holds the string denoting the host edge name in mutations.
	EdgeHost = "host"
	// EdgeTasks holds the string denoting the tasks edge name in mutations.
	EdgeTasks = "tasks"
	// Table holds the table name of the beacon in the database.
	Table = "beacons"
	// HostTable is the table that holds the host relation/edge.
	HostTable = "beacons"
	// HostInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostInverseTable = "hosts"
	// HostColumn is the table column denoting the host relation/edge.
	HostColumn = "beacon_host"
	// TasksTable is the table that holds the tasks relation/edge.
	TasksTable = "tasks"
	// TasksInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TasksInverseTable = "tasks"
	// TasksColumn is the table column denoting the tasks relation/edge.
	TasksColumn = "task_beacon"
)

// Columns holds all SQL columns for beacon fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPrincipal,
	FieldIdentifier,
	FieldAgentIdentifier,
	FieldLastSeenAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "beacons"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"beacon_host",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName func() string
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// PrincipalValidator is a validator for the "principal" field. It is called by the builders before save.
	PrincipalValidator func(string) error
	// DefaultIdentifier holds the default value on creation for the "identifier" field.
	DefaultIdentifier func() string
	// IdentifierValidator is a validator for the "identifier" field. It is called by the builders before save.
	IdentifierValidator func(string) error
	// AgentIdentifierValidator is a validator for the "agent_identifier" field. It is called by the builders before save.
	AgentIdentifierValidator func(string) error
)

// OrderOption defines the ordering options for the Beacon queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPrincipal orders the results by the principal field.
func ByPrincipal(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrincipal, opts...).ToFunc()
}

// ByIdentifier orders the results by the identifier field.
func ByIdentifier(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIdentifier, opts...).ToFunc()
}

// ByAgentIdentifier orders the results by the agent_identifier field.
func ByAgentIdentifier(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAgentIdentifier, opts...).ToFunc()
}

// ByLastSeenAt orders the results by the last_seen_at field.
func ByLastSeenAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastSeenAt, opts...).ToFunc()
}

// ByHostField orders the results by host field.
func ByHostField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHostStep(), sql.OrderByField(field, opts...))
	}
}

// ByTasksCount orders the results by tasks count.
func ByTasksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTasksStep(), opts...)
	}
}

// ByTasks orders the results by tasks terms.
func ByTasks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTasksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newHostStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HostInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, HostTable, HostColumn),
	)
}
func newTasksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TasksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, TasksTable, TasksColumn),
	)
}
