package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/v7ktory/htmx+go/internal/model"
	postgresdb "github.com/v7ktory/htmx+go/pkg/database/postgres/sqlc"
	"github.com/v7ktory/htmx+go/pkg/hash"
)

type AuthService struct {
	Rdb *redis.Client
	Pdb *postgresdb.Queries
}

func NewAuthService(Rdb *redis.Client, Pdb *postgresdb.Queries) *AuthService {
	return &AuthService{
		Rdb: Rdb,
		Pdb: Pdb,
	}
}

func (s *AuthService) Signup(name, email, password string) (postgresdb.User, error) {
	// Hash the password for security
	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		return postgresdb.User{}, fmt.Errorf("failed to hash password: %v", err)
	}

	// Create a new user in the database
	u, err := s.Pdb.CreateUser(context.Background(), postgresdb.CreateUserParams{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	})
	if err != nil {
		return postgresdb.User{}, fmt.Errorf("failed to create user: %v", err)
	}

	return u, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	// Retrieve user from the database
	user, err := s.Pdb.GetUser(context.Background(), email)
	if err != nil {
		return "", err
	}

	// Check if the provided password matches the user's password
	if !hash.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid password")
	}

	// Generate a new session ID
	sessionID := uuid.NewString()
	// Create a user session object
	userSession := model.UserSession{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	// Marshal user session object to JSON
	jsonData, err := json.Marshal(userSession)
	if err != nil {
		return "", err
	}

	// Set the user session in the Redis database with a 24-hour expiration
	err = s.Rdb.Set(context.Background(), sessionID, string(jsonData), 24*time.Hour).Err()
	if err != nil {
		return "", err
	}

	// Return the generated session ID
	return sessionID, nil
}
func (s *AuthService) SignOut(sessionID string) error {
	ctx := context.Background()

	// Delete the session with the given session ID from the Redis database.
	err := s.Rdb.Del(ctx, sessionID).Err()
	if err != nil {
		return fmt.Errorf("failed to revoke session: %w", err)
	}

	return nil
}
