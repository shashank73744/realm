// Code generated by entc, DO NOT EDIT.

package implant

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kcarretto/realm/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SessionID applies equality check predicate on the "sessionID" field. It's identical to SessionIDEQ.
func SessionID(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionID), v))
	})
}

// ProcessName applies equality check predicate on the "processName" field. It's identical to ProcessNameEQ.
func ProcessName(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcessName), v))
	})
}

// SessionIDEQ applies the EQ predicate on the "sessionID" field.
func SessionIDEQ(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionID), v))
	})
}

// SessionIDNEQ applies the NEQ predicate on the "sessionID" field.
func SessionIDNEQ(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSessionID), v))
	})
}

// SessionIDIn applies the In predicate on the "sessionID" field.
func SessionIDIn(vs ...string) predicate.Implant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSessionID), v...))
	})
}

// SessionIDNotIn applies the NotIn predicate on the "sessionID" field.
func SessionIDNotIn(vs ...string) predicate.Implant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSessionID), v...))
	})
}

// SessionIDGT applies the GT predicate on the "sessionID" field.
func SessionIDGT(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSessionID), v))
	})
}

// SessionIDGTE applies the GTE predicate on the "sessionID" field.
func SessionIDGTE(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSessionID), v))
	})
}

// SessionIDLT applies the LT predicate on the "sessionID" field.
func SessionIDLT(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSessionID), v))
	})
}

// SessionIDLTE applies the LTE predicate on the "sessionID" field.
func SessionIDLTE(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSessionID), v))
	})
}

// SessionIDContains applies the Contains predicate on the "sessionID" field.
func SessionIDContains(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSessionID), v))
	})
}

// SessionIDHasPrefix applies the HasPrefix predicate on the "sessionID" field.
func SessionIDHasPrefix(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSessionID), v))
	})
}

// SessionIDHasSuffix applies the HasSuffix predicate on the "sessionID" field.
func SessionIDHasSuffix(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSessionID), v))
	})
}

// SessionIDEqualFold applies the EqualFold predicate on the "sessionID" field.
func SessionIDEqualFold(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSessionID), v))
	})
}

// SessionIDContainsFold applies the ContainsFold predicate on the "sessionID" field.
func SessionIDContainsFold(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSessionID), v))
	})
}

// ProcessNameEQ applies the EQ predicate on the "processName" field.
func ProcessNameEQ(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcessName), v))
	})
}

// ProcessNameNEQ applies the NEQ predicate on the "processName" field.
func ProcessNameNEQ(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProcessName), v))
	})
}

// ProcessNameIn applies the In predicate on the "processName" field.
func ProcessNameIn(vs ...string) predicate.Implant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProcessName), v...))
	})
}

// ProcessNameNotIn applies the NotIn predicate on the "processName" field.
func ProcessNameNotIn(vs ...string) predicate.Implant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProcessName), v...))
	})
}

// ProcessNameGT applies the GT predicate on the "processName" field.
func ProcessNameGT(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProcessName), v))
	})
}

// ProcessNameGTE applies the GTE predicate on the "processName" field.
func ProcessNameGTE(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProcessName), v))
	})
}

// ProcessNameLT applies the LT predicate on the "processName" field.
func ProcessNameLT(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProcessName), v))
	})
}

// ProcessNameLTE applies the LTE predicate on the "processName" field.
func ProcessNameLTE(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProcessName), v))
	})
}

// ProcessNameContains applies the Contains predicate on the "processName" field.
func ProcessNameContains(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProcessName), v))
	})
}

// ProcessNameHasPrefix applies the HasPrefix predicate on the "processName" field.
func ProcessNameHasPrefix(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProcessName), v))
	})
}

// ProcessNameHasSuffix applies the HasSuffix predicate on the "processName" field.
func ProcessNameHasSuffix(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProcessName), v))
	})
}

// ProcessNameIsNil applies the IsNil predicate on the "processName" field.
func ProcessNameIsNil() predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldProcessName)))
	})
}

// ProcessNameNotNil applies the NotNil predicate on the "processName" field.
func ProcessNameNotNil() predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldProcessName)))
	})
}

// ProcessNameEqualFold applies the EqualFold predicate on the "processName" field.
func ProcessNameEqualFold(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProcessName), v))
	})
}

// ProcessNameContainsFold applies the ContainsFold predicate on the "processName" field.
func ProcessNameContainsFold(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProcessName), v))
	})
}

// HasTarget applies the HasEdge predicate on the "target" edge.
func HasTarget() predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TargetTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TargetTable, TargetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTargetWith applies the HasEdge predicate on the "target" edge with a given conditions (other predicates).
func HasTargetWith(preds ...predicate.Target) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TargetInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TargetTable, TargetColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasConfig applies the HasEdge predicate on the "config" edge.
func HasConfig() predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConfigTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ConfigTable, ConfigColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConfigWith applies the HasEdge predicate on the "config" edge with a given conditions (other predicates).
func HasConfigWith(preds ...predicate.ImplantConfig) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConfigInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ConfigTable, ConfigColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Implant) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Implant) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Implant) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		p(s.Not())
	})
}
