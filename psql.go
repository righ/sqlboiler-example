package main
  
import (
	"context"
	"database/sql"
	"fmt"

	models "sqlboiler-example/models_psql"

	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open(
		"postgres",
		"host=postgres dbname=db user=usr password=pw sslmode=disable")
	jets, _ := models.Jets().All(context.Background(), db)
	fmt.Println("jets:", jets)
}
