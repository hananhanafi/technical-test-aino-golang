package main

import (
	"github.com/hananhanafi/technical-test-aino-golang/database"
	"github.com/hananhanafi/technical-test-aino-golang/routes"
)

func main() {
	// initialize database
	database.Init()
	// initialize routes
	e := routes.Init()

	// run server
	e.Logger.Fatal(e.Start(":"))
}
