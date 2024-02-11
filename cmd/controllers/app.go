package controllers

import (
	"fmt"
	"net/http"
)

func Serve(app *Application) {
	fs := http.FileServer(http.Dir("templates/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/request.gohtml", app.RequestHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/request.gohtml", http.StatusTemporaryRedirect)
	})

	// make port env variable
	fmt.Println("Listening (http://localhost:30000/) ...")
	http.ListenAndServe(":30000", nil)
}
