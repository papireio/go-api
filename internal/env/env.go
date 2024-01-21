package env

type Config struct {
	Port int `env:"PORT,default=3001"`
}
