package sqlbuilder

import (
    "fmt"
    "strings"
)
// ColumnValue represents a single database column and its corresponding value
// to be used in SQL INSERT statements.
type ColumnValue struct {
	Column string
	Value  string
}

// BuildSelectQuery generates a valid SQL SELECT query for the specified table and columns.
// If the columns slice is empty or nil, it defaults to the wildcard "*".
func BuildSelectQuery(table string, columns []string) string {
    columnsString := "*"
    if len(columns) > 0 {
        columnsString = strings.Join(columns, ", ")
    }
	return fmt.Sprintf("SELECT %s FROM %s;", columnsString, table)
}

// BuildBatchInserts generates a slice of individual SQL INSERT statements for the provided rows.
// Empty rows are skipped, and values must be enclosed in single quotes.
func BuildBatchInserts(table string, rows [][]ColumnValue) []string {
	inserts := make([]string, 0, len(rows))
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		columns := make([]string, 0, len(row))
		values := make([]string, 0, len(row))
		for _, cv := range row {
			columns = append(columns, cv.Column)
			values = append(values, "'"+cv.Value+"'")
		}
		query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
			table,
			strings.Join(columns, ", "),
			strings.Join(values, ", "),
		)
		inserts = append(inserts, query)
	}

	return inserts
}
