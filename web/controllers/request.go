package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/learngolangwithkushan/service"
)

func (app *Application) RequestHandler(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		TransactionId string
		Success       bool
		Response      bool
	}{
		TransactionId: "",
		Success:       false,
		Response:      false,
	}
	if r.FormValue("submitted") == "true" {
		helloValue := r.FormValue("hello")
		log.Println(helloValue)

		responseFromService, err := service.ReadWebPage(helloValue)
		if err != nil {
			log.Println(err, "Error happen ")
		}

		data.TransactionId = responseFromService.Title + " : " + responseFromService.HtmlVersion + " : " + responseFromService.HeadLines +
			" : " + strconv.Itoa(responseFromService.NumberOfLinks)
		data.Success = true
		data.Response = true
	}
	renderTemplate(w, r, "request.html", data)
}
