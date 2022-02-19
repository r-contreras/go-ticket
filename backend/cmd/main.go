package main

import (
	api "backend/cmd/api"
	implementations "backend/cmd/implementations"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func setupConfig() api.Config {
	var cfg api.Config
	flag.IntVar(&cfg.Port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "development", "Application environment (development|production)")
	flag.Parse()
	return cfg
}

func setupApp(cfg api.Config) *implementations.Application {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app:= &implementations.Application{
		Config:  cfg,
		Logger: logger,
	}
	return app
}
func main(){
	
	cfg := setupConfig()
	app := setupApp(cfg)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
		Handler: app.Routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 30*time.Second,
	}

	app.Logger.Println("Starting server on port", cfg.Port)

	err := srv.ListenAndServe()

	if err != nil {
		log.Println(err)
	}
}