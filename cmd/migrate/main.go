package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jckonewalik/migration-demo-yt/database"
	_ "github.com/lib/pq"
)

func main() {

	postgresDb, err := database.NewPosgresDB()

	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(postgresDb.Db, &postgres.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres", driver)

	if err != nil {
		log.Fatal(err)
	}

	switch cmd := os.Args[len(os.Args)-1]; cmd {
	case "up":
		m.Up()
	case "down":
		m.Down()
	default:
		log.Fatalf("invalid option: %s", cmd)
	}
}
