package env

type Config struct {
	Port              int    `env:"PORT,default=3001"`
	SessionServiceUrl string `env:"SESSION_SERVICE_URL,default=localhost:50001"`
	UsersServiceUrl   string `env:"USERS_SERVICE_URL,default=localhost:50000"`
}
