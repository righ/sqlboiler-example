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

	count, err := models.Pilots().Count(ctx, db)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("num of pilots:", count)
}
