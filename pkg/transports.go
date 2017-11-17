package hasher

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	//ErrJSONUnMarshall indicates a bad request where json unmarshalling failed
	ErrJSONUnMarshall = errors.New("failed to parse json")
)

// MakeHashSHA256Handler returns an http.Handler populated with go-kit endpoint routes
func MakeHashSHA256Handler(e endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
	}

	hashSHA256Handler := httptransport.NewServer(
		e,
		decodeHashSHA256Request,
		encodeHashSHA256Response,
		options...,
	)

	r.Methods("POST").Path("/hash").Handler(hashSHA256Handler)
	return r
}

func decodeHashSHA256Request(_ context.Context, r *http.Request) (interface{}, error) {
	var request hashSHA256Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, ErrJSONUnMarshall
	}
	return request, nil
}

func encodeHashSHA256Response(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	return json.NewEncoder(w).Encode(resp)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrJSONUnMarshall:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
