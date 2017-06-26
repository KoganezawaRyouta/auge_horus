package config

import "os"

type ApiServerConfig struct {
	LogFile string
	Port    string
	PidFile string
}

func ApiConfigNew() *ApiServerConfig {
	return &ApiServerConfig{
		LogFile: os.Getenv("API_LOG_PATH"),
		Port:    os.Getenv("API_PORT"),
		PidFile: os.Getenv("API_PID_FILE"),
	}
}
