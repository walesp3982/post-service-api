package services

import (
	"walesp3982/golang-post-api/config"
	"walesp3982/golang-post-api/repository"
)

type Service struct {
	Auth AuthService
	User UserService
	Post PostService
}

func New(repository repository.Repository, config config.AppConfig) Service {
	return Service{
		Auth: *NewAuthService(
			&repository.RefreshToken,
			repository.User,
			config.JWTRefreshExpiry,
			config.JWTAccessExpiry,
			config.JWTSecret),
		User: *NewUserService(
			repository.User,
		),
		Post: *NewPostService(repository.Post),
	}
}
