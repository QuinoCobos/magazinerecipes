package sql

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	l "magazinerecipes/logger"
	u "magazinerecipes/utils"
)

const (
	dateFormat = "20060102"
)

var (
	db     *sql.DB
	schema string
)

//Connect creates a mysql connection
func Connect(user, password, dataBaseName string) (*sql.DB, error) {
	schema = dataBaseName
	dataBase, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", user, password, schema))
	db = dataBase
	return dataBase, err
}

//Duplicate creates a copy of given table and insert all data
func Duplicate(tableName string) error {
	today := time.Now().Format(dateFormat)
	duplicateTable := fmt.Sprintf("CREATE TABLE %s_%s LIKE %s", tableName, today, tableName)
	l.Sql(duplicateTable)
	_, err := db.Exec(duplicateTable)
	if err != nil {
		return err
	}
	duplicateData := fmt.Sprintf("INSERT INTO %s_%s SELECT * FROM %s", tableName, today, tableName)
	l.Sql(duplicateData)
	_, err2 := db.Exec(duplicateData)
	return err2
}

//Truncate erases all data from a table
func Truncate(tableName string) (sql.Result, error) {
	query := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
	l.Sql(query)
	return db.Exec(query)
}

//ShowTables returns all tables in schema
func ShowTables() (*sql.Rows, error) {
	query := fmt.Sprintf("SHOW FULL TABLES FROM %s WHERE table_type = 'BASE TABLE'", schema)
	l.Sql(query)
	return db.Query(query)
}

func Insert(table string, fields, values []string) (sql.Result, error) {
	if len(fields) != len(values) {
		panic(errors.New("Can't insert, different number of fields and values"))
	}
	for i, value := range values {
		values[i] = u.Literal(value)
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(fields, ","), strings.Join(values, ","))
	l.Sql(query)
	return db.Exec(query)
}

func Select(table, fields, where string) (*sql.Rows, error) {
	if fields == "" {
		fields = "*"
	}
	query := fmt.Sprintf("SELECT %s FROM %s WHERE 1=1 AND %s", fields, table, where)
	l.Sql(query)
	return db.Query(query)
}

func SelectQuery(query string) (*sql.Rows, error) {
	l.Sql(query)
	return db.Query(query)
}

func MaxField(table, field, where string) *sql.Row {
	query := fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE 1=1 AND %s", field, table, where)
	l.Sql(query)
	return db.QueryRow(query)
}

func Count(table, field, where string) *sql.Row {
	query := fmt.Sprintf("SELECT COUNT(%s) FROM %s WHERE 1=1 AND %s", field, table, where)
	l.Sql(query)
	return db.QueryRow(query)
}

func CountDistinct(table, field, where string) *sql.Row {
	query := fmt.Sprintf("SELECT COUNT(DISTINCT %s) FROM %s WHERE 1=1 AND %s", field, table, where)
	l.Sql(query)
	return db.QueryRow(query)
}
