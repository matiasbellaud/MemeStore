package datamanagement

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func AddDeleteUpdateDB(query string, args ...interface{}) sql.Result {
	db, err := sql.Open("postgres", "user=bella password=171104 host=localhost dbname=test_pgadmin sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer db.Close()

	res, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return res
}

func SelectDB(query string, args ...interface{}) *sql.Rows {
	db, err := sql.Open("postgres", "user=bella password=171104 host=localhost dbname=test_pgadmin sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer db.Close()
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return rows
}
