package load

import "github.com/spf13/viper"

type MongoConfig struct {
	Host       string
	Port       int
	Database   string
	Collection string
}

type Config struct {
	MongoConfig MongoConfig

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
		MongoConfig: MongoConfig{
			Host:       viper.GetString("mongo.host"),
			Port:       viper.GetInt("mongo.port"),
			Database:   viper.GetString("mongo.database"),
			Collection: viper.GetString("mongo.collection"),
		},
		ServerHost: viper.GetString("server.host"),
		ServerPort: viper.GetInt("server.port"),
	}
	return &cfg, nil
}
