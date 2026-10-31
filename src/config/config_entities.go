package config

type Config struct {
	ApiConfig apiConfig
}

type apiConfig struct {
	url  string
	port string
}
