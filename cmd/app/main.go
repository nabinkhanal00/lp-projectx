package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nabinkhanal00/lp-calculator/controllers/calculate"
)

func main() {
	// returns a type that implements http.Handler interface
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// use rate limiter to prevent DOS attacks
	r.Use(httprate.LimitByIP(100, time.Minute))

	// handler to calculate the result
	r.Post("/calculate", calculate.CalculateController)

	// handler to serve the built frontend files
	r.Handle("/*", http.FileServer(http.Dir("./web")))

	// run the server
	port := fmt.Sprintf(":%v", getenv("SERVER_PORT", "3000"))
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
