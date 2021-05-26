package main

import (
	"context"
	"database/sql"

	models "sqlboiler-example/models_mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	db, _ := sql.Open("mysql", "usr:pw@tcp(mysql)/db")
	pilot := models.Pilot{
		ID:   1,
		Name: "Righ",
	}
	pilot.Insert(context.Background(), db, boil.Infer())
}
