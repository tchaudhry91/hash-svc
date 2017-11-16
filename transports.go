package main

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	svc := hashService{}

	hashEndpoint := makeHashSHA256Endpoint(svc)
	muxRouter := makeHashSHA256Handler(hashEndpoint)

	http.Handle("/", muxRouter)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func makeHashSHA256Handler(e endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()

	hashSHA256Handler := httptransport.NewServer(
		e,
		decodeHashSHA256Request,
		encodeHashSHA256Response,
	)

	r.Methods("POST").Path("/hash").Handler(hashSHA256Handler)
	return r
}

func decodeHashSHA256Request(_ context.Context, r *http.Request) (interface{}, error) {
	var request hashSHA256Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeHashSHA256Response(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	return json.NewEncoder(w).Encode(resp)
}
