package models

type Station struct {
	ID   int    `json:"id"`
	Kota string `json:"kota"`
	Name string `json:"name"`
}

type StationResponse struct {
	ID   int    `json:"id"`
	Kota string `json:"kota"`
	Name string `json:"name"`
}

func (Station) TableName() string {
	return "station"
}

func (StationResponse) TableName() string {
	return "station"
}
