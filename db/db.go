package db

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	dsn := "postgresql://postgres.ixojfdzeruupylbmocoy:1Dreamgr34.@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	runMigration(db)
	return db
}

func runMigration(db *sql.DB) {
	file, err := os.ReadFile("ddl.sql")
	if err != nil {
		log.Fatal("Failed to open ddl.sql", err)
	}

	statements := strings.Split(string(file), ";")
	for _, statement := range statements {
		trimmed := strings.TrimSpace(statement)
		if trimmed != "" {
			_, err := db.Exec(trimmed)
			if err != nil {
				log.Fatal("Failed to execute migration: ", err)
			}
		}
	}
	log.Println("Database migration completed successfully")
}
