package repositories

import (
	"server/models"

	"gorm.io/gorm"
)

type StationRepository interface {
	FindStations() ([]models.Station, error)
	GetStationById(ID int) (models.Station, error)
	CreateStation(station models.Station) (models.Station, error)
}

func RepositoryStation(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindStations() ([]models.Station, error) {
	var station []models.Station
	err := r.db.Find(&station).Error

	return station, err
}

func (r *repository) GetStationById(ID int) (models.Station, error) {
	var stationId models.Station
	err := r.db.First(&stationId, ID).Error

	return stationId, err
}

func (r *repository) CreateStation(station models.Station) (models.Station, error) {
	err := r.db.Create(&station).Error
	return station, err
}
