package app

type DBConfig struct {
	User     string
	Password string
	Host     string
	Name     string
	Port     int
}

type BaseCommConfig struct {
	Port   string
	Domain string
}

type GRPCConfig struct {
	Subscriber_Auth BaseCommConfig
}

type CommConfig struct {
	REST BaseCommConfig
	GRPC GRPCConfig
}

type AppConfig struct {
	CommConfig
	JWT_secret      string
	JWT_exp_minutes int
}

type Config struct {
	App AppConfig
	DB  DBConfig
}

var appConfig Config = Config{
	App: AppConfig{
		CommConfig: CommConfig{
			REST: BaseCommConfig{
				Port:   "8000",
				Domain: "localhost",
			},
			GRPC: GRPCConfig{
				Subscriber_Auth: BaseCommConfig{
					Domain: "localhost",
					Port:   "50000",
				},
			},
		},
		JWT_secret:      "secret",
		JWT_exp_minutes: 60,
	},
	DB: DBConfig{
		User:     "user",
		Password: "password",
		Name:     "user_db",
		Port:     5000,
		Host:     "localhost",
	},
}

func GetConfig() Config {
	return appConfig
}
