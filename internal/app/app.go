package app

import (
	"crudapp/internal/transport"
	"database/sql"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pressly/goose/v3"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml: "server"`
	Database struct {
		Driver 	string `yaml:"driver"`
		Dsn		string `yaml:"dsn"`
	} `yaml:"database"`
	Migrations struct {
		Dir string `yaml:"dir"`
	} `yaml:"migration"`
}

type App struct {
	config Config
}

func NewApp(configPath string) (*App, error) {
	date, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err = yaml.Unmarshal(date, &cfg); err != nil {
		return nil, err
	}
	return &App{cfg}, nil
}

func (a *App) Ran() {
	db, err := sql.Open(a.config.Database.Driver, a.config.Database.Dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к бд:", err)
	}
	defer db.Close()

	if err := goose.Up(db, a.config.Migrations.Dir); err != nil {
		log.Fatal("Ошибка миграций:", err)
	}

	e := echo.New()
	transport.RegisterRoutes(e, db)
	e.Logger.Fatal(e.Start(a.config.Server.Port))
}