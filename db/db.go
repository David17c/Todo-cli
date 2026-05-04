package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func Init() (*sql.DB, error) {
	baseDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	appDir := filepath.Join(baseDir, "Todo-cli")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return nil, err
	}

	dbPath := filepath.Join(appDir, "tasks.db")

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	// Create tables ONCE here
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY,
		taskname TEXT NOT NULL,
		status INTEGER NOT NULL
	);
	`)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
