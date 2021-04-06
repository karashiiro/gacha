// Code generated by entc, DO NOT EDIT.

package drop

const (
	// Label holds the string label denoting the drop type in the database.
	Label = "drop"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldObjectID holds the string denoting the object_id field in the database.
	FieldObjectID = "object_id"
	// FieldRate holds the string denoting the rate field in the database.
	FieldRate = "rate"
	// FieldSeriesID holds the string denoting the series_id field in the database.
	FieldSeriesID = "series_id"
	// EdgeSeries holds the string denoting the series edge name in mutations.
	EdgeSeries = "series"
	// Table holds the table name of the drop in the database.
	Table = "drops"
	// SeriesTable is the table the holds the series relation/edge.
	SeriesTable = "drops"
	// SeriesInverseTable is the table name for the Series entity.
	// It exists in this package in order to avoid circular dependency with the "series" package.
	SeriesInverseTable = "series"
	// SeriesColumn is the table column denoting the series relation/edge.
	SeriesColumn = "series_id"
)

// Columns holds all SQL columns for drop fields.
var Columns = []string{
	FieldID,
	FieldObjectID,
	FieldRate,
	FieldSeriesID,
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
