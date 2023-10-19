package config

// ServerConfig
type ServerConfig struct {
	Name            string `mapstructure:"name"`
	Env             string `mapstructure:"env"`
	Level           string `mapstructure:"level"`
	ShutdownTimeout int    `mapstructure:"shutdown_timeout"`
}

// GinConfig
type GinConfig struct {
	HttpPort string `mapstructure:"http_port"`
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

func GetConfig() ConfigSetup {
	return config
}

func GetServerConfig() ServerConfig {
	return config.ServerConfig
}

func GetGinConfig() GinConfig {
	return config.GinConfig
}

func GetDBConfig() DBConfig {
	return config.DBConfig
}
