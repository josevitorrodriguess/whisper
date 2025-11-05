package repository

import (
	"context"
	"errors"

	"github.com/josevitorrodriguess/whisper/server/internal/models"
	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) Delete(ctx context.Context, userID string) error {
	return r.DB.WithContext(ctx).Delete(&models.User{}, "id = ?", userID).Error
}

func (r *UserRepository) GetByFirebaseUID(ctx context.Context, firebaseUID string) (*models.User, error) {
	var user models.User
	err := r.DB.WithContext(ctx).Where("id = ?", firebaseUID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
