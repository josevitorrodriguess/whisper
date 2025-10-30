package services

import (
	"github.com/josevitorrodriguess/whisper/server/internal/models"
	"github.com/josevitorrodriguess/whisper/server/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(u *models.User) (*models.User, error) {
	if u == nil {
		return nil, nil
	}
	if err := s.repo.CreateUser(u); err != nil {
		return nil, err
	}
	return s.repo.GetByFirebaseUID(u.ID)
}


func (s *UserService) DeleteUser(userID string) error {
	return s.repo.DeleteUser(userID)
}