package storage

import "database/sql"

type UserStore struct {
	db *sql.DB
}

func NewUserStore(dbPath string) (*UserStore, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	store := &UserStore{db: db}
	if err := store.initialize(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *UserStore) initialize() error {
	_, err := s.db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id TEXT PRIMARY KEY,
            username TEXT UNIQUE NOT NULL,
            password_hash TEXT NOT NULL,
            created_at DATETIME NOT NULL
        )
    `)
	return err
}
