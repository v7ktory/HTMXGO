package session

import (
	"context"
	"encoding/json"
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
	err := s.Rdb.Set(context.Background(), sessionID, string(jsonData), 24*time.Hour).Err()
	if err != nil {
		return "", err
	}
	return sessionID, nil
}

func (s *SessionManager) GetSession(session string) (*model.UserSession, error) {

	data, err := s.Rdb.Get(context.Background(), session).Result()
	if err != nil {
		return nil, err
	}

	var userSession model.UserSession
	err = json.Unmarshal([]byte(data), &userSession)
	if err != nil {
		return nil, err
	}

	return &userSession, nil

}
