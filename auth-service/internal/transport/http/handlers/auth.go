package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/danmory/geocoding-service/auth-service/internal/core"
	"github.com/danmory/geocoding-service/auth-service/internal/service"
	"github.com/julienschmidt/httprouter"
)

func HandleRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var user core.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if token, err := service.RegisterUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.Write([]byte(token))
		w.WriteHeader(http.StatusCreated)
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var user core.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if token, err := service.LoginUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.Write([]byte(token))
		w.WriteHeader(http.StatusOK)
	}
}
