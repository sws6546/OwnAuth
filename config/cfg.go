package config

import (
	"os"
	"strconv"
)

var (
	Port   = getPort()
	JwtKey = getJwtKey()

	PgHost     = getPgHost()
	PgPort     = getPgPort()
	PgUser     = getPgUser()
	PgPassword = getPgPassword()
	PgDbName   = getPgDbName()
	PgSslMode  = getPgSslMode()
)

func getPort() int {
	port := os.Getenv("PORT")
	if port == "" {
		return 8080
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		return 8080
	}
	return p
}

func getPgSslMode() string {
	sslMode := os.Getenv("PG_SSL_MODE")
	if sslMode == "" {
		return "disable"
	}
	return sslMode
}

func getPgDbName() string {
	dbName := os.Getenv("PG_DB_NAME")
	if dbName == "" {
		return "ownAuth"
	}
	return dbName
}

func getPgUser() string {
	user := os.Getenv("PG_USER")
	if user == "" {
		return "postgres"
	}
	return user
}
func getPgPassword() string {
	password := os.Getenv("PG_PASSWORD")
	if password == "" {
		return "postgres"
	}
	return password
}

func getPgPort() int {
	port := os.Getenv("PG_PORT")
	if port == "" {
		return 5432
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		return 5432
	}
	return p
}

func getPgHost() string {
	key := os.Getenv("PG_HOST")
	if key == "" {
		return "localhost"
	}
	return key
}

func getJwtKey() string {
	key := os.Getenv("JWT_KEY")
	if key == "" {
		return "superSecretKey"
	}
	return key
}
