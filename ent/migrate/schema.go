// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DropsColumns holds the columns for the "drops" table.
	DropsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "object_id", Type: field.TypeUint32},
		{Name: "rate", Type: field.TypeFloat32},
		{Name: "series_id", Type: field.TypeUint32, Nullable: true},
	}
	// DropsTable holds the schema information for the "drops" table.
	DropsTable = &schema.Table{
		Name:       "drops",
		Columns:    DropsColumns,
		PrimaryKey: []*schema.Column{DropsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "drops_series_drops",
				Columns:    []*schema.Column{DropsColumns[3]},
				RefColumns: []*schema.Column{SeriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SeriesColumns holds the columns for the "series" table.
	SeriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// SeriesTable holds the schema information for the "series" table.
	SeriesTable = &schema.Table{
		Name:        "series",
		Columns:     SeriesColumns,
		PrimaryKey:  []*schema.Column{SeriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DropsTable,
		SeriesTable,
	}
)

func init() {
	DropsTable.ForeignKeys[0].RefTable = SeriesTable
}
