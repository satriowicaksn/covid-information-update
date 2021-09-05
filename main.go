package main

import "covid-information-update/routes"

func main() {
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
