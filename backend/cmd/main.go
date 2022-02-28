package main

import (
	api "backend/cmd/api"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func setupConfig() api.Config {
	var cfg api.Config
	flag.IntVar(&cfg.Port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "development", "Application environment (development|production)")
	flag.StringVar(&cfg.Dsn, "dsn", "postgres://rcontreras@localhost/go_ticket?sslmode=disable", "Postgres connection string")
	flag.Parse()
	return cfg
}

func main() {

	cfg := setupConfig()
	app := api.SetupApp(cfg)
	err := app.OpenDB()
	if err != nil {
		app.Logger.Fatal(err)
	}
	defer app.CloseDB()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Println("Starting server on port", cfg.Port)

	err = srv.ListenAndServe()

	if err != nil {
		log.Println(err)
	}
}
