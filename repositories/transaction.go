package repositories

import (
	"server/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetTicketById(ID int) (models.Ticket, error)
	GetTicketTransaction(UserID int) ([]models.Transaction, error)
	CreateTransaction(trans models.Transaction) (models.Transaction, error)
	GetTransactionById(ID int) (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	GetTransUser(UserID int) (models.Transaction, error)
	Payment(payment models.Transaction) (models.Transaction, error)
	GetPaymentByIdTrans(ID int) (models.Transaction, error)
	UpdateTransaction(status string, ID string) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("Ticket.StartStation").Preload("Ticket.DestinationSTation").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTicketById(ID int) (models.Ticket, error) {
	var tiket models.Ticket
	err := r.db.Preload("StartStation").Preload("DestinationStation").First(&tiket, "id = ?", ID).Error

	return tiket, err
}

func (r *repository) CreateTransaction(trans models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Ticket.StartStation").Preload("Ticket.DestinationStation").Create(&trans).Error

	return trans, err
}

func (r *repository) GetTransactionById(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Ticket").Preload("User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Ticket.StartStation").Preload("Ticket.DestinationStation").First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) GetTicketTransaction(UserID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("Ticket.StartStation").Preload("Ticket.DestinationStation").Where("user_id = ?", UserID).Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransUser(UserID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Ticket.StartStation").Preload("Ticket.DestinationStation").Where("user = ?", UserID).Where("status = ?", "pending").Find(&transaction).Error

	return transaction, err
}

func (r *repository) Payment(payment models.Transaction) (models.Transaction, error) {
	err := r.db.Save(&payment).Error

	return payment, err
}

func (r *repository) GetPaymentByIdTrans(ID int) (models.Transaction, error) {
	var payment models.Transaction
	err := r.db.Preload("Ticket").Where("transaction_id = ?", ID).Find(&payment).Error

	return payment, err
}

func (r *repository) UpdateTransaction(status string, ID string) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("Ticket").First(&transaction, ID)

	if status != transaction.Status && status == "success" {
		var ticket models.Ticket
		r.db.First(&ticket, transaction.Ticket.ID)
		ticket.Qty = ticket.Qty - 1
		r.db.Save(&ticket)
	}

	transaction.Status = status
	error := r.db.Save(&transaction).Error
	return transaction, error
}
