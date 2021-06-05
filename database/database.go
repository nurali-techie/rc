package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/nurali-techie/rc/config"
)

func GetDatabase() (*sql.DB, func(db *sql.DB)) {
	db := openDB()
	closeFn := func(db *sql.DB) {
		if db != nil {
			db.Close()
		}
	}

	err := setupDB(db)
	if err != nil {
		closeFn(db)
		panic(err)
	}
	return db, closeFn
}

func openDB() *sql.DB {
	dbPath := config.GetDBPath()
	if dbPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		dbPath = fmt.Sprintf("%s/.rc", homeDir)
	}

	err := os.Mkdir(dbPath, 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	dbFile := fmt.Sprintf("%s/rc.db", dbPath)
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}
	return db
}

func setupDB(db *sql.DB) error {
	err := createTable(db, "CREATE TABLE IF NOT EXISTS commands (key TEXT PRIMARY KEY, text TEXT)")
	if err != nil {
		return err
	}
	err = createTable(db, "CREATE TABLE IF NOT EXISTS memory (seq INTEGER PRIMARY KEY, key TEXT)")
	return err
}

func createTable(db *sql.DB, query string) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
