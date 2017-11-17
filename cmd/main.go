package main

import (
	"flag"
	hashservice "github.com/tchaudhry91/hash-svc/pkg"
	"log"
	"net/http"
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

	svc := hashservice.NewHashService()
	endpoint := hashservice.MakeHashSHA256Endpoint(svc)
	transportHandler := hashservice.MakeHashSHA256Handler(endpoint)

	http.Handle("/", transportHandler)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
