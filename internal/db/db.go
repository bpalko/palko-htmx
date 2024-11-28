package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Initialize initializes the database connection and creates tables if they don't exist.
func Initialize() {
	var err error
	DB, err = sql.Open("sqlite3", "./gunbuild.db")
	if err != nil {
		log.Fatal(err)
	}

}

// SaveBuild saves a new build to the database.
func SaveBuild(build Build) error {
	query := `INSERT INTO builds (barrel, grip, sight) VALUES (?, ?, ?)`
	_, err := DB.Exec(query, build.Barrel, build.Grip, build.Sight)
	return err
}
