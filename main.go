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
	if len(port) == 0 {
		port = "8081"
	}
	// run server
	e.Logger.Fatal(e.Start(":" + port))
}
