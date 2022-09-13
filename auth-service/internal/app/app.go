package app

import (
	"log"
	"net/http"
	"os"

	"github.com/danmory/geocoding-service/auth-service/internal/storages/psql"
	"github.com/danmory/geocoding-service/auth-service/internal/transport/http/handlers"
	"github.com/julienschmidt/httprouter"
)

func RunApp() {
	psql.GetDatabase()
	router := httprouter.New()
	router.POST("/login", handlers.HandleLogin)
	router.POST("/register", handlers.HandleRegister)
	log.Fatal(http.ListenAndServe(os.Getenv("APP_ADDRESS"), router))
}
