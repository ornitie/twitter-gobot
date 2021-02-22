package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	_ = godotenv.Load()
	db := initializeDatabase()

	defer db.Close()

	files := readMigrationFiles()
	files = filterMigrations(db, files)

	runMigrations(db, files)
}

func readMigrationFiles() []string {
	files, err := filepath.Glob("./migrations/*.sql")

	if err != nil {
		log.Fatalf("No migration files found [%v]", err)
	}

	sort.Strings(files)

	return files
}

func filterMigrations(db *sql.DB, files []string) []string {
	var pendingMigrations []string
	for _, file := range files {
		filename := filepath.Base(file)
		var id int
		row := db.QueryRow("select id from migrations where name='%1'", filename)
		_ = row.Scan(&id)
		if id == 0 {
			pendingMigrations = append(pendingMigrations, file)
		}
	}

	return pendingMigrations
}

func runMigrations(db *sql.DB, files []string) {
	for _, filename := range files {
		file, _ := ioutil.ReadFile(filename)
		_, err := db.Exec(string(file))

		if err != nil {
			log.Fatalf("Error %v running migration %s", err, filename)
		}

		query := fmt.Sprintf("insert into migrations(name) values('%s')", filepath.Base(filename))

		_, err = db.Exec(query)

		if err != nil {
			log.Fatalf("Error %v inserting migration %s", err, filepath.Base(filename))
		}
	}
}

func initializeDatabase() *sql.DB {
	host, _ := os.LookupEnv("DATABASE_HOST")
	port, _ := os.LookupEnv("DATABASE_PORT")
	user, _ := os.LookupEnv("DATABASE_USER")
	password, _ := os.LookupEnv("DATABASE_PASSWORD")
	database, _ := os.LookupEnv("DATABASE_NAME")

	databaseInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		database,
	)

	db, err := sql.Open("postgres", databaseInfo)

	if err != nil {
		panic(err)
	}

	return db
}
