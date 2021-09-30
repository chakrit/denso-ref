package main

import (
	"dx/controllers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	home := controllers.HomeController{}
	r.Get("/", home.Index)
	r.Post("/", home.Create)

	r.Get("/new", home.New)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalln("error", err)
	}
}
