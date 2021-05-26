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

	exists, err := models.Pilots().Exists(ctx, db)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(exists)
}
