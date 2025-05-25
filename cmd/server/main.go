package main

import (
	"database/sql"
	"github.com/kaerubo/kaeruashi/internal/router"
	"github.com/kaerubo/kaeruashi/internal/wire"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://kaeruashi:pass@localhost:5432/kaeruashi-dev?sslmode=disable"
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close()

	h := wire.InitializeHandler(db)

	e := echo.New()

	router.RegisterRoutes(e, h)

	e.Logger.Fatal(e.Start(":8080"))
}
