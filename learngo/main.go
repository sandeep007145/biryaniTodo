package main

import (
	"log"
	"net/http"

	"./another"
	"./route"
)

func usemax() http.Handler {
	return route.Routes()
}

func main() {
	another.Chota()
	// mongodbconnect.Adapter()
	log.Fatal(http.ListenAndServe(":5555", usemax()))
}
