package main

import (
	"covid-information-update/routes"
	"log"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT MUST BE SET")
	}

	e := routes.New()
	e.Logger.Fatal(e.Start(":" + port))
}
