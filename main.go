package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"main/apis"
	"main/logger"
	"net/http"
)

func main() {
	//Настройка логера
	logger.Setup()
	//Закрытие логирования в файл после закрытия основной горутины
	defer logger.File.Close()

	r := chi.NewRouter()
	r.Use(logger.LoggerMiddleware)
	r.Use(middleware.Recoverer)
	// Маршрут для получения списка заметок
	r.Get("/notes", apis.GetNotes)
	// Маршрут для добавления новой заметки
	r.Post("/notes", apis.CreateNote)
	http.ListenAndServe(":8080", r)
}
