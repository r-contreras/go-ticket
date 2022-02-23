package api

import (
	"log"
)

type Application struct {
	Config Config
	Logger *log.Logger
}

type Config struct {
	Port int
	Env  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}
