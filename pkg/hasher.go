package hasher

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// HashService is a simple interface for a service that returns a SHA256 hash of a
// given string
type HashService interface {
	HashSHA256(context.Context, string) (string, error)
}

type hashService struct{}

// HashSHA256 returns SHA256 hash of the provided string
func (hashService) HashSHA256(_ context.Context, input string) (string, error) {
	if input == "" {
		return "", ErrEmpty
	}
	hash := sha256.Sum256([]byte(input))
	hashSlice := hash[:]
	hashString := hex.EncodeToString(hashSlice)
	return hashString, nil
}

// ErrEmpty is returned when there is a blank input string to hash
var ErrEmpty = errors.New("Empty input string")

// NewHashService is a constructor to the hashService
func NewHashService() HashService {
	return hashService{}
}
