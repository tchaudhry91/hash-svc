package hasher

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

// MakeHashSHA256Endpoint returns a go-kit endpoint for the hashService
func MakeHashSHA256Endpoint(svc HashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(hashSHA256Request)
		v, err := svc.HashSHA256(ctx, req.S)
		if err != nil {
			return hashSHA256Response{v, err.Error()}, nil
		}
		return hashSHA256Response{v, ""}, nil
	}
}
