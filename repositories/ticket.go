package repositories

import (
	"server/models"

	"gorm.io/gorm"
)

type TicketRepository interface {
	FindTicket() ([]models.Ticket, error)
	FindStationById(StationID int) (models.Station, error)
	// GetTicketById(ID int) (models.Ticket, error) //nggk dpkek
	AddTicket(ticket models.Ticket) (models.Ticket, error)
	GetTicket(ID int) (models.Ticket, error)

	FilterStation(stationID int) (models.Station, error)
	FilterTicket(start int, destination int, date string) ([]models.Ticket, error)

	CreateTransactionQty(transaction models.Transaction) (models.Transaction, error)
}

func RepositoryTicket(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTicket() ([]models.Ticket, error) {
	var ticket []models.Ticket
	err := r.db.Preload("StartStation").Preload("DestinationStation").Find(&ticket).Error

	return ticket, err
}

func (r *repository) FindStationById(station int) (models.Station, error) {
	var stationID models.Station
	err := r.db.Find(&stationID, stationID).Error

	return stationID, err
}

// nggk dpkek
// func (r *repository) GetTicketById(ID int) (models.Ticket, error) {
// 	var getTicketID models.Ticket
// 	err := r.db.Preload("Station").First(&getTicketID, ID).Error

// 	return getTicketID, err
// }

func (r *repository) AddTicket(ticket models.Ticket) (models.Ticket, error) {
	err := r.db.Save(&ticket).Error

	return ticket, err
}

func (r *repository) GetTicket(ID int) (models.Ticket, error) {
	var ticket models.Ticket
	err := r.db.Preload("StartStation").Preload("DestinationStation").First(&ticket, "id = ?", ID).Error

	return ticket, err
}

func (r *repository) FilterStation(statiun int) (models.Station, error) {
	var station models.Station
	err := r.db.Where("station = ?", station.ID).First(&station).Error

	return station, err
}

func (r *repository) FilterTicket(start int, destination int, date string) ([]models.Ticket, error) {
	var ticket []models.Ticket
	err := r.db.Preload("StartStation").Preload("DestinationStation").Where("start_destination_id = ? AND destination_station_id = ? AND start_date = ?", start, destination, date).Find(&ticket).Error

	return ticket, err
}

func (r *repository) CreateTransactionQty(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error
	return transaction, err
}
