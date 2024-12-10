package config

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}
