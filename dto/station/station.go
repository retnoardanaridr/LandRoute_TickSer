package stationdto

type StationRequest struct {
	Kota string `json:"kota" form:"kota" gorm:"type: varchar(255)"`
	Name string `json:"name" form:"name" gorm:"type: varchar(255)"`
}
