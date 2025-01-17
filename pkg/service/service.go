package service

import (
	restapitodo "rest-api-todo"
	"rest-api-todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user restapitodo.User) (int, error)
}

type TodoList interface {

}

type TodoItem interface {

}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
