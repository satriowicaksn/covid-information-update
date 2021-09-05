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

func FetchDataFromApi() ([]models.DataIndonesia, error) {
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

func GetDataNow() (models.DataIndonesia, error) {
	var ctx = context.Background()
	result := []models.DataIndonesia{}
	db, err := config.Connect()
	if err != nil {
		return result[0], err
	}
	csr, err := db.Collection("indonesia").Find(ctx, bson.M{"daerah": "Indonesia"})
	if err != nil {
		return result[0], err
	}
	defer csr.Close(ctx)

	for csr.Next(ctx) {
		var row models.DataIndonesia
		err := csr.Decode(&row)
		if err != nil {
			return result[0], err
		}
		result = append(result, row)
	}
	count := len(result)
	return result[count-1], nil
}

func CheckDataUpdate() (bool, error) {
	fromApi, err := FetchDataFromApi()
	if err != nil {
		return false, err
	}
	fromDatabase, err := GetDataNow()
	if err != nil {
		return false, err
	}
	if fromApi[0].Positif != fromDatabase.Positif {
		return true, nil
	} else if fromApi[0].Sembuh != fromDatabase.Sembuh {
		return true, nil
	} else if fromApi[0].Meninggal != fromDatabase.Meninggal {
		return true, nil
	} else if fromApi[0].Dirawat != fromDatabase.Dirawat {
		return true, nil
	}
	return false, nil
}

func GetDataLog() (interface{}, error) {
	var ctx = context.Background()
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}
	csr, err := db.Collection("indonesia").Find(ctx, bson.M{"daerah": "Indonesia"})
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
	return result, nil
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

	dataIndonesia := models.DataIndonesia{
		Update_time: time.Now().Add(7 * time.Hour),
		Daerah:      "Indonesia",
		Positif:     newData[0].Positif,
		Sembuh:      newData[0].Sembuh,
		Meninggal:   newData[0].Meninggal,
		Dirawat:     newData[0].Dirawat,
	}
	_, err = db.Collection("indonesia").InsertOne(ctx, dataIndonesia)
	if err != nil {
		return nil, err
	}
	return dataIndonesia, nil
}
