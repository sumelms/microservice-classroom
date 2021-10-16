package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	classroom "github.com/sumelms/microservice-classroom/internal/classroom/database"
	subscription "github.com/sumelms/microservice-classroom/internal/subscription/database"
	"github.com/sumelms/microservice-classroom/pkg/config"
	database "github.com/sumelms/microservice-classroom/pkg/database/gorm"
	"github.com/sumelms/microservice-classroom/pkg/seed"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("Unable to load the configuration: %s", err)
	}

	// Database
	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err)
	}
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)

	for _, s := range allSeeds() {
		fmt.Printf("Executing seed '%s'...\n", s.Name)
		if err := s.Run(db); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", s.Name, err) // nolint: gocritic
		}
	}
}

func loadConfig() (*config.Config, error) {
	// Configuration
	configPath := os.Getenv("SUMELMS_CONFIG_PATH")
	if configPath == "" {
		configPath = "./config.yml"
	}

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func allSeeds() []seed.Seed {
	return []seed.Seed{
		classroom.Classrooms(),
		subscription.Subscriptions(),
	}
}
