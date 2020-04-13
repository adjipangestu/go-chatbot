package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Indonesia struct {
	Name    	string    `json:"name"`
	Positif     string    `json:"positif"`
	Sembuh    	string    `json:"sembuh"`
	Meninggal   string    `json:"meninggal"`
}

func (i *Indonesia) GetData() string{
	url := "https://api.kawalcorona.com/indonesia/"
	//indonesia
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	var indonesia []Indonesia
	jsonErr := json.Unmarshal([]byte(responseData), &indonesia)
	if jsonErr != nil {
		panic(jsonErr)
	}

	kasus_positif := indonesia[0].Positif
	kasus_sembuh := indonesia[0].Sembuh
	kasus_meninggal := indonesia[0].Meninggal

	pesan := ""
	pesan = "*Data Kasus Covid-19 Indonesia*\n\nPositif : " + kasus_positif + ",\nSembuh : " + kasus_sembuh + ",\nMeninggal : " + kasus_meninggal + "\n\n#jagakesehatan #dirumahaja"

	return pesan
}