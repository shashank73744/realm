// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcarretto/realm/tavern/internal/ent/beacon"
	"github.com/kcarretto/realm/tavern/internal/ent/host"
	"github.com/kcarretto/realm/tavern/internal/ent/predicate"
	"github.com/kcarretto/realm/tavern/internal/ent/task"
)

// BeaconUpdate is the builder for updating Beacon entities.
type BeaconUpdate struct {
	config
	hooks    []Hook
	mutation *BeaconMutation
}

// Where appends a list predicates to the BeaconUpdate builder.
func (bu *BeaconUpdate) Where(ps ...predicate.Beacon) *BeaconUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetPrincipal sets the "principal" field.
func (bu *BeaconUpdate) SetPrincipal(s string) *BeaconUpdate {
	bu.mutation.SetPrincipal(s)
	return bu
}

// SetNillablePrincipal sets the "principal" field if the given value is not nil.
func (bu *BeaconUpdate) SetNillablePrincipal(s *string) *BeaconUpdate {
	if s != nil {
		bu.SetPrincipal(*s)
	}
	return bu
}

// ClearPrincipal clears the value of the "principal" field.
func (bu *BeaconUpdate) ClearPrincipal() *BeaconUpdate {
	bu.mutation.ClearPrincipal()
	return bu
}

// SetIdentifier sets the "identifier" field.
func (bu *BeaconUpdate) SetIdentifier(s string) *BeaconUpdate {
	bu.mutation.SetIdentifier(s)
	return bu
}

// SetNillableIdentifier sets the "identifier" field if the given value is not nil.
func (bu *BeaconUpdate) SetNillableIdentifier(s *string) *BeaconUpdate {
	if s != nil {
		bu.SetIdentifier(*s)
	}
	return bu
}

// SetAgentIdentifier sets the "agent_identifier" field.
func (bu *BeaconUpdate) SetAgentIdentifier(s string) *BeaconUpdate {
	bu.mutation.SetAgentIdentifier(s)
	return bu
}

// SetNillableAgentIdentifier sets the "agent_identifier" field if the given value is not nil.
func (bu *BeaconUpdate) SetNillableAgentIdentifier(s *string) *BeaconUpdate {
	if s != nil {
		bu.SetAgentIdentifier(*s)
	}
	return bu
}

// ClearAgentIdentifier clears the value of the "agent_identifier" field.
func (bu *BeaconUpdate) ClearAgentIdentifier() *BeaconUpdate {
	bu.mutation.ClearAgentIdentifier()
	return bu
}

// SetLastSeenAt sets the "last_seen_at" field.
func (bu *BeaconUpdate) SetLastSeenAt(t time.Time) *BeaconUpdate {
	bu.mutation.SetLastSeenAt(t)
	return bu
}

// SetNillableLastSeenAt sets the "last_seen_at" field if the given value is not nil.
func (bu *BeaconUpdate) SetNillableLastSeenAt(t *time.Time) *BeaconUpdate {
	if t != nil {
		bu.SetLastSeenAt(*t)
	}
	return bu
}

// ClearLastSeenAt clears the value of the "last_seen_at" field.
func (bu *BeaconUpdate) ClearLastSeenAt() *BeaconUpdate {
	bu.mutation.ClearLastSeenAt()
	return bu
}

// SetHostID sets the "host" edge to the Host entity by ID.
func (bu *BeaconUpdate) SetHostID(id int) *BeaconUpdate {
	bu.mutation.SetHostID(id)
	return bu
}

// SetHost sets the "host" edge to the Host entity.
func (bu *BeaconUpdate) SetHost(h *Host) *BeaconUpdate {
	return bu.SetHostID(h.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (bu *BeaconUpdate) AddTaskIDs(ids ...int) *BeaconUpdate {
	bu.mutation.AddTaskIDs(ids...)
	return bu
}

// AddTasks adds the "tasks" edges to the Task entity.
func (bu *BeaconUpdate) AddTasks(t ...*Task) *BeaconUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.AddTaskIDs(ids...)
}

// Mutation returns the BeaconMutation object of the builder.
func (bu *BeaconUpdate) Mutation() *BeaconMutation {
	return bu.mutation
}

// ClearHost clears the "host" edge to the Host entity.
func (bu *BeaconUpdate) ClearHost() *BeaconUpdate {
	bu.mutation.ClearHost()
	return bu
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (bu *BeaconUpdate) ClearTasks() *BeaconUpdate {
	bu.mutation.ClearTasks()
	return bu
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (bu *BeaconUpdate) RemoveTaskIDs(ids ...int) *BeaconUpdate {
	bu.mutation.RemoveTaskIDs(ids...)
	return bu
}

// RemoveTasks removes "tasks" edges to Task entities.
func (bu *BeaconUpdate) RemoveTasks(t ...*Task) *BeaconUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.RemoveTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BeaconUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BeaconUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BeaconUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BeaconUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BeaconUpdate) check() error {
	if v, ok := bu.mutation.Principal(); ok {
		if err := beacon.PrincipalValidator(v); err != nil {
			return &ValidationError{Name: "principal", err: fmt.Errorf(`ent: validator failed for field "Beacon.principal": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Identifier(); ok {
		if err := beacon.IdentifierValidator(v); err != nil {
			return &ValidationError{Name: "identifier", err: fmt.Errorf(`ent: validator failed for field "Beacon.identifier": %w`, err)}
		}
	}
	if v, ok := bu.mutation.AgentIdentifier(); ok {
		if err := beacon.AgentIdentifierValidator(v); err != nil {
			return &ValidationError{Name: "agent_identifier", err: fmt.Errorf(`ent: validator failed for field "Beacon.agent_identifier": %w`, err)}
		}
	}
	if _, ok := bu.mutation.HostID(); bu.mutation.HostCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Beacon.host"`)
	}
	return nil
}

func (bu *BeaconUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(beacon.Table, beacon.Columns, sqlgraph.NewFieldSpec(beacon.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Principal(); ok {
		_spec.SetField(beacon.FieldPrincipal, field.TypeString, value)
	}
	if bu.mutation.PrincipalCleared() {
		_spec.ClearField(beacon.FieldPrincipal, field.TypeString)
	}
	if value, ok := bu.mutation.Identifier(); ok {
		_spec.SetField(beacon.FieldIdentifier, field.TypeString, value)
	}
	if value, ok := bu.mutation.AgentIdentifier(); ok {
		_spec.SetField(beacon.FieldAgentIdentifier, field.TypeString, value)
	}
	if bu.mutation.AgentIdentifierCleared() {
		_spec.ClearField(beacon.FieldAgentIdentifier, field.TypeString)
	}
	if value, ok := bu.mutation.LastSeenAt(); ok {
		_spec.SetField(beacon.FieldLastSeenAt, field.TypeTime, value)
	}
	if bu.mutation.LastSeenAtCleared() {
		_spec.ClearField(beacon.FieldLastSeenAt, field.TypeTime)
	}
	if bu.mutation.HostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   beacon.HostTable,
			Columns: []string{beacon.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.HostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   beacon.HostTable,
			Columns: []string{beacon.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   beacon.TasksTable,
			Columns: []string{beacon.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedTasksIDs(); len(nodes) > 0 && !bu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   beacon.TasksTable,
			Columns: []string{beacon.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   beacon.TasksTable,
			Columns: []string{beacon.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{beacon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BeaconUpdateOne is the builder for updating a single Beacon entity.
type BeaconUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BeaconMutation
}

// SetPrincipal sets the "principal" field.
func (buo *BeaconUpdateOne) SetPrincipal(s string) *BeaconUpdateOne {
	buo.mutation.SetPrincipal(s)
	return buo
}

// SetNillablePrincipal sets the "principal" field if the given value is not nil.
func (buo *BeaconUpdateOne) SetNillablePrincipal(s *string) *BeaconUpdateOne {
	if s != nil {
		buo.SetPrincipal(*s)
	}
	return buo
}

// ClearPrincipal clears the value of the "principal" field.
func (buo *BeaconUpdateOne) ClearPrincipal() *BeaconUpdateOne {
	buo.mutation.ClearPrincipal()
	return buo
}

// SetIdentifier sets the "identifier" field.
func (buo *BeaconUpdateOne) SetIdentifier(s string) *BeaconUpdateOne {
	buo.mutation.SetIdentifier(s)
	return buo
}

// SetNillableIdentifier sets the "identifier" field if the given value is not nil.
func (buo *BeaconUpdateOne) SetNillableIdentifier(s *string) *BeaconUpdateOne {
	if s != nil {
		buo.SetIdentifier(*s)
	}
	return buo
}

// SetAgentIdentifier sets the "agent_identifier" field.
func (buo *BeaconUpdateOne) SetAgentIdentifier(s string) *BeaconUpdateOne {
	buo.mutation.SetAgentIdentifier(s)
	return buo
}

// SetNillableAgentIdentifier sets the "agent_identifier" field if the given value is not nil.
func (buo *BeaconUpdateOne) SetNillableAgentIdentifier(s *string) *BeaconUpdateOne {
	if s != nil {
		buo.SetAgentIdentifier(*s)
	}
	return buo
}

// ClearAgentIdentifier clears the value of the "agent_identifier" field.
func (buo *BeaconUpdateOne) ClearAgentIdentifier() *BeaconUpdateOne {
	buo.mutation.ClearAgentIdentifier()
	return buo
}

// SetLastSeenAt sets the "last_seen_at" field.
func (buo *BeaconUpdateOne) SetLastSeenAt(t time.Time) *BeaconUpdateOne {
	buo.mutation.SetLastSeenAt(t)
	return buo
}

// SetNillableLastSeenAt sets the "last_seen_at" field if the given value is not nil.
func (buo *BeaconUpdateOne) SetNillableLastSeenAt(t *time.Time) *BeaconUpdateOne {
	if t != nil {
		buo.SetLastSeenAt(*t)
	}
	return buo
}

// ClearLastSeenAt clears the value of the "last_seen_at" field.
func (buo *BeaconUpdateOne) ClearLastSeenAt() *BeaconUpdateOne {
	buo.mutation.ClearLastSeenAt()
	return buo
}

// SetHostID sets the "host" edge to the Host entity by ID.
func (buo *BeaconUpdateOne) SetHostID(id int) *BeaconUpdateOne {
	buo.mutation.SetHostID(id)
	return buo
}

// SetHost sets the "host" edge to the Host entity.
func (buo *BeaconUpdateOne) SetHost(h *Host) *BeaconUpdateOne {
	return buo.SetHostID(h.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (buo *BeaconUpdateOne) AddTaskIDs(ids ...int) *BeaconUpdateOne {
	buo.mutation.AddTaskIDs(ids...)
	return buo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (buo *BeaconUpdateOne) AddTasks(t ...*Task) *BeaconUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.AddTaskIDs(ids...)
}

// Mutation returns the BeaconMutation object of the builder.
func (buo *BeaconUpdateOne) Mutation() *BeaconMutation {
	return buo.mutation
}

// ClearHost clears the "host" edge to the Host entity.
func (buo *BeaconUpdateOne) ClearHost() *BeaconUpdateOne {
	buo.mutation.ClearHost()
	return buo
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (buo *BeaconUpdateOne) ClearTasks() *BeaconUpdateOne {
	buo.mutation.ClearTasks()
	return buo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (buo *BeaconUpdateOne) RemoveTaskIDs(ids ...int) *BeaconUpdateOne {
	buo.mutation.RemoveTaskIDs(ids...)
	return buo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (buo *BeaconUpdateOne) RemoveTasks(t ...*Task) *BeaconUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.RemoveTaskIDs(ids...)
}

// Where appends a list predicates to the BeaconUpdate builder.
func (buo *BeaconUpdateOne) Where(ps ...predicate.Beacon) *BeaconUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BeaconUpdateOne) Select(field string, fields ...string) *BeaconUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Beacon entity.
func (buo *BeaconUpdateOne) Save(ctx context.Context) (*Beacon, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BeaconUpdateOne) SaveX(ctx context.Context) *Beacon {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BeaconUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BeaconUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BeaconUpdateOne) check() error {
	if v, ok := buo.mutation.Principal(); ok {
		if err := beacon.PrincipalValidator(v); err != nil {
			return &ValidationError{Name: "principal", err: fmt.Errorf(`ent: validator failed for field "Beacon.principal": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Identifier(); ok {
		if err := beacon.IdentifierValidator(v); err != nil {
			return &ValidationError{Name: "identifier", err: fmt.Errorf(`ent: validator failed for field "Beacon.identifier": %w`, err)}
		}
	}
	if v, ok := buo.mutation.AgentIdentifier(); ok {
		if err := beacon.AgentIdentifierValidator(v); err != nil {
			return &ValidationError{Name: "agent_identifier", err: fmt.Errorf(`ent: validator failed for field "Beacon.agent_identifier": %w`, err)}
		}
	}
	if _, ok := buo.mutation.HostID(); buo.mutation.HostCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Beacon.host"`)
	}
	return nil
}

func (buo *BeaconUpdateOne) sqlSave(ctx context.Context) (_node *Beacon, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(beacon.Table, beacon.Columns, sqlgraph.NewFieldSpec(beacon.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Beacon.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, beacon.FieldID)
		for _, f := range fields {
			if !beacon.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != beacon.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Principal(); ok {
		_spec.SetField(beacon.FieldPrincipal, field.TypeString, value)
	}
	if buo.mutation.PrincipalCleared() {
		_spec.ClearField(beacon.FieldPrincipal, field.TypeString)
	}
	if value, ok := buo.mutation.Identifier(); ok {
		_spec.SetField(beacon.FieldIdentifier, field.TypeString, value)
	}
	if value, ok := buo.mutation.AgentIdentifier(); ok {
		_spec.SetField(beacon.FieldAgentIdentifier, field.TypeString, value)
	}
	if buo.mutation.AgentIdentifierCleared() {
		_spec.ClearField(beacon.FieldAgentIdentifier, field.TypeString)
	}
	if value, ok := buo.mutation.LastSeenAt(); ok {
		_spec.SetField(beacon.FieldLastSeenAt, field.TypeTime, value)
	}
	if buo.mutation.LastSeenAtCleared() {
		_spec.ClearField(beacon.FieldLastSeenAt, field.TypeTime)
	}
	if buo.mutation.HostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   beacon.HostTable,
			Columns: []string{beacon.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.HostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   beacon.HostTable,
			Columns: []string{beacon.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   beacon.TasksTable,
			Columns: []string{beacon.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !buo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   beacon.TasksTable,
			Columns: []string{beacon.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   beacon.TasksTable,
			Columns: []string{beacon.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Beacon{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{beacon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
