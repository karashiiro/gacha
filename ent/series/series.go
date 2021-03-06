// Code generated by entc, DO NOT EDIT.

package series

const (
	// Label holds the string label denoting the series type in the database.
	Label = "series"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeDrops holds the string denoting the drops edge name in mutations.
	EdgeDrops = "drops"
	// Table holds the table name of the series in the database.
	Table = "series"
	// DropsTable is the table the holds the drops relation/edge.
	DropsTable = "drops"
	// DropsInverseTable is the table name for the Drop entity.
	// It exists in this package in order to avoid circular dependency with the "drop" package.
	DropsInverseTable = "drops"
	// DropsColumn is the table column denoting the drops relation/edge.
	DropsColumn = "series_id"
)

// Columns holds all SQL columns for series fields.
var Columns = []string{
	FieldID,
	FieldName,
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
