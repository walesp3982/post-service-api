package services

import (
	"context"
	"errors"
	"walesp3982/golang-post-api/model"
	"walesp3982/golang-post-api/repository"

	"github.com/google/uuid"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) GetById(ctx context.Context, id uuid.UUID) *model.User {
	return s.repository.GetById(ctx, id)
}

func (s *UserService) ChangePassword(
	ctx context.Context,
	id uuid.UUID,
	newPassword string,
	oldPassword string) error {

	// Get User beetween password
	user := s.repository.GetById(ctx, id)

	if user == nil {
		return errors.New("User not found")
	}

	// Verify old password
	if !user.VerifyPassword(oldPassword) {
		return errors.New("Incorrect password")
	}
	user.ChangePassword(newPassword)
	err := s.repository.Update(ctx, user)

	return err
}

func (s *UserService) ChangeMetadata(ctx context.Context, id uuid.UUID, name string) error {
	user := s.repository.GetById(ctx, id)
	if user == nil {
		return errors.New("User not found")
	}

	user.Name = name
	err := s.repository.Update(ctx, user)
	return err
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	user := s.repository.GetById(ctx, id)
	if user == nil {
		return errors.New("User not found")
	}

	err := s.repository.Delete(ctx, id)
	return err
}
