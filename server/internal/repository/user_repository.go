package repository

import (
	"github.com/josevitorrodriguess/whisper/server/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) DeleteUser(userID uint) error {
	return r.DB.Delete(&models.User{}, userID).Error
}

func (r *UserRepository) GetByFirebaseUID(firebaseUID string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("id = ?", firebaseUID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
