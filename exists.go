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

	exists1, _ := models.PilotExists(ctx, db, 1)
	exists2, _ := models.PilotExists(ctx, db, 2)

	fmt.Println(exists1, exists2)
}
