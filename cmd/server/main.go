package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jerryddie77/jerry_notes/internal/config"
	"github.com/jerryddie77/jerry_notes/internal/handler"
	_ "github.com/lib/pq"
)

func initDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func main() {

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("loadConfig: %v", err)
	}

	db, err := initDB(cfg)
	if err != nil {
		log.Fatalf("initDB %v", err)
	}
 
	h := handler.NewHandler()
	r := handler.NewRouter(h)
	
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("run server: %v", err)
	}

}
