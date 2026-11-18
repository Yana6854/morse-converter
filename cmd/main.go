package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер, он будет писать сообщения с датой и временем в консоль
	logger := log.New(os.Stdout, "SERVER: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Вызываем функцию которая собирает сервер
	srv := server.NewServer(logger)

	// Запускаем сервер
	err := srv.Server.ListenAndServe()
	if err != nil {
		logger.Fatalf("Ошибка при запуске сервера: %s", err.Error())
	}
}
