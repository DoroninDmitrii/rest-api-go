package repository

import (
	restapitodo "rest-api-todo"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user restapitodo.User) (int, error)
	GetUser(username, password string) (restapitodo.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
