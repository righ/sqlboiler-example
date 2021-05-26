package main

import (
	"context"
	"database/sql"
	"fmt"

	models "sqlboiler-example/models_mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func main() {
	ctx := context.Background()
	db, _ := sql.Open("mysql", "usr:pw@tcp(mysql)/db")

	pilots, err := models.Pilots(
		qm.Select("languages.language AS name"),
		qm.InnerJoin("languages ON languages.id = pilots.id"),
	).All(ctx, db)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("num of pilots:", len(pilots))
	for _, pilot := range pilots {
		fmt.Println(pilot.Name)
	}
}
