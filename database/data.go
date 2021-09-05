package database

import (
	"context"
	"covid-information-update/config"
	"covid-information-update/models"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func FetchDataFromApi() (interface{}, error) {
	req, err := http.NewRequest("GET", "https://api.kawalcorona.com/indonesia", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	result := []models.DataIndonesia{}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil

}

func GetDataToday() (interface{}, error) {
	var ctx = context.Background()
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}
	csr, err := db.Collection("indonesia").Find(ctx, bson.M{"name": "Indonesia"})
	if err != nil {
		return nil, err
	}
	defer csr.Close(ctx)
	result := make([]models.DataIndonesia, 0)
	for csr.Next(ctx) {
		var row models.DataIndonesia
		err := csr.Decode(&row)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	count := len(result)
	return result[count-1], nil
}

func GetDataLog() (interface{}, error) {
	var ctx = context.Background()
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}
	csr, err := db.Collection("indonesia").Find(ctx, bson.M{"name": "Indonesia"})
	if err != nil {
		return nil, err
	}
	defer csr.Close(ctx)
	result := make([]models.DataIndonesia, 0)
	for csr.Next(ctx) {
		var row models.DataIndonesia
		err := csr.Decode(&row)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result[0].Daerah, nil
}

func PostData() (interface{}, error) {
	var ctx = context.Background()
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}
	newData, err := FetchDataFromApi()
	if err != nil {
		return nil, err
	}
	newDataCasting := newData.([]models.DataIndonesia)

	dataIndonesia := models.DataIndonesia{
		Update_time: time.Now().Add(7 * time.Hour),
		Daerah:      "Indonesia",
		Positif:     newDataCasting[0].Positif,
		Sembuh:      newDataCasting[0].Sembuh,
		Meninggal:   newDataCasting[0].Meninggal,
		Dirawat:     newDataCasting[0].Dirawat,
	}
	_, err = db.Collection("indonesia").InsertOne(ctx, dataIndonesia)
	if err != nil {
		return nil, err
	}
	return dataIndonesia, nil
}
