package config

import "os"

type Config struct {
	RunType string
	Host    string
	Port    string
}

func GetConfig() *Config {
	runType := os.Getenv("RUN_TYPE")
	if runType == "" {
		runType = "Server"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		RunType: runType,
		Host:    host,
		Port:    port,
	}
}
