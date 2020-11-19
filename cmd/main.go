package main

import (
	"log"
	"mmt/go-search/pkg/api"
)

func main() {
	log.Fatal(api.Serve(8080))
}
