package service

import (
	"crypto/sha1"
	"fmt"
	restapitodo "rest-api-todo"
	"rest-api-todo/pkg/repository"
)

const salt = "dsdsdwqwewel;fldl;f;ld"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user restapitodo.User) (int, error) {
  user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

  return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
