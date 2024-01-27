package service

import (
	"github.com/redis/go-redis/v9"
	postgresdb "github.com/v7ktory/htmx+go/pkg/database/postgres/sqlc"
)

type Auth interface {
	Signup(name, email, password string) (postgresdb.User, error)
	Login(email, password string) (string, error)
	SignOut(sessionID string) error
}
type Service struct {
	Auth
}

func NewService(Rdb *redis.Client, Pdb *postgresdb.Queries) *Service {
	return &Service{
		Auth: NewAuthService(Rdb, Pdb),
	}
}
