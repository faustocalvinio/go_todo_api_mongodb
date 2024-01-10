package routes

import (
	"net/http"
	"todo_api_mongodb/controllers"

	"github.com/go-chi/chi/v5"
)
func Router() http.Handler {
	router := chi.NewRouter()
	router.Get("/",controllers.FetchTodos )
	return router
}