package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)
// Server — структура, которая хранит логгер и сам сервер
type Server struct {
	Logger *log.Logger
	Server *http.Server
}
// NewServer создает и настраивает новый сервер
func NewServer(logger *log.Logger) *Server {
    // Создаем маршрутизатор
	mux := http.NewServeMux()

	// Регистрируем маршруты
	mux.HandleFunc("/", handlers.IndexHandler) // Главная страница
	mux.HandleFunc("/upload", handlers.UploadHandler) // Страница загрузки

	// Настраиваем http.Server по ТЗ
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
    // Возвращаем готовую структуру
	return &Server{
		Logger: logger,
		Server: httpServer,
	}
}