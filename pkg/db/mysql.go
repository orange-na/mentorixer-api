package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		"root",
		"Yuto8181nmb",
		"localhost",
		"3306",
		"todos",
	)
	return sql.Open("mysql", dsn)
}