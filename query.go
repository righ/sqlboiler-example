package main

import (
	"database/sql"
	"fmt"

	models "sqlboiler-example/models_mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func main() {
	db, _ := sql.Open("mysql", "usr:pw@tcp(mysql)/db")

	rows, err := models.Pilots(
		qm.Select("id", "name"),
	).Query.Query(db)
	if err != nil {
		fmt.Println("err:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			fmt.Println("err:", err)
			continue
		}
		fmt.Println(id, name)
	}
}
