package repositories

import (
	"github.com/putrapratamanst/ecommerce/user-service/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

type UserRepository interface {
	FindByEmailOrPhone(emailOrPhone string) (*models.User, error)
	CreateUser(user *models.User) error
	FindByUserID(userID string) (*models.User, error)
}

func (r *userRepository) FindByEmailOrPhone(emailOrPhone string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ? OR phone = ?", emailOrPhone, emailOrPhone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}


func (r *userRepository) FindByUserID(userID string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}