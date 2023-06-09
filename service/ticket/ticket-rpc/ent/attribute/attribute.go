// Code generated by ent, DO NOT EDIT.

package attribute

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the attribute type in the database.
	Label = "attribute"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldEntityID holds the string denoting the entity_id field in the database.
	FieldEntityID = "entity_id"
	// FieldAttributeCode holds the string denoting the attribute_code field in the database.
	FieldAttributeCode = "attribute_code"
	// FieldBackendClass holds the string denoting the backend_class field in the database.
	FieldBackendClass = "backend_class"
	// FieldBackendType holds the string denoting the backend_type field in the database.
	FieldBackendType = "backend_type"
	// FieldBackendTable holds the string denoting the backend_table field in the database.
	FieldBackendTable = "backend_table"
	// FieldFrontendClass holds the string denoting the frontend_class field in the database.
	FieldFrontendClass = "frontend_class"
	// FieldFrontendType holds the string denoting the frontend_type field in the database.
	FieldFrontendType = "frontend_type"
	// FieldFrontendLabel holds the string denoting the frontend_label field in the database.
	FieldFrontendLabel = "frontend_label"
	// FieldSourceClass holds the string denoting the source_class field in the database.
	FieldSourceClass = "source_class"
	// FieldDefaultValue holds the string denoting the default_value field in the database.
	FieldDefaultValue = "default_value"
	// FieldIsFilterable holds the string denoting the is_filterable field in the database.
	FieldIsFilterable = "is_filterable"
	// FieldIsSearchable holds the string denoting the is_searchable field in the database.
	FieldIsSearchable = "is_searchable"
	// FieldIsRequired holds the string denoting the is_required field in the database.
	FieldIsRequired = "is_required"
	// FieldRequiredValidateClass holds the string denoting the required_validate_class field in the database.
	FieldRequiredValidateClass = "required_validate_class"
	// EdgeOptionID holds the string denoting the option_id edge name in mutations.
	EdgeOptionID = "option_id"
	// Table holds the table name of the attribute in the database.
	Table = "attributes"
	// OptionIDTable is the table that holds the option_id relation/edge.
	OptionIDTable = "attribute_options"
	// OptionIDInverseTable is the table name for the AttributeOption entity.
	// It exists in this package in order to avoid circular dependency with the "attributeoption" package.
	OptionIDInverseTable = "attribute_options"
	// OptionIDColumn is the table column denoting the option_id relation/edge.
	OptionIDColumn = "attribute_option_id"
)

// Columns holds all SQL columns for attribute fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldEntityID,
	FieldAttributeCode,
	FieldBackendClass,
	FieldBackendType,
	FieldBackendTable,
	FieldFrontendClass,
	FieldFrontendType,
	FieldFrontendLabel,
	FieldSourceClass,
	FieldDefaultValue,
	FieldIsFilterable,
	FieldIsSearchable,
	FieldIsRequired,
	FieldRequiredValidateClass,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultIsFilterable holds the default value on creation for the "is_filterable" field.
	DefaultIsFilterable uint8
	// DefaultIsSearchable holds the default value on creation for the "is_searchable" field.
	DefaultIsSearchable uint8
	// DefaultIsRequired holds the default value on creation for the "is_required" field.
	DefaultIsRequired uint8
)

// OrderOption defines the ordering options for the Attribute queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByEntityID orders the results by the entity_id field.
func ByEntityID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntityID, opts...).ToFunc()
}

// ByAttributeCode orders the results by the attribute_code field.
func ByAttributeCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAttributeCode, opts...).ToFunc()
}

// ByBackendClass orders the results by the backend_class field.
func ByBackendClass(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBackendClass, opts...).ToFunc()
}

// ByBackendType orders the results by the backend_type field.
func ByBackendType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBackendType, opts...).ToFunc()
}

// ByBackendTable orders the results by the backend_table field.
func ByBackendTable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBackendTable, opts...).ToFunc()
}

// ByFrontendClass orders the results by the frontend_class field.
func ByFrontendClass(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrontendClass, opts...).ToFunc()
}

// ByFrontendType orders the results by the frontend_type field.
func ByFrontendType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrontendType, opts...).ToFunc()
}

// ByFrontendLabel orders the results by the frontend_label field.
func ByFrontendLabel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrontendLabel, opts...).ToFunc()
}

// BySourceClass orders the results by the source_class field.
func BySourceClass(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSourceClass, opts...).ToFunc()
}

// ByDefaultValue orders the results by the default_value field.
func ByDefaultValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefaultValue, opts...).ToFunc()
}

// ByIsFilterable orders the results by the is_filterable field.
func ByIsFilterable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsFilterable, opts...).ToFunc()
}

// ByIsSearchable orders the results by the is_searchable field.
func ByIsSearchable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsSearchable, opts...).ToFunc()
}

// ByIsRequired orders the results by the is_required field.
func ByIsRequired(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsRequired, opts...).ToFunc()
}

// ByRequiredValidateClass orders the results by the required_validate_class field.
func ByRequiredValidateClass(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequiredValidateClass, opts...).ToFunc()
}

// ByOptionIDCount orders the results by option_id count.
func ByOptionIDCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOptionIDStep(), opts...)
	}
}

// ByOptionID orders the results by option_id terms.
func ByOptionID(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOptionIDStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOptionIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OptionIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OptionIDTable, OptionIDColumn),
	)
}
