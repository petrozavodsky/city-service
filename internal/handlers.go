package internal

import (
	city "city_service/pkg"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"log"
	"net/http"
	"strconv"
)

// AddCity хендлер добавления гродоа
func AddCity(storage *city.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		defer func(Body io.ReadCloser) {

			err := Body.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}(r.Body)

		w.Header().Set("Content-Type", "application/json")

		requestCity := city.City{}
		if err := json.NewDecoder(r.Body).Decode(&requestCity); err != nil {
			log.Fatalln(err)
		}

		requestCity.SetId(storage.GenerateId())

		storage.AddCity(requestCity.GetId(), &requestCity)

		response := fmt.Sprintf("Город с ID %d создан\n", requestCity.GetId())

		body := MakeBody()
		w.WriteHeader(http.StatusCreated)
		body.SetMessage(response)

		if err := json.NewEncoder(w).Encode(body); err != nil {
			log.Fatalln(err)
		}
		return
	}

}

// GetCity хендлер получения грода по id
func GetCity(storage *city.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			body := MakeBody()
			w.WriteHeader(http.StatusBadRequest)
			body.SetMessage("Идентификатор не корректен")

			if err := json.NewEncoder(w).Encode(body); err != nil {
				log.Fatalln(err)
			}
			return
		}

		getCity := storage.GetCity(id)

		if getCity == nil {
			body := MakeBody()
			w.WriteHeader(http.StatusNotFound)
			body.SetMessage("Город не найден")

			if err := json.NewEncoder(w).Encode(body); err != nil {
				log.Fatalln(err)
			}
			return
		}

		if err := json.NewEncoder(w).Encode(getCity); err != nil {
			log.Fatalln(err)
		}
	}
}

// DeleteCity хендлер удаления грода по id
func DeleteCity(storage *city.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		defer func(Body io.ReadCloser) {

			err := Body.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}(r.Body)

		w.Header().Set("Content-Type", "application/json")

		request := map[string]string{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Fatalln(err)
		}

		id, err := strconv.Atoi(request["id"])
		if err != nil {
			log.Fatalln(err)
		}

		response := fmt.Sprintf("Город %s удален\n")
		storage.DeleteCity(id)

		body := MakeBody()
		w.WriteHeader(http.StatusOK)
		body.SetMessage(response)

		if err := json.NewEncoder(w).Encode(body); err != nil {
			log.Fatalln(err)
		}
		return
	}

}

// UpdatePopulation хендлер обновления численности
func UpdatePopulation(storage *city.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		defer func(Body io.ReadCloser) {

			err := Body.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}(r.Body)

		w.Header().Set("Content-Type", "application/json")

		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			body := MakeBody()
			w.WriteHeader(http.StatusBadRequest)
			body.SetMessage("Идентификатор не корректен")

			if err := json.NewEncoder(w).Encode(body); err != nil {
				log.Fatalln(err)
			}
			return
		}

		request := map[string]int{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Fatalln(err)
		}

		population := request["population"]
		if err != nil {
			body := MakeBody()
			w.WriteHeader(http.StatusBadRequest)
			body.SetMessage("Численность не корректена")

			if err := json.NewEncoder(w).Encode(body); err != nil {
				log.Fatalln(err)
			}
			return
		}

		requestCity := storage.GetCity(id)

		requestCity.SetPopulation(population)

		storage.UpdatePopulation(requestCity.GetId(), requestCity)

		response := fmt.Sprintf(
			"Город %s с ID %d обновлен (конкретнее численность населения)\n",
			requestCity.GetName(),
			requestCity.GetId(),
		)

		body := MakeBody()
		w.WriteHeader(http.StatusOK)
		body.SetMessage(response)

		if err := json.NewEncoder(w).Encode(body); err != nil {
			log.Fatalln(err)
		}
		return
	}

}

// GetCitiesByRegion хендлер получения городов по региону
func GetCitiesByRegion(storage *city.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		key := chi.URLParam(r, "key")

		getCity := storage.GetByRegion(key)

		if getCity == nil {
			body := MakeBody()
			w.WriteHeader(http.StatusNotFound)
			body.SetMessage("Регион не найден")

			if err := json.NewEncoder(w).Encode(body); err != nil {
				log.Fatalln(err)
			}
			return
		}

		if err := json.NewEncoder(w).Encode(getCity); err != nil {
			log.Fatalln(err)
		}
	}
}

// GetCitiesByDistrict хендлер получения городов по району
func GetCitiesByDistrict(storage *city.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		key := chi.URLParam(r, "key")

		getCity := storage.GetByDistrict(key)

		if getCity == nil {
			body := MakeBody()
			w.WriteHeader(http.StatusNotFound)
			body.SetMessage("Район не найден")

			if err := json.NewEncoder(w).Encode(body); err != nil {
				log.Fatalln(err)
			}
			return
		}

		if err := json.NewEncoder(w).Encode(getCity); err != nil {
			log.Fatalln(err)
		}
	}
}

// GetCitiesByPopulation хендлен получения городов по числености населения
func GetCitiesByPopulation(storage *city.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		fromStr := r.URL.Query().Get("from")

		if "" == fromStr {
			body := MakeBody()
			w.WriteHeader(http.StatusBadRequest)
			body.SetMessage("Идентификатор диапазова не корректен")

			if err := json.NewEncoder(w).Encode(body); err != nil {
				log.Println(err)
			}
			return
		}

		from, err := strconv.Atoi(fromStr)

		if err != nil {
			log.Fatalln(err)
		}

		toStr := r.URL.Query().Get("to")

		if "" == toStr {
			body := MakeBody()
			w.WriteHeader(http.StatusBadRequest)
			body.SetMessage("Идентификатор диапазова не корректен")

			if err := json.NewEncoder(w).Encode(body); err != nil {
				log.Println(err)
			}
			return
		}

		to, err := strconv.Atoi(toStr)

		if err != nil {
			log.Fatalln(err)
		}

		getCities := storage.GetByPopulation(from, to)

		if from >= to {
			body := MakeBody()
			w.WriteHeader(http.StatusBadRequest)
			body.SetMessage("Интервал задание неверно")

			if err := json.NewEncoder(w).Encode(body); err != nil {
				log.Fatalln(err)
			}
			return
		}

		if getCities == nil {
			body := MakeBody()
			w.WriteHeader(http.StatusNotFound)
			body.SetMessage("Интервал пуст")

			if err := json.NewEncoder(w).Encode(body); err != nil {
				log.Fatalln(err)
			}
			return
		}

		if err := json.NewEncoder(w).Encode(getCities); err != nil {
			log.Fatalln(err)
		}

	}

}

// GetCitiesByFoundation хендлен получения городов по году основания
func GetCitiesByFoundation(storage *city.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		fromStr := r.URL.Query().Get("from")

		if "" == fromStr {
			body := MakeBody()
			w.WriteHeader(http.StatusBadRequest)
			body.SetMessage("Идентификатор диапазова не корректен")

			if err := json.NewEncoder(w).Encode(body); err != nil {
				log.Println(err)
			}
			return
		}

		from, err := strconv.Atoi(fromStr)

		if err != nil {
			log.Fatalln(err)
		}

		toStr := r.URL.Query().Get("to")

		if "" == toStr {
			body := MakeBody()
			w.WriteHeader(http.StatusBadRequest)
			body.SetMessage("Идентификатор диапазова не корректен")

			if err := json.NewEncoder(w).Encode(body); err != nil {
				log.Println(err)
			}
			return
		}

		to, err := strconv.Atoi(toStr)

		if err != nil {
			log.Fatalln(err)
		}

		getCities := storage.GetByFoundation(from, to)

		if from >= to {
			body := MakeBody()
			w.WriteHeader(http.StatusBadRequest)
			body.SetMessage("Интервал задание неверно")

			if err := json.NewEncoder(w).Encode(body); err != nil {
				log.Fatalln(err)
			}
			return
		}

		if getCities == nil {
			body := MakeBody()
			w.WriteHeader(http.StatusNotFound)
			body.SetMessage("Интервал пуст")

			if err := json.NewEncoder(w).Encode(body); err != nil {
				log.Fatalln(err)
			}
			return
		}

		if err := json.NewEncoder(w).Encode(getCities); err != nil {
			log.Fatalln(err)
		}

	}
}
