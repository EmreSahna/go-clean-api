package config

import "flag"

const version = "1.0.0"

type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn string
	}
	Version string
}

func LoadConfig() *Config {
	cfg := &Config{}
	flag.IntVar(&cfg.Port, "port", 8080, "Server port")
	flag.StringVar(&cfg.Env, "env", "development", "App environment (development|staging|production)")
	flag.StringVar(&cfg.Db.Dsn, "dsn", "postgres://postgres:postgres@localhost/clean_api", "Postgres DSN")
	flag.StringVar(&cfg.Version, "version", version, "App version")
	flag.Parse()
	return cfg
}
