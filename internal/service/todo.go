package service

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	postgresdb "github.com/v7ktory/htmx+go/pkg/database/postgres/sqlc"
)

type TodoService struct {
	Pdb *postgresdb.Queries
}

func NewTodoService(Pdb *postgresdb.Queries) *TodoService {
	return &TodoService{
		Pdb: Pdb,
	}
}

func (s *TodoService) CreateTodo(title, description, email string) error {
	user, err := s.Pdb.GetUser(context.Background(), email)
	if err != nil {
		log.Println("Error getting user:", err)
	}

	todo := postgresdb.CreateTodoParams{
		Title:       title,
		Description: pgtype.Text{String: description, Valid: true},
		Completed:   pgtype.Bool{Bool: false, Valid: true},
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
		UserID:      user.ID,
	}

	if _, err := s.Pdb.CreateTodo(context.Background(), todo); err != nil {
		log.Println("Error creating todo:", err)
	}

	return nil
}
