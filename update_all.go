package main

import (
	"context"
	"database/sql"

	models "sqlboiler-example/models_mysql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	db, _ := sql.Open("mysql", "usr:pw@tcp(mysql)/db")

	models.Pilots().UpdateAll(ctx, db, map[string]interface{}{
		"name": "Saki",
	})
}
