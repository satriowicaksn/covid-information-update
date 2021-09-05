package models

import "time"

type DataIndonesia struct {
	Update_time time.Time `json:"update_time" bson:"update_time"`
	Daerah      string    `json:"daerah" bson:"daerah"`
	Positif     string    `json:"positif" bson:"positif"`
	Sembuh      string    `json:"sembuh" bson:"sembuh"`
	Meninggal   string    `json:"meninggal" bson:"meninggal"`
	Dirawat     string    `json:"dirawat" bson:"dirawat"`
}
