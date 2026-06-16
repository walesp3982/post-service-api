package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DatabaseUrl      string
	JWTSecret        string
	JWTAccessExpiry  uint // In minutes
	JWTRefreshExpiry uint // In Hours
}

func getEnvVariable(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic("Cannot get env variable: " + value)
	}
	return value
}

func convertUint(value string) uint {
	convert, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return uint(convert)
}

func New() AppConfig {
	godotenv.Load()

	databaseUrl := getEnvVariable("DATABASE_URL")
	jwtSecret := getEnvVariable("JWT_SECRET")
	jwtAccessExpiry := convertUint(getEnvVariable("JWT_ACCESS_EXPIRY"))
	jwtRefreshExpiry := convertUint(getEnvVariable("JWT_REFRESH_EXPIRY"))

	return AppConfig{
		DatabaseUrl:      databaseUrl,
		JWTSecret:        jwtSecret,
		JWTAccessExpiry:  jwtAccessExpiry,
		JWTRefreshExpiry: jwtRefreshExpiry,
	}
}
