package main

import (
	"github.com/BrownieBrown/schafkopf/internal/database/sqlite"
	"github.com/BrownieBrown/schafkopf/internal/utils"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := sqlite.Init(cfg.SqliteConfig)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	defer func(db *sqlite.Database) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing database: %v", err)
		}
	}(db)

}
