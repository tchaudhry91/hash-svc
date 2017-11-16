package main

import (
	"context"
	"testing"
)

var hasher hashService
var ctx context.Context

const (
	string1   = "hello"
	response1 = "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
)

func init() {
	hasher = hashService{}
}

func TestHash(t *testing.T) {
	response, err := hasher.HashSHA256(ctx, string1)
	if err != nil {
		t.Fatal("Error while Hashing:", err)
	}
	if response != response1 {
		t.Fatalf("Hash failed for %s.\nExpected:%s\nReceived:%s\n",
			string1,
			response1,
			response,
		)
	}
}

func TestEmpty(t *testing.T) {
	_, err := hasher.HashSHA256(ctx, "")
	if err != nil {
		if err == ErrEmpty {
			return
		}
		t.Fatal("Error while Hashing:", err)
	}
}
