package main

import (
	"fmt"

	"github.com/learngolangwithkushan/web"
	"github.com/learngolangwithkushan/web/controllers"
)

func main() {
	fmt.Print("Start 1")
	app := &controllers.Application{}
	web.Serve(app)
}
