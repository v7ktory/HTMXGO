package service

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/v7ktory/htmx+go/internal/model"
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

func (s *TodoService) CreateTodo(ctx context.Context, title, description, email string) error {
	user, err := s.Pdb.GetUser(ctx, email)
	if err != nil {
		log.Println("Error getting user:", err)
		return err
	}

	todo := postgresdb.CreateTodoParams{
		Title:       title,
		Description: pgtype.Text{String: description, Valid: true},
		Completed:   pgtype.Bool{Bool: false, Valid: true},
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
		UserID:      user.ID,
	}

	if _, err := s.Pdb.CreateTodo(ctx, todo); err != nil {
		log.Println("Error creating todo:", err)
		return err
	}

	return nil
}

func (s *TodoService) GetTodos(ctx context.Context, id int32) ([]model.Todo, error) {
	todosDB, err := s.Pdb.GetTodos(ctx, id)
	if err != nil {
		log.Println("Error getting todos:", err)
		return nil, err
	}

	var todos []model.Todo
	for _, t := range todosDB {
		todo := model.Todo{
			Title:       t.Title,
			Description: t.Description.String,
			Completed:   t.Completed.Bool,
			CreatedAt:   t.CreatedAt.Time,
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, id int32) error {
	err := s.Pdb.DeleteTodo(ctx, id)
	if err != nil {
		log.Println("Error deleting todo:", err)
		return err
	}
	return nil
}