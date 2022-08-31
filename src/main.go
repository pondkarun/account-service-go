package main

import (
	"account-service/Routers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func init() {
}

func main() {
	r := chi.NewRouter()
	port := "3000"

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("404"))
	})

	r.Mount("/account", Routers.AccountRouter())
	fmt.Println("Run Port : " + port)

	http.ListenAndServe(":"+port, r)

}
