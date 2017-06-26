package config

type Config struct {
	DB        *DBConfig
	Batch     *BatchConfig
	AppServer *AppServerConfig
	ApiServer *ApiServerConfig
}

func ConfigNew() *Config {
	return &Config{DB: DBConfigNew(), Batch: BatchConfigNew(), ApiServer: ApiConfigNew(), AppServer: AppConfigNew()}
}
