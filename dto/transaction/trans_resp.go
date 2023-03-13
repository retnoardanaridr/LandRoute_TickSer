package transactiondto

import (
	"server/models"
	"time"
)

type TransactionResponse struct {
	ID        int                   `json:"id" gorm:"primary_key:auto_increment"`
	UserID    int                   `json:"user_id"`
	User      models.UserResponse   `json:"user"`
	TicketID  int                   `json:"ticket_id"`
	Ticket    models.TicketResponse `json:"ticket"`
	Status    string                `json:"status"`
	CreatedAt time.Time             `json:"-"`
	UpdatedAt time.Time             `json:"-"`
}
