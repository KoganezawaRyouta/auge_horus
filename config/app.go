package config

import "os"

type AppServerConfig struct {
	LogFile string
	Port    string
	PidFile string
}

func AppConfigNew() *AppServerConfig {
	return &AppServerConfig{
		LogFile: os.Getenv("APP_LOG_PATH"),
		Port:    os.Getenv("APP_PORT"),
		PidFile: os.Getenv("APP_PID_FILE"),
	}
}
