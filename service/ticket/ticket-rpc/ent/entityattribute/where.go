// Code generated by ent, DO NOT EDIT.

package entityattribute

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldUpdatedAt, v))
}

// AttributeID applies equality check predicate on the "attribute_id" field. It's identical to AttributeIDEQ.
func AttributeID(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldAttributeID, v))
}

// EntityID applies equality check predicate on the "entity_id" field. It's identical to EntityIDEQ.
func EntityID(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldEntityID, v))
}

// AttributeSetID applies equality check predicate on the "attribute_set_id" field. It's identical to AttributeSetIDEQ.
func AttributeSetID(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldAttributeSetID, v))
}

// AttributeGroupID applies equality check predicate on the "attribute_group_id" field. It's identical to AttributeGroupIDEQ.
func AttributeGroupID(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldAttributeGroupID, v))
}

// Sequence applies equality check predicate on the "sequence" field. It's identical to SequenceEQ.
func Sequence(v uint8) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldSequence, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLTE(FieldUpdatedAt, v))
}

// AttributeIDEQ applies the EQ predicate on the "attribute_id" field.
func AttributeIDEQ(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldAttributeID, v))
}

// AttributeIDNEQ applies the NEQ predicate on the "attribute_id" field.
func AttributeIDNEQ(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNEQ(FieldAttributeID, v))
}

// AttributeIDIn applies the In predicate on the "attribute_id" field.
func AttributeIDIn(vs ...uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldIn(FieldAttributeID, vs...))
}

// AttributeIDNotIn applies the NotIn predicate on the "attribute_id" field.
func AttributeIDNotIn(vs ...uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNotIn(FieldAttributeID, vs...))
}

// AttributeIDGT applies the GT predicate on the "attribute_id" field.
func AttributeIDGT(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGT(FieldAttributeID, v))
}

// AttributeIDGTE applies the GTE predicate on the "attribute_id" field.
func AttributeIDGTE(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGTE(FieldAttributeID, v))
}

// AttributeIDLT applies the LT predicate on the "attribute_id" field.
func AttributeIDLT(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLT(FieldAttributeID, v))
}

// AttributeIDLTE applies the LTE predicate on the "attribute_id" field.
func AttributeIDLTE(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLTE(FieldAttributeID, v))
}

// EntityIDEQ applies the EQ predicate on the "entity_id" field.
func EntityIDEQ(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldEntityID, v))
}

// EntityIDNEQ applies the NEQ predicate on the "entity_id" field.
func EntityIDNEQ(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNEQ(FieldEntityID, v))
}

// EntityIDIn applies the In predicate on the "entity_id" field.
func EntityIDIn(vs ...uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldIn(FieldEntityID, vs...))
}

// EntityIDNotIn applies the NotIn predicate on the "entity_id" field.
func EntityIDNotIn(vs ...uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNotIn(FieldEntityID, vs...))
}

// EntityIDGT applies the GT predicate on the "entity_id" field.
func EntityIDGT(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGT(FieldEntityID, v))
}

// EntityIDGTE applies the GTE predicate on the "entity_id" field.
func EntityIDGTE(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGTE(FieldEntityID, v))
}

// EntityIDLT applies the LT predicate on the "entity_id" field.
func EntityIDLT(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLT(FieldEntityID, v))
}

// EntityIDLTE applies the LTE predicate on the "entity_id" field.
func EntityIDLTE(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLTE(FieldEntityID, v))
}

// AttributeSetIDEQ applies the EQ predicate on the "attribute_set_id" field.
func AttributeSetIDEQ(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldAttributeSetID, v))
}

// AttributeSetIDNEQ applies the NEQ predicate on the "attribute_set_id" field.
func AttributeSetIDNEQ(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNEQ(FieldAttributeSetID, v))
}

// AttributeSetIDIn applies the In predicate on the "attribute_set_id" field.
func AttributeSetIDIn(vs ...uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldIn(FieldAttributeSetID, vs...))
}

// AttributeSetIDNotIn applies the NotIn predicate on the "attribute_set_id" field.
func AttributeSetIDNotIn(vs ...uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNotIn(FieldAttributeSetID, vs...))
}

// AttributeSetIDGT applies the GT predicate on the "attribute_set_id" field.
func AttributeSetIDGT(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGT(FieldAttributeSetID, v))
}

// AttributeSetIDGTE applies the GTE predicate on the "attribute_set_id" field.
func AttributeSetIDGTE(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGTE(FieldAttributeSetID, v))
}

// AttributeSetIDLT applies the LT predicate on the "attribute_set_id" field.
func AttributeSetIDLT(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLT(FieldAttributeSetID, v))
}

// AttributeSetIDLTE applies the LTE predicate on the "attribute_set_id" field.
func AttributeSetIDLTE(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLTE(FieldAttributeSetID, v))
}

// AttributeGroupIDEQ applies the EQ predicate on the "attribute_group_id" field.
func AttributeGroupIDEQ(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldAttributeGroupID, v))
}

// AttributeGroupIDNEQ applies the NEQ predicate on the "attribute_group_id" field.
func AttributeGroupIDNEQ(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNEQ(FieldAttributeGroupID, v))
}

// AttributeGroupIDIn applies the In predicate on the "attribute_group_id" field.
func AttributeGroupIDIn(vs ...uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldIn(FieldAttributeGroupID, vs...))
}

// AttributeGroupIDNotIn applies the NotIn predicate on the "attribute_group_id" field.
func AttributeGroupIDNotIn(vs ...uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNotIn(FieldAttributeGroupID, vs...))
}

// AttributeGroupIDGT applies the GT predicate on the "attribute_group_id" field.
func AttributeGroupIDGT(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGT(FieldAttributeGroupID, v))
}

// AttributeGroupIDGTE applies the GTE predicate on the "attribute_group_id" field.
func AttributeGroupIDGTE(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGTE(FieldAttributeGroupID, v))
}

// AttributeGroupIDLT applies the LT predicate on the "attribute_group_id" field.
func AttributeGroupIDLT(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLT(FieldAttributeGroupID, v))
}

// AttributeGroupIDLTE applies the LTE predicate on the "attribute_group_id" field.
func AttributeGroupIDLTE(v uint64) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLTE(FieldAttributeGroupID, v))
}

// SequenceEQ applies the EQ predicate on the "sequence" field.
func SequenceEQ(v uint8) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldEQ(FieldSequence, v))
}

// SequenceNEQ applies the NEQ predicate on the "sequence" field.
func SequenceNEQ(v uint8) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNEQ(FieldSequence, v))
}

// SequenceIn applies the In predicate on the "sequence" field.
func SequenceIn(vs ...uint8) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldIn(FieldSequence, vs...))
}

// SequenceNotIn applies the NotIn predicate on the "sequence" field.
func SequenceNotIn(vs ...uint8) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldNotIn(FieldSequence, vs...))
}

// SequenceGT applies the GT predicate on the "sequence" field.
func SequenceGT(v uint8) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGT(FieldSequence, v))
}

// SequenceGTE applies the GTE predicate on the "sequence" field.
func SequenceGTE(v uint8) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldGTE(FieldSequence, v))
}

// SequenceLT applies the LT predicate on the "sequence" field.
func SequenceLT(v uint8) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLT(FieldSequence, v))
}

// SequenceLTE applies the LTE predicate on the "sequence" field.
func SequenceLTE(v uint8) predicate.EntityAttribute {
	return predicate.EntityAttribute(sql.FieldLTE(FieldSequence, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EntityAttribute) predicate.EntityAttribute {
	return predicate.EntityAttribute(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EntityAttribute) predicate.EntityAttribute {
	return predicate.EntityAttribute(func(s *sql.Selector) {
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
func Not(p predicate.EntityAttribute) predicate.EntityAttribute {
	return predicate.EntityAttribute(func(s *sql.Selector) {
		p(s.Not())
	})
}
