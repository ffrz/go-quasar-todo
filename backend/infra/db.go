package infra

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

func NewSQLite(path string) (*sql.DB, error) {
	// pastikan folder ada
	dir := "./" + path
	if fi := os.DirFS(dir); fi == nil {
		os.MkdirAll("data", os.ModePerm)
	}

	dsn := fmt.Sprintf("file:%s?_foreign_keys=1", path)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	// migration sederhana
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
