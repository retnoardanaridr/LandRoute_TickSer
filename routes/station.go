package routes

import (
	"server/handlers"
	"server/pkg/mysql"
	"server/repositories"

	"github.com/gorilla/mux"
)

func StationRoutes(r *mux.Router) {
	StationRepository := repositories.RepositoryStation(mysql.DB)
	h := handlers.HandlerStation(StationRepository)

	r.HandleFunc("/stations", h.FindStations).Methods("GET")
	r.HandleFunc("/station/{id}", h.GetStationById).Methods("GET")
	r.HandleFunc("/station", h.CreateStation).Methods("POST")
}
