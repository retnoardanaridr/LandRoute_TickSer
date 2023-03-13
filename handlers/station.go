package handlers

import (
	"encoding/json"
	"net/http"
	dto "server/dto/result"
	stationdto "server/dto/station"
	"server/models"
	"server/repositories"
	"strconv"

	"github.com/gorilla/mux"
)

type handlerStation struct {
	StationRepository repositories.StationRepository
}

func HandlerStation(stationRepository repositories.StationRepository) *handlerStation {
	return &handlerStation{stationRepository}
}

func (h *handlerStation) FindStations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	station, err := h.StationRepository.FindStations()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: station}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerStation) GetStationById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	station, err := h.StationRepository.GetStationById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: station}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerStation) CreateStation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(stationdto.StationRequest)
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	Field := models.Station{
		Kota: request.Kota,
		Name: request.Name,
	}

	station, err := h.StationRepository.CreateStation(Field)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: station}
	json.NewEncoder(w).Encode(response)
}
