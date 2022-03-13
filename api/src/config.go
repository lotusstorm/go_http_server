package src

import (
	"fmt"
	"os"
)

type Configuration struct {
	Host       string
	Port       string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	DBString   string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (c *Configuration) loadEnv() {
	c.Host = getEnv("HOST", "0.0.0.0")
	c.Port = getEnv("PORT", "5000")
	c.DBUser = getEnv("DB_USER", "pg")
	c.DBPassword = getEnv("DB_PASSWORD", "pass")
	c.DBHost = getEnv("DB_HOST", "db")
	c.DBPort = getEnv("DB_PORT", "5432")
	c.DBName = getEnv("DB_NAME", "crud")
	c.DBString = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}
