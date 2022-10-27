package internal

import (
	city "city_service/pkg"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

// FileReader чтение файла
func FileReader(storage *city.Storage, fileStr string) {

	file, err := os.Open(fileStr)

	//Обработка ошибки открытия файла
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()

		//Обработка ошибки закрытия файла
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for _, line := range records {

		id, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}

		population, err := strconv.Atoi(line[4])
		if err != nil {
			log.Fatal(err)
		}

		foundation, err := strconv.Atoi(line[5])
		if err != nil {
			log.Fatal(err)
		}

		cityObj := city.City{
			Id:         id,
			Name:       line[1],
			Region:     line[2],
			District:   line[3],
			Population: population,
			Foundation: foundation,
		}

		storage.AddCity(id, &cityObj)

	}

}

// FileWriter  Запись файла
func FileWriter(storage *city.Storage, fileStr string) {
	file, err := os.OpenFile(fileStr, os.O_RDWR|os.O_CREATE, os.ModePerm)

	//Обработка ошибки открытия файла
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		//Обработка ошибки закрытия файла
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	data := storage.StoreToStrings()

	writer := csv.NewWriter(file)

	err = writer.WriteAll(data)

	//Обработка ошибки записи
	if err != nil {
		log.Fatalln(err)
	}
}
