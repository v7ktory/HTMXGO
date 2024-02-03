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
	CreateTodo(ctx context.Context, title, description, email string) error
	GetTodos(ctx context.Context, userID int32) ([]model.Todo, error)
	UpdateTodo(ctx context.Context, userID, todoID int32) error
	DeleteTodo(ctx context.Context, userID, todoID int32) error
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
