package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/danmory/geocoding-service/data-service/internal/middlewares"
	"github.com/danmory/geocoding-service/data-service/internal/transport/http/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file " + err.Error())
	}
}

func Run() {
	r := getRouter()
	srv := &http.Server{
		Addr:    os.Getenv("APP_ADDRESS"),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(middlewares.Auth)
	r.HandleFunc("/place", handlers.SearchByCoordinates)
	r.HandleFunc("/coords", handlers.SearchByName)
	return r
}
