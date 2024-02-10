package api

const baseURL = "https://pokeapi.co/api/v2"

type Config struct {
	Next     string
	Previous string
}

func NewConfig() *Config {
	return &Config{}
}
