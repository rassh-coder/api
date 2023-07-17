package service

import (
	"MyApi/pkg/models"
	"MyApi/pkg/repository"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	salt           = "qqqwe12sdfvhhyjldcxv43ogv"
	secretKey      = "qqqasde123dbmvcs741asd"
	expirationTime = 10 * time.Minute
	adminKey       = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.Eu0izcSRVwn901UcFpwYzmhaBKw4BJHjRwYZQfpzPek"
)

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (s *AuthService) CreateUser(input models.UserInputFields) (int64, error) {
	input.Password = generatePasswordHash(input.Password)
	return s.repos.CreateUser(input)
}

func (s *AuthService) SignIn(creds models.UserInputCreds) (string, error) {
	creds.Password = generatePasswordHash(creds.Password)
	user, err := s.repos.GetUser(creds)
	if err != nil {
		return "", err
	}
	var token string
	if creds.UserName == "rassh" {
		token, err = generateSUToken(user)
		if err != nil {
			return "", err
		}
	} else {
		token, err = generateToken(user)
		if err != nil {
			return "", err
		}
	}

	return token, nil
}

func generatePasswordHash(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))

	return fmt.Sprintf("%x", hasher.Sum([]byte(salt)))
}

func generateToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"ExpiredAt":  time.Now().Add(expirationTime),
		"IssuedAt":   time.Now(),
		"id":         user.Id,
		"user_name":  user.UserName,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
	}

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func generateSUToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{"id": user.Id}

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func parseToken(token string) (bool, error) {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return nil, nil
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func ParseSUToken(token string) bool {
	if token == adminKey {
		return true
	}
	return false
}
