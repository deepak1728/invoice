package main

import (
	"invoice/db"
	"invoice/routes"
	l "log"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	err := db.ConnectToMongo()

	if err != nil {
		log.Error().Err(err).Send()
		l.Panic(err)
	}

	l.Fatal(http.ListenAndServe(":8181", routes.CreateRouter()))
}
