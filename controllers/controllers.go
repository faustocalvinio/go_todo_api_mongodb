package controllers

import (
	"fmt"
	"net/http"
	"github.com/thedevsaddam/renderer"
)
var rnd *renderer.Render
func FirstController() {
	fmt.Println("Controller works")
}

func SecondController(w http.ResponseWriter, r *http.Request) {
	rnd = renderer.New()
	fmt.Println("Request",r.Body)
	rnd.JSON(w, http.StatusAccepted, renderer.M{
		"data": "data",
	})
}
