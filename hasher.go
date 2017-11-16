package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

type HashService interface {
	HashSHA256(context.Context, string) (string, error)
}

type hashService struct{}

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
