package main

import (
	"log"
	"net/http"

	"github.com/marcolino/jukebox/gen/openapi"
	"github.com/marcolino/jukebox/internal/api"
)

func main() {
	handler := api.Handler{}

	server, err := openapi.NewServer(&handler)
	if err != nil {
		log.Fatal(err)
	}

	if err = http.ListenAndServe(":9090", server); err != nil {
		log.Fatal(err)
	}

}
