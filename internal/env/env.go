package env

type Config struct {
	Port            int    `env:"PORT,default=3001"`
	UsersServiceUrl string `env:"USERS_SERVICE_URL,default=localhost:50000"`
}
