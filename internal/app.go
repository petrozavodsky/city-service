package internal

import (
	city "city_service/pkg"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func WebService(storage *city.Storage, port int) *http.Server {

	server := &http.Server{Addr: "localhost:" + strconv.Itoa(port), Handler: handlers(storage)}

	return server
}

func handlers(storage *city.Storage) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/city", AddCity(storage))

	router.Get("/city/{id}", GetCity(storage))

	router.Delete("/city/{id}", DeleteCity(storage))

	router.Put("/city/{id}", UpdatePopulation(storage))

	router.Get("/cities-by-region/{key}", GetCitiesByRegion(storage))

	router.Get("/cities-by-district/{key}", GetCitiesByDistrict(storage))

	router.Get("/cities-by-population", GetCitiesByPopulation(storage))

	router.Get("/cities-by-foundation", GetCitiesByFoundation(storage))

	return router
}
