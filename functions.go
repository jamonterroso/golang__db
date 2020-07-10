package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func obtenerResultadosdb() (*ResponseDatabase, error) {

	var resp ResponseDatabase
	cnn, err := sql.Open("mysql", "docker:docker@tcp(db:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}

	sqlQuery := `
		SELECT  
			id, 
			name
		FROM test_tb 
	`

	row, err := cnn.Query(
		sqlQuery,
	)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var c Database
		err := row.Scan(
			&c.ID,
			&c.Name,
		)
		if err != nil {
			return nil, err
		}
		resp.Database = append(resp.Database, c)
	}

	return &resp, nil

}
