package main

import (
	"log"
	"net/http"

	httpapi "public-hiking-route-station/internal/http"
	"public-hiking-route-station/internal/seed"
	"public-hiking-route-station/internal/store"
)

func main() {
	s := store.New(seed.DefaultRecords())
	log.Fatal(http.ListenAndServe(":8080", httpapi.Router(s)))
}
