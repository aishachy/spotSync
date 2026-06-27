package main

import (
	"spotSync/internal/config"
	"spotSync/internal/server"
)

func main() {
	cfg := config.LoadEnv()

	db := config.ConnectDatabase(cfg)

	server.Start(db, cfg)
}
