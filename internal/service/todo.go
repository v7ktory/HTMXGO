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

func (s *TodoService) GetTodos(ctx context.Context, userID int32) ([]model.Todo, error) {
	todosDB, err := s.Pdb.GetTodos(ctx, userID)
	if err != nil {
		log.Println("Error getting todos:", err)
		return nil, err
	}

	var todos []model.Todo
	for _, t := range todosDB {
		todo := model.Todo{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description.String,
			Completed:   t.Completed.Bool,
			CreatedAt:   t.CreatedAt.Time,
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, userID, todoID int32) error {
	err := s.Pdb.UpdateTodo(ctx, postgresdb.UpdateTodoParams{
		Completed: pgtype.Bool{Bool: true, Valid: true},
		ID:        todoID,
		UserID:    userID,
	})

	return err
}
func (s *TodoService) DeleteTodo(ctx context.Context, userID, todoID int32) error {
	err := s.Pdb.DeleteTodo(ctx, postgresdb.DeleteTodoParams{
		UserID: userID,
		ID:     todoID,
	})

	return err
}
