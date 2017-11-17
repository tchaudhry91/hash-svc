package main

import (
	"flag"
	kitlog "github.com/go-kit/kit/log"
	hashservice "github.com/tchaudhry91/hash-svc/pkg"
	"log"
	"net/http"
	"os"
)

func main() {
	var serverAddr string
	flag.StringVar(
		&serverAddr,
		"serverAddr",
		":8080",
		"Fully qualified server address like 0.0.0.0:8080",
	)
	flag.Parse()

	logger := kitlog.NewJSONLogger(os.Stderr)

	svc := hashservice.NewHashService()
	svc = hashservice.NewLoggingMiddleware(logger, svc)
	endpoint := hashservice.MakeHashSHA256Endpoint(svc)
	transportHandler := hashservice.MakeHashSHA256Handler(endpoint)

	http.Handle("/", transportHandler)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
