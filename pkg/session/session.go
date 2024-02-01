package session

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/v7ktory/htmx+go/internal/model"
)

type SessionManager struct {
	Rdb *redis.Client
}

func NewSessionManager(rdb *redis.Client) *SessionManager {
	return &SessionManager{
		Rdb: rdb,
	}
}

func (s *SessionManager) GenerateSession(data model.UserSession) (string, error) {
	sessionID := uuid.NewString()
	jsonData, _ := json.Marshal(data)
	ctx := context.Background()
	err := s.Rdb.Set(ctx, sessionID, string(jsonData), 24*time.Hour).Err()
	if err != nil {
		log.Println("Error setting session:", err)
		return "", err
	}
	return sessionID, nil
}

func (s *SessionManager) GetSession(sessionID string) (*model.UserSession, error) {
	data, err := s.Rdb.Get(context.Background(), sessionID).Result()
	if err != nil {
		log.Println("Error getting session:", err)
		return nil, err
	}

	var userSession model.UserSession
	if err := json.Unmarshal([]byte(data), &userSession); err != nil {
		log.Println("Error unmarshaling session:", err)
		return nil, err
	}

	return &userSession, nil
}
