package main

import (
	"context"
	"database/sql"
	"fmt"

	models "sqlboiler-example/models_mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/queries"
)

type PilotAndJet struct {
	models.Pilot    `boil:",bind"`
	models.Language `boil:",bind"`
}

func main() {
	var pajs []PilotAndJet
	db, _ := sql.Open("mysql", "usr:pw@tcp(mysql)/db")
	if err := queries.Raw(`
		SELECT
		  p.id AS "pilot.id", p.name, COALESCE(l.language, 'NO LANG') AS language
		FROM pilots p
		  LEFT JOIN pilot_languages pl ON pl.pilot_id = p.id
		  LEFT JOIN languages l ON l.id = pl.language_id AND p.id != 1`,
	).Bind(context.Background(), db, &pajs); err != nil {
		fmt.Println(err)
	}
	for _, paj := range pajs {
		fmt.Println(paj)
	}
}
