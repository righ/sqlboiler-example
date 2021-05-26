package main

import (
	"context"
	"database/sql"
	"fmt"

	models "sqlboiler-example/models_mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func main() {
	ctx := context.Background()
	db, _ := sql.Open("mysql", "usr:pw@tcp(mysql)/db")

	jets, err := models.Jets(
		qm.Select("pilots.name as name", "MIN(jets.age) AS age"),
		qm.InnerJoin("pilots on pilots.id = jets.pilot_id"),
		qm.Where("age >= 20"),
		qm.GroupBy("pilots.name"),
		qm.Having("count(*) >= 2"),
		qm.OrderBy("age DESC"),
		qm.Limit(2),
		qm.Offset(1),
	).All(ctx, db)
	if err != nil {
		fmt.Println(err)
	}
	for _, jet := range jets {
		fmt.Println(jet.Name, jet.Age)
	}
}
