package main

import (
	"context"
	"database/sql"

	models "sqlboiler-example/models_mysql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "usr:pw@tcp(mysql)/db")
	ctx := context.Background()
	pilot, _ := models.FindPilot(ctx, db, 1)
	pilot.Delete(ctx, db)
}
