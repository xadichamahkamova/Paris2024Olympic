package load

import "github.com/spf13/viper"

type ServiceConfig struct {
	Host string
	Port int
}

type Config struct {
	ServerHost     string
	ServerPort     int
	UserService    ServiceConfig
	MedalService   ServiceConfig
	CountryService ServiceConfig
	EventService   ServiceConfig
	AthleteService ServiceConfig
	LiveService    ServiceConfig
}

func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		ServerHost: viper.GetString("server.host"),
		ServerPort: viper.GetInt("server.port"),

		UserService: ServiceConfig{
			Host: viper.GetString("services.user_service.host"),
			Port: viper.GetInt("services.user_service.port"),
		},
		MedalService: ServiceConfig{
			Host: viper.GetString("services.medal_service.host"),
			Port: viper.GetInt("services.medal_service.port"),
		},
		CountryService: ServiceConfig{
			Host: viper.GetString("services.country_service.host"),
			Port: viper.GetInt("services.country_service.port"),
		},
		EventService: ServiceConfig{
			Host: viper.GetString("services.event_service.host"),
			Port: viper.GetInt("services.event_service.port"),
		},
		AthleteService: ServiceConfig{
			Host: viper.GetString("services.athlete_service.host"),
			Port: viper.GetInt("services.athlete_service.port"),
		},
		LiveService: ServiceConfig{
			Host: viper.GetString("services.live_service.host"),
			Port: viper.GetInt("services.live_service.port"),
		},
	}
	return &cfg, nil
}
