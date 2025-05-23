package service

import (
	restapitodo "rest-api-todo"
	"rest-api-todo/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list restapitodo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
