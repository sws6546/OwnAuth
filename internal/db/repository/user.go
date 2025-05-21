package repository

import (
	"github.com/sws6546/OwnAuth/internal/db"
)

type User struct{}

func (u User) InitTable() error {
	// TODO : model
	db.DbConn.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT NOT NULL PRIMARY KEY,
		);
	`)
}
