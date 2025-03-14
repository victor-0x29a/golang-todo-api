package main

import (
	"fmt"
	"golang-api/configs"
	"golang-api/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/", handlers.Create)
	r.Get("/{id}", handlers.Get)
	r.Get("/list", handlers.List)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)

	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.GetAPI().Port), r)

	if err != nil {
		panic(err)
	}
}
