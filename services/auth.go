package services

import (
	"context"
	"errors"
	"time"
	"walesp3982/golang-post-api/model"
	"walesp3982/golang-post-api/repository"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	repository     *repository.RefreshTokenRepository
	userRepository *repository.UserRepository
	timeRefresh    uint // In hours
	timeAcess      uint // In minutes
	secretJWT      string
}

func NewAuthService(
	repo *repository.RefreshTokenRepository,
	userRepo repository.UserRepository,
	timeRefresh uint,
	timeAcess uint,
	secretJWT string,

) *AuthService {
	return &AuthService{
		repository:     repo,
		userRepository: &userRepo,
		timeRefresh:    timeRefresh,
		timeAcess:      timeAcess,
		secretJWT:      secretJWT,
	}
}

/**
 * Generate a refresh token with expired time pass in settings
 */
func (a *AuthService) Login(ctx context.Context, email string, password string) (string, error) {
	// Getting a user by email
	user := (*a.userRepository).GetByEmail(ctx, email)

	if user == nil {
		return "", errors.New("User not found error")
	}

	// Check credencials of user
	if !user.VerifyPassword(password) {
		return "", errors.New("Credencials incorrect")
	}

	// Create a new Refresh Token, With 6 hours for debug
	new_refresh := model.NewRefreshToken(user.Id, 120)

	// Save refresh token in repository
	(*a.repository).Save(ctx, &new_refresh)

	return new_refresh.Token, nil
}

func (a *AuthService) GetAccessToken(ctx context.Context, token string) (string, error) {
	// Get a refreshToken Data in db
	refreshToken := (*a.repository).GetByToken(ctx, token)
	if refreshToken == nil {
		return "", errors.New("Refresh Token not found")
	}

	if !refreshToken.IsValid() {
		return "", errors.New("Refresh Token invalid")
	}

	// Get user
	user := (*a.userRepository).GetById(ctx, refreshToken.Id)

	if user == nil {
		return "", errors.New("User not found")
	}
	claims := jwt.MapClaims{
		"name": user.Name,
		"sub":  user.Id.String(),
		"exp":  time.Now().Add(time.Duration(a.timeAcess)).Unix(),
		"iat":  time.Now().Unix(),
	}

	j := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := j.SignedString([]byte(a.secretJWT))
	if err != nil {
		return "", errors.New("Cannot create token")
	}

	return t, nil
}

func (a *AuthService) RegisterUser(ctx context.Context, name string, email string, password string) (*model.User, error) {
	user := model.NewUser(name, email, password)

	// Check if email have not register in db
	if (*a.userRepository).GetByEmail(ctx, email) != nil {
		return nil, errors.New("Email was registred in db")
	}
	// Save user in db
	error := (*a.userRepository).Save(ctx, &user)
	if error != nil {
		return nil, errors.New("Error to create user")
	}

	return &user, nil
}

func (a *AuthService) Logout(ctx context.Context, token string) error {
	// Get RefreshToken of db
	refreshToken := (*a.repository).GetByToken(ctx, token)
	if refreshToken == nil {
		return errors.New("refresh Token not found")
	}
	if !refreshToken.IsValid() {
		return errors.New("token has expired")
	}

	refreshToken.Revolked = false
	err := (*a.repository).Update(ctx, refreshToken)
	if err != nil {
		return errors.New("Cannot update refreshToken")
	}
	return nil
}
