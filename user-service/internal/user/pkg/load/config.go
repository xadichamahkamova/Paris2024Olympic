package load

import (
	"github.com/spf13/viper"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type RedisConfig struct {
	Host string
	Port int
}

type Config struct {
	Postgres        PostgresConfig
	Redis           RedisConfig
	UserServiceHost string
	UserServicePort int
}

func Load(path string) (*Config, error) {

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

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
		Redis: RedisConfig{
			Host: viper.GetString("redis.host"),
			Port: viper.GetInt("redis.port"),
		},
		UserServiceHost: viper.GetString("server.host"),
		UserServicePort: viper.GetInt("server.port"),
	}
	return &cfg, nil
}
