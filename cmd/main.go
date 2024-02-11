package main

import (
	"log"

	"github.com/learngolangwithkushan/cmd/controllers"
)

func main() {
	log.Print("Starting the application")
	app := &controllers.Application{}
	controllers.Serve(app)
}
