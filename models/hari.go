package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Data 	Data 	`json:"data"`
	Update	Update `json:"update"`
}

type Data struct {
	Id   				int   `json:"id"`
	Odp     			int    `json:"jumlah_odp"`
	Pdp   				int    `json:"jumlah_pdp"`
	Spesimen  			int    `json:"total_spesimen"`
	SpesimenNegatif  	int    `json:"total_spesimen_negatif"`
}

type Update struct {
	Total	Total	`json:"total"`
	Penambahan	Penambahan	`json:"penambahan"`
}

type Total struct {
	Positif   	int   `json:"jumlah_positif"`
	Sembuh     	int    `json:"jumlah_sembuh"`
	Meninggal   int    `json:"jumlah_meninggal"`
	Dirawat  	int   `json:"jumlah_dirawat"`
}

type Penambahan struct {
	Positif   	int   `json:"jumlah_positif"`
	Sembuh     	int    `json:"jumlah_sembuh"`
	Meninggal   int    `json:"jumlah_meninggal"`
	Dirawat  	int    `json:"jumlah_dirawat"`
	Tanggal  	string    `json:"tanggal"`
	Update  	string    `json:"updated"`
}

func (r *Response) GetData() string {
	url := "https://data.covid19.go.id/public/api/update.json"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject Response
	json.Unmarshal(data, &responseObject)


	pesan := ""
	pesan =
		"*Data Kasus Covid-19 Indonesia*\n\n" +
		"Positif : " + strconv.Itoa(responseObject.Update.Total.Positif) + ",\n" +
		"Sembuh : " + strconv.Itoa(responseObject.Update.Total.Sembuh) + ",\n" +
		"Meninggal : " + strconv.Itoa(responseObject.Update.Total.Meninggal) + "\n\n" +

		"ODP : " + strconv.Itoa(responseObject.Data.Odp) + ",\n" +
		"PDP : " + strconv.Itoa(responseObject.Data.Pdp) + ",\n" +
		"Spesimen : " + strconv.Itoa(responseObject.Data.Spesimen) + ",\n" +
		"Spesimen Negatif : " + strconv.Itoa(responseObject.Data.SpesimenNegatif) + ",\n\n" +

		"*Data per hari ini* " +"\n\n" +
		"Positif : " + strconv.Itoa(responseObject.Update.Penambahan.Positif) + ",\n" +
		"Sembuh : " + strconv.Itoa(responseObject.Update.Penambahan.Sembuh) + ",\n" +
		"Meninggal : " + strconv.Itoa(responseObject.Update.Penambahan.Meninggal) + "\n\n" +

		"*Updated* : " + responseObject.Update.Penambahan.Update + "\n\n" +

		"#jagakesehatan"

	return pesan
}
