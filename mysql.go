package main

import (
	"context"
	"database/sql"
	"fmt"

	models "sqlboiler-example/models_mysql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "usr:pw@tcp(mysql)/db")
	jets, _ := models.Jets().All(context.Background(), db)
	fmt.Println("jets:", jets)
}
