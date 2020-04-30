package structs

import (
	"database/sql"

	l "magazinerecipes/logger"
	u "magazinerecipes/utils"
)

//Table struct
type Table struct {
	Name string `json:"Tables_in_new_strategy"`
	Type string `json:"Table_type"`
}

func (t *Table) String() string {
	return t.Name + ";" + t.Type
}

func GetTables(r *sql.Rows) []Table {
	tables := []Table{}
	for r.Next() {
		var t Table
		u.Check(r.Scan(&t.Name, &t.Type))
		l.Info(t.String())
		tables = append(tables, t)
	}
	return tables
}
