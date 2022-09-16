package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danmory/geocoding-service/data-service/internal/core"
	"github.com/danmory/geocoding-service/data-service/internal/service"
)

func getCoordinatesFromQuery(req *http.Request) (float32, float32, error) {
	lat, err := strconv.ParseFloat(req.URL.Query().Get("lat"), 32)
	if err != nil {
		return 0, 0, err
	}
	lon, err := strconv.ParseFloat(req.URL.Query().Get("lon"), 32)
	if err != nil {
		return 0, 0, err
	}
	return float32(lat), float32(lon), nil
}

func writePlaceToResponse(rw http.ResponseWriter, place *core.Place) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(place)
}

func SearchByCoordinates(rw http.ResponseWriter, req *http.Request) {
	var place *core.Place
	lat, lon, err := getCoordinatesFromQuery(req)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	coords := core.PlaceCoordinates{Lat: lat, Lon: lon}
	place = service.CheckCoordinatesCache(coords)
	if place != nil {
		writePlaceToResponse(rw, place)
		return
	}
	place = service.SearchByCoordinates(coords)
	if place != nil {
		writePlaceToResponse(rw, place)
		return
	}
	rw.WriteHeader(http.StatusNotFound)
}

func SearchByName(rw http.ResponseWriter, req *http.Request) {
	var place *core.Place
	name := req.URL.Query().Get("name")
	if name == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	placeName := core.PlaceName{Name: name}
	place = service.CheckNameCache(placeName)
	if place != nil {
		writePlaceToResponse(rw, place)
		return
	}
	place = service.SearchByName(placeName)
	if place != nil {
		writePlaceToResponse(rw, place)
		return
	}
	rw.WriteHeader(http.StatusNotFound)
}
