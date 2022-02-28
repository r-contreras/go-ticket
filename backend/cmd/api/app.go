package api

import (
	"backend/cmd/infrastructure"
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Application struct {
	Config Config
	Logger *log.Logger
	Model  *infrastructure.DbModel
}

type Config struct {
	Port int
	Env  string
	Dsn  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func SetupApp(cfg Config) *Application {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &Application{
		Config: cfg,
		Logger: logger,
		Model:  nil,
	}
	return app
}

func (app *Application) OpenDB() error {
	db, err := sql.Open("postgres", app.Config.Dsn)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return err
	}
	app.Model = infrastructure.GetDbModel(db)

	return nil
}

func (app *Application) CloseDB() error {
	err := app.Model.DB.Close()

	if err != nil {
		return err
	}
	return nil
}
