package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sws6546/OwnAuth/config"
)

var (
	host     = config.PgHost
	port     = config.PgPort
	user     = config.PgUser
	dbname   = config.PgDbName
	password = config.PgPassword
	sslmode  = config.PgSslMode
)

var Db *sql.DB

func InitConnection() error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	Db = db

	return nil
}

func CloseDb() {
	Db.Close()
}
