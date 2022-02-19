package api

type Config struct{
	Port int
	Env string
}

type AppStatus struct{
	Status string `json:"status"`
	Environment string `json:"environment"`
	Version string `json:"version"`
}