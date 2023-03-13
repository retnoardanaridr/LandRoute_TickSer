package models

import "time"

type Transaction struct {
	ID            int            `json:"id"`
	TransactionID int            `json:"transaction_id"`
	UserID        int            `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User          UserResponse   `json:"user"`
	TicketID      int            `json:"ticket_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Ticket        TicketResponse `json:"ticket"`
	Total         int            `json:"total" form:"total"`
	Qty           int            `json:"qty" form:"qty"`
	Status        string         `json:"status" gorm:"type: varchar(255)"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
}
