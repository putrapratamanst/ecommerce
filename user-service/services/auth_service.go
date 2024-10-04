package services

import (
	"errors"

	"github.com/putrapratamanst/ecommerce/user-service/models"
	"github.com/putrapratamanst/ecommerce/user-service/repositories"
	"github.com/putrapratamanst/ecommerce/user-service/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(emailOrPhone, password string) (string, error)
	Register(user *models.User) error
	Detail(userID string) (*models.User, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Login(emailOrPhone, password string) (string, error) {
	user, err := s.userRepo.FindByEmailOrPhone(emailOrPhone)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	tokenString, err := utils.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *authService) Register(user *models.User) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.CreateUser(user)
}

func (s *authService) Detail(userID string) (*models.User,error) {
	user, err := s.userRepo.FindByUserID(userID)
	if err != nil {
		return &models.User{}, errors.New("user not found")
	}

	return user, nil
}