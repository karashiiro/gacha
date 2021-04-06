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
		{Name: "rate", Type: field.TypeFloat32},
		{Name: "series", Type: field.TypeString},
	}
	// DropsTable holds the schema information for the "drops" table.
	DropsTable = &schema.Table{
		Name:        "drops",
		Columns:     DropsColumns,
		PrimaryKey:  []*schema.Column{DropsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DropsTable,
	}
)

func init() {
}
