package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/putrapratamanst/ecommerce/user-service/config"
	"github.com/putrapratamanst/ecommerce/user-service/models"
	"github.com/putrapratamanst/ecommerce/user-service/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(emailOrPhone, password string) (string, error)
	Register(user *models.User) error
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(config.JWT_SECRET)
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