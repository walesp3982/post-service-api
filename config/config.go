package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DatabaseUrl      string
	JWTSecret        string
	JWTAccessExpiry  int // In minutes
	JWTRefreshExpiry int // In Hours
}

func getEnvVariable(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic("Cannot get env variable: " + value)
	}
	return value
}

func convertInt(value string) int {
	convert, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return convert
}

func New() AppConfig {
	godotenv.Load()

	databaseUrl := getEnvVariable("DATABASE_URL")
	jwtSecret := getEnvVariable("JWT_SECRET")
	jwtAccessExpiry := convertInt(getEnvVariable("JWT_ACCESS_EXPIRY"))
	jwtRefreshExpiry := convertInt(getEnvVariable("JWT_REFRESH_EXPIRY"))

	return AppConfig{
		DatabaseUrl:      databaseUrl,
		JWTSecret:        jwtSecret,
		JWTAccessExpiry:  jwtAccessExpiry,
		JWTRefreshExpiry: jwtRefreshExpiry,
	}
}
