package main

import (
	"city_service/internal"
	city "city_service/pkg"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var port int
	var fileName string

	flag.StringVar(&fileName, "file", "cities.csv", "File")
	flag.IntVar(&port, "port", 8080, "Port")
	flag.Parse()

	//подготовка хранилища
	storage := city.MakeStorage()

	// чтение данных
	internal.FileReader(storage, fileName)

	// подготовка сервера
	server := internal.WebService(storage, port)

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Ожидание сигнала прерывания процесса
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Шатдаун с таймаутом до завершение
		shutdownCtx, _ := context.WithTimeout(serverCtx, 4*time.Second)

		go func() {

			<-shutdownCtx.Done()
			//Таймаут вышел, принудительное завершение
			if shutdownCtx.Err() == context.DeadlineExceeded {

				// Сохранение состояний
				internal.FileWriter(storage, fileName)
				log.Fatal("Завершение по таймауту")
			}
		}()

		// Запус graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}

		// Сохранение состояний
		internal.FileWriter(storage, fileName)
		serverStopCtx()
	}()

	//Стерт сервера
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Ожижание остановки сервера
	<-serverCtx.Done()

}
