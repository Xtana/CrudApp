package main

import (
	"crudapp/internal/app"
	"log"
	"os"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config.yaml"
	}

	application, err := app.NewApp(configPath)
	if err != nil {
		log.Fatalf("Ошибка инициализации приложения: %v", err)
	}

	application.Run()
}
