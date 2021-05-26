package main

import (
	"context"
	"database/sql"
	"fmt"

	models "sqlboiler-example/models_mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/queries"
)

func main() {
	ctx := context.Background()
	db, _ := sql.Open("mysql", "usr:pw@tcp(mysql)/db")

	queries.Raw(`
		INSERT INTO pilots (id, name)
		VALUES (6, 'Diddy')
	`).Exec(db)

	pilot, _ := models.FindPilot(ctx, db, 6)
	fmt.Println(pilot)
}
