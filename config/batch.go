package config

import "os"

type BatchConfig struct {
	LogFile string
}

func BatchConfigNew() *BatchConfig {
	return &BatchConfig{LogFile: os.Getenv("BATCH_LOG_PATH")}
}
