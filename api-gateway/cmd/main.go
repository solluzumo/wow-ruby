package main

import (
	"log"

	"github.com/solluzumo/wow-ruby/gateway/internal/app"
	"github.com/solluzumo/wow-ruby/gateway/internal/config"
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
	defer application.AuthServiceRPC.Close()
	log.Println("App is created.")

	if err := application.Run(); err != nil {
		log.Fatal("Failed to run app:", err)
	}

	log.Println("App is running.")

}
