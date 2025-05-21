package main

import (
	"github.com/sws6546/OwnAuth/internal/api"
	"github.com/sws6546/OwnAuth/internal/db"
)

func main() {
	db.InitConnection()
	defer db.CloseDb()

	api.ServeApi()
}
