package app

import (
	"city_service/internal/handlers"
	"github.com/go-chi/chi"
)

func WebService(port int) {

	router := chi.NewRouter()

	router.Post("/city", handlers.AddCity())

	router.Get("/city{id}", handlers.GetCity())

	router.Delete("/city", handlers.DeleteCity())

}
