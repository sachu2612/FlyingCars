package postgres

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/example/car-api/models"
)

type AuthService struct {
	userRepo    UserRepository
	secretKey   []byte
	expDuration time.Duration
}

func NewAuthService(ur UserRepository, secretKey string, expDuration time.Duration) *AuthService {
	return &AuthService{
		userRepo:    ur,
		secretKey:   []byte(secretKey),
		expDuration: expDuration,
	}
}

func (as *AuthService) GenerateToken(u *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.ID,
		"exp": time.Now().Add(as.expDuration).Unix(),
	})
	tokenString, err := token.SignedString(as.secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (as *AuthService) VerifyToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return as.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int64(claims["sub"].(float64))
		user, err := as.userRepo.GetUserByID(userID)
		if err != nil {
			return nil, err
		}
		return user, nil
	} else {
		return nil, err
	}
}
