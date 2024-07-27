package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Payload contains the payload data of the token
type Payload struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// NewPayload creates a new Payload for a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return payload, nil
}
