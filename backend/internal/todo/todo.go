package todo

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func Migrate(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS todos (
id INTEGER PRIMARY KEY AUTOINCREMENT,
title TEXT NOT NULL,
completed INTEGER NOT NULL DEFAULT 0,
created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);`)
	return err
}

func List(db *sql.DB) ([]Todo, error) {
	rows, err := db.Query("SELECT id, title, completed, created_at FROM todos ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []Todo
	for rows.Next() {
		var t Todo
		var doneInt int
		if err := rows.Scan(&t.ID, &t.Title, &doneInt, &t.CreatedAt); err != nil {
			return nil, err
		}
		t.Completed = doneInt == 1
		res = append(res, t)
	}
	return res, nil
}

func Create(db *sql.DB, title string) (int64, error) {
	res, err := db.Exec("INSERT INTO todos (title, completed) VALUES (?, 0)", title)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func Toggle(db *sql.DB, id int64, completed bool) error {
	_, err := db.Exec("UPDATE todos SET completed = ? WHERE id = ?", boolToInt(completed), id)
	return err
}

func Delete(db *sql.DB, id int64) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
