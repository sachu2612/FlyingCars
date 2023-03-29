package postgres

import (
	"errors"

	"github.com/example/car-api/models"
	"github.com/example/car-api/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (us *UserService) CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return us.userRepository.CreateUser(user)
}

func (us *UserService) AuthenticateUser(username string, password string) (string, error) {
	user, err := us.userRepository.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	token, err := generateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateToken(user *models.User) (string, error) {
	// TODO: Implement JWT token generation
	return "", nil
}
