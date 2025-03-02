package main

import (
	"log"
	"orchstrator/internal/config"
	"orchstrator/internal/server"
	"orchstrator/pkg/db/mongodb"
	"time"
)

func main() {

	// Reading config
	cfg, err := config.New("local.env")
	if err != nil {
		log.Fatalf("failed to read config: %s", err.Error())
	}

	// Connect to MongoDB
	timeoutMDB := 10 * time.Second
	mdb, err := mongodb.New(cfg.MongoDBUri, cfg.MongoDBDatabase, timeoutMDB)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %s", err.Error())
	}

	// Init server
	e := server.New(mdb)

	// Start server
	e.Start(cfg.RestServerPort)
}
