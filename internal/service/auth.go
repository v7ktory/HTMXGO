package service

import (
	"context"
	"encoding/json"
	"log"
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

func (s *AuthService) Signup(ctx context.Context, user *model.User) (postgresdb.User, error) {
	// Hash the password for security
	hashedPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing password: %v\n", err)
		return postgresdb.User{}, err
	}

	// Create a new user in the database
	u, err := s.Pdb.CreateUser(ctx, postgresdb.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	})
	if err != nil {
		log.Println("Error creating user:", err)
		return postgresdb.User{}, err
	}

	return u, nil
}

func (s *AuthService) Login(ctx context.Context, user *model.User) (string, error) {
	// Retrieve user from the database
	u, err := s.Pdb.GetUser(ctx, user.Email)
	if err != nil {
		log.Println("Error getting user:", err)
		return "", err
	}

	// Check if the provided password matches the user's password
	if !hash.CheckPasswordHash(user.Password, u.Password) {
		log.Println("Error checking password:", err)
		return "", err
	}

	// Generate a new session ID
	sessionID := uuid.NewString()
	// Create a user session object
	userSession := model.UserSession{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
	// Marshal user session object to JSON
	jsonData, err := json.Marshal(userSession)
	if err != nil {
		log.Println("Error marshaling object:", err)
		return "", err
	}

	// Set the user session in the Redis database with a 24-hour expiration
	err = s.Rdb.Set(ctx, sessionID, string(jsonData), 24*time.Hour).Err()
	if err != nil {
		log.Println("Error setting session:", err)
		return "", err
	}

	// Return the generated session ID
	return sessionID, nil
}
func (s *AuthService) SignOut(ctx context.Context, sessionID string) error {
	// Delete the session with the given session ID from the Redis database.
	err := s.Rdb.Del(ctx, sessionID).Err()
	if err != nil {
		log.Println("Error deleting session:", err)
		return err
	}

	return nil
}
