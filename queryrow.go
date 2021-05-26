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

	row := models.Pilots(
		qm.Select("id", "name"),
	).Query.QueryRow(db)

	var id int
	var name string
	if err := row.Scan(&id, &name); err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(id, name)
}
