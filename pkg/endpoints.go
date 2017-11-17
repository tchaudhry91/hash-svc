package hasher

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type hashSHA256Request struct {
	S string `json:"s"`
}

type hashSHA256Response struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

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
