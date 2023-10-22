package config

// ServerConfig
type ServerConfig struct {
	Name            string `mapstructure:"name"`
	Env             string `mapstructure:"env"`
	Level           string `mapstructure:"level"`
	ShutdownTimeout int    `mapstructure:"shutdown_timeout"`
}

// DBConfig
type DBConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Address  string `mapstructure:"address"`
	Database string `mapstructure:"database"`

	LogMode        bool `mapstructure:"log_mode"`
	MaxIdle        int  `mapstructure:"max_idle"`
	MaxOpen        int  `mapstructure:"max_open"`
	ConnMaxLifeMin int  `mapstructure:"conn_max_life_min"`
}

// scanConfig
type ScannerConfig struct {
	PipelineNumber int `mapstructure:"pipeline_number"`
	ScanBlockFrom  int `mapstructure:"scan_block_from"`
}

func GetConfig() ConfigSetup {
	return config
}

func GetServerConfig() ServerConfig {
	return config.ServerConfig
}

func GetDBConfig() DBConfig {
	return config.DBConfig
}

func GetScannerConfig() ScannerConfig {
	return config.ScannerConfig
}
