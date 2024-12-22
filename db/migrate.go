package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rishu/microservice/config"
)

func main() {
	// Load configuration
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Construct PostgreSQL connection string
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		conf.PostgresConfig.User,
		conf.PostgresConfig.Password,
		conf.PostgresConfig.Host,
		conf.PostgresConfig.Port,
		conf.PostgresConfig.DBName,
		"disable",
	)

	// Command-line argument (up or down)
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run migrate.go <up|down>")
	}
	action := os.Args[1]

	fmt.Println(os.Args)

	// Initialize migration
	migrationsDir := fmt.Sprintf("file://./db/%v/migrations", os.Args[2])
	m, err := migrate.New(migrationsDir, dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	switch action {
	case "up":
		err = m.Up()
	case "down":
		err = m.Steps(-1)
	case "snapshot":
		err = snapshotDB(dbURL, conf.PostgresConfig.DBName)
	default:
		log.Fatalf("Invalid action: %s. Use 'up' or 'down'.", action)
	}

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration successful!")
}

// snapshotDB takes a snapshot of the current database schema and writes it to a file.
func snapshotDB(dbURL, dbName string) error {
	// Open connection to the database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Query to get the list of tables in the database (considering public schema and others)
	rows, err := db.Query(`
		SELECT table_schema || '.' || table_name
		FROM information_schema.tables
		WHERE table_schema NOT IN ('information_schema', 'pg_catalog')
		ORDER BY table_schema, table_name
	`)
	if err != nil {
		return fmt.Errorf("failed to fetch table list: %v", err)
	}
	defer rows.Close()

	// Start building the schema snapshot
	var schemaBuilder strings.Builder
	schemaBuilder.WriteString(fmt.Sprintf("-- Snapshot for database: %s\n", dbName))

	// Loop through the tables and fetch their DDL (CREATE TABLE statements)
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return fmt.Errorf("failed to scan table name: %v", err)
		}

		// Get the DDL for each table
		var ddl string
		err := db.QueryRow(fmt.Sprintf(`
			SELECT 'CREATE TABLE ' || table_name || E'\n(\n' ||
			string_agg(column_name || ' ' || data_type, E',\n') || E'\n);\n'
			FROM information_schema.columns
			WHERE table_schema = '%s' AND table_name = '%s'
			GROUP BY table_name`,
			strings.Split(table, ".")[0], strings.Split(table, ".")[1])).Scan(&ddl)
		if err != nil {
			// If there's an error (table doesn't exist or some other issue), log it and continue
			log.Printf("Skipping table %s due to error: %v\n", table, err)
			continue
		}

		// Append the DDL to the schema builder
		schemaBuilder.WriteString(ddl + "\n\n")
	}

	// Write the schema snapshot to a file
	snapshotFile := fmt.Sprintf("db/%s/latest.sql", dbName)
	err = os.WriteFile(snapshotFile, []byte(schemaBuilder.String()), 0644)
	if err != nil {
		return fmt.Errorf("failed to write schema to file: %v", err)
	}

	log.Printf("Snapshot saved to %s\n", snapshotFile)
	return nil
}
