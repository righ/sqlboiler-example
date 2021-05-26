package main

import (
	"context"
	"database/sql"
	"fmt"

	models "sqlboiler-example/models_mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	ctx := context.Background()
	db, _ := sql.Open("mysql", "usr:pw@tcp(mysql)/db")
	pilot := models.Pilot{
		ID:   1,
		Name: "Righ",
	}
	if err := pilot.Upsert(ctx, db, boil.Infer(), boil.Infer()) {
		fmt.Println(err)
	}

	jet1 := models.Jet{
		ID:      1,
		PilotID: pilot.ID,
		Age:     10,
		Name:    "F-15",
		Color:   "Gray",
	}
	if err := jet1.Upsert(ctx, db, boil.Infer(), boil.Infer()) {
		fmt.Println(err)
	}
	
	jet1.Age = 0
	jet1.Name = "MRJ"
	jet1.Color = "White"
	if err := jet1.Upsert(ctx, db, boil.Whitelist("name", "color"), boil.Infer()); err != nil {
		fmt.Println(err)
	}

	jet2 := models.Jet{
		ID:      2,
		PilotID: pilot.ID,
		Age:     20,
		Name:    "F-14",
		Color:   "Gray",
	}
	if err := jet2.Upsert(ctx, db, boil.Whitelist("name", "color"), boil.Infer()); err != nil {
		fmt.Println(err)
	}
}
