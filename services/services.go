package services

import (
	"walesp3982/golang-post-api/config"
	"walesp3982/golang-post-api/repository"
)

type Service struct {
	Auth AuthService
}

func New(repository repository.Repository, config config.AppConfig) Service {
	return Service{
		Auth: *NewAuthService(
			&repository.RefreshToken,
			repository.User,
			uint(config.JWTRefreshExpiry),
			uint(config.JWTAccessExpiry),
			config.JWTSecret),
	}
}
