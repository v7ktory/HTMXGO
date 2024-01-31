package service

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/v7ktory/htmx+go/internal/model"
	postgresdb "github.com/v7ktory/htmx+go/pkg/database/postgres/sqlc"
)

type Auth interface {
	Signup(ctx context.Context, user *model.User) (postgresdb.User, error)
	Login(ctx context.Context, user *model.User) (string, error)
	SignOut(ctx context.Context, sessionID string) error
}

type Todo interface {
	CreateTodo(title, description, email string) error
	// GetTodos(id int32) ([]model.Todo, error)
	// UpdateTodo(title, description string, id int32, completed bool) error
	// DeleteTodo(id int32) error
}
type Service struct {
	Auth
	Todo
}

func NewService(Rdb *redis.Client, Pdb *postgresdb.Queries) *Service {
	return &Service{
		Auth: NewAuthService(Rdb, Pdb),
		Todo: NewTodoService(Pdb),
	}
}
