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

	pilot1, _ := models.FindPilot(ctx, db, 1)
	pilot2, _ := models.FindPilot(ctx, db, 2)

	fmt.Println(pilot1, pilot2)
}
