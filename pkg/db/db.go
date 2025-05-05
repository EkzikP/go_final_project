package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

type TasksStore struct {
	db *sql.DB
}

func New(db *sql.DB) *TasksStore {
	return &TasksStore{db: db}
}

func (s *TasksStore) Initialize() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS scheduler (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			date CHAR(8) NOT NULL DEFAULT '',
			title VARCHAR NOT NULL DEFAULT '',
			comment TEXT NOT NULL DEFAULT '',
			repeat VARCHAR NOT NULL DEFAULT ''
		);
		CREATE INDEX IF NOT EXISTS scheduler_date ON scheduler(date);
	`)
	return err
}
