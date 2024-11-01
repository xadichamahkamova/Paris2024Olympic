package load 

import "github.com/spf13/viper"

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Config struct {
	Postgres PostgresConfig

	ServerHost string
	ServerPort int
}

func Load(path string) (*Config, error) {

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		Postgres: PostgresConfig{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetInt("postgres.port"),
			User:     viper.GetString("postgres.user"),
			Password: viper.GetString("postgres.password"),
			Database: viper.GetString("postgres.name"),
		},
		ServerHost: viper.GetString("server.host"),
		ServerPort: viper.GetInt("server.port"),
	}
	return &cfg, nil
}
