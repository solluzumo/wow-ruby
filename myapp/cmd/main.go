package main

import (
	"log"
	_ "wow-ruby/docs"
	"wow-ruby/internal/app"
	"wow-ruby/internal/config"
	_ "wow-ruby/internal/handlers"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	log.Println("Config is loaded.")

	application, err := app.New(cfg)
	if err != nil {
		log.Fatal("Failed to create app:", err)
	}

	log.Println("App is created.")

	if err := application.Run(); err != nil {
		log.Fatal("Failed to run app:", err)
	}
	log.Println("App is running.")
}
