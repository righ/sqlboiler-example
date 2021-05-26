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
	pilot := models.Pilot{
		ID:   1,
		Name: "",
	}
	pilot.Reload(context.Background(), db)
	fmt.Println(pilot.ID, pilot.Name)
}
