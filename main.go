package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"main/apis"
	"net/http"
)

func main() {

	r := chi.NewRouter()

	// Middleware для логгирования и восстановления
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Маршрут для получения списка заметок
	r.Get("/notes", apis.GetNotes)

	// Маршрут для добавления новой заметки
	r.Post("/notes", apis.CreateNote)

	// Запускаем сервер
	http.ListenAndServe(":8080", r)
}
