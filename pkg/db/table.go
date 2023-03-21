package db

import "github.com/lib/pq"

// postgres does not support prepared statements for
// statements that modify schemas
// https://github.com/lib/pq/issues/694#issuecomment-356180769

func DropTable(name string) string {
	return "DROP TABLE IF EXISTS " + pq.QuoteIdentifier(name)
}

func CreateTable(name, schema string) string {
	return "CREATE TABLE " + pq.QuoteIdentifier(name) + "(" +
		schema +
		")"
}
