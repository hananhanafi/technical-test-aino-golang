package main

import (
	"os"

	"github.com/hananhanafi/technical-test-aino-golang/database"
	"github.com/hananhanafi/technical-test-aino-golang/routes"
)

func main() {
	// initialize database
	database.Init()
	// initialize routes
	e := routes.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}
	// run server
	e.Logger.Fatal(e.Start(":" + port))
}
