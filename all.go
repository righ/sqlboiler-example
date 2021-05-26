package main

import (
	"context"
	"database/sql"
	"fmt"

	models "sqlboiler-example/models_mysql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	db, _ := sql.Open("mysql", "usr:pw@tcp(mysql)/db")

	pilots, err := models.Pilots().All(ctx, db)
	if err != nil {
		fmt.Println("err:", err)
	}
	for _, pilot := range pilots {
		fmt.Println(pilot.Name)
	}
}
