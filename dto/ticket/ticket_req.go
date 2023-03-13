package ticketdto

type TicketRequest struct {
	NameTrain            string `json:"name_train" gorm:"type: varchar(255)"`
	TypeTrain            string `json:"type_train" gorm:"type: varchat(255)"`
	StartDate            string `json:"start_date" gorm:"type: varchar(255)"`
	StartStationID       int    `json:"start_station_id,string,omitempty" form:"start_station_id"`
	StartTime            string `json:"start_time" gorm:"type: varchar(255)"`
	DestinationStationID int    `json:"destination_station_id,string,omitempty" form:"destination_station_id"`
	ArrivalTime          string `json:"arrival_time" gorm:"type: varchar(255)"`
	Price                int    `json:"price,string,omitempty" form:"price"`
	Qty                  int    `json:"qty,string,omitempty" form:"qty"`
}

type FilterRequest struct {
	StartDate            string `json:"start_date" form:"start_date"`
	StartStationID       int    `json:"start_station_id" form:"start_station_id"`
	DestinationStationID int    `json:"destination_station_id" form:"destination_station_id"`
}

type TransTicket struct {
	Qty int `json:"qty" form:"qty"`
}
