package main

import (
	"technical-test-aino-golang/database"
	"technical-test-aino-golang/routes"
)

func main() {
	// initialize database
	database.Init()
	// initialize routes
	e := routes.Init()

	// run server
	e.Logger.Fatal(e.Start(":1234"))
}
