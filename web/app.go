package web

import (
	"fmt"
	"net/http"

	"github.com/learngolangwithkushan/web/controllers"
)

func Serve(app *controllers.Application) {
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/request.html", app.RequestHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/request.html", http.StatusTemporaryRedirect)
	})

	fmt.Println("Listening (http://localhost:30000/) ...")
	http.ListenAndServe(":30000", nil)
}
