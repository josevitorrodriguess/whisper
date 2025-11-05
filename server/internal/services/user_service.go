package services

import (
	"context"
	"errors"

	"github.com/josevitorrodriguess/whisper/server/internal/models"
	"github.com/josevitorrodriguess/whisper/server/internal/repository"
	"github.com/josevitorrodriguess/whisper/server/internal/validations"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) EnsureUserExists(ctx context.Context, uid, email, username, photoURL string) error {
	user := &models.User{
		ID:       uid,
		Email:    email,
		Username: username,
		PhotoURL: photoURL,
	}

	if err := validations.UserIsValid(*user); err != nil {
		return err
	}

	_, err := s.repo.GetByFirebaseUID(ctx, uid)
	if err == nil {
		return nil
	}

	if !errors.Is(err, repository.ErrUserNotFound) {
		return err
	}

	return s.repo.Create(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, userID string) error {
	return s.repo.Delete(ctx, userID)
}
