package app

import (
	"crudapp/internal/transport"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pressly/goose/v3"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Driver string `yaml:"driver"`
		Dsn    string `yaml:"dsn"`
	} `yaml:"database"`
	Migrations struct {
		Dir string `yaml:"dir"`
	} `yaml:"migrations"`
}

type App struct {
	config Config
}

func NewApp(configPath string) (*App, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения конфигурации: %w", err)
	}

	var cfg Config
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("ошибка парсинга конфигурации: %w", err)
	}
	return &App{cfg}, nil
}

func (a *App) Run() {
	db, err := sql.Open(a.config.Database.Driver, a.config.Database.Dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	if err := goose.Up(db, a.config.Migrations.Dir); err != nil {
		log.Fatalf("Ошибка миграций: %v", err)
	}

	e := echo.New()
	transport.RegisterRoutes(e, db)
	e.Logger.Fatal(e.Start(a.config.Server.Port))
}