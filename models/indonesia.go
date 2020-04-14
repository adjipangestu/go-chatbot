package models

import (
	"encoding/json"
	"fmt"
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

type Provinsi struct {
	Name   			string   `json:"Provinsi"`
	Kasus_Posi     	int32    `json:"Kasus_Posi"`
	Kasus_Semb    	int32    `json:"Kasus_Semb"`
	Kasus_Meni   	int32    `json:"Kasus_Meni"`
}

type Attributes struct {
	Attributes Provinsi   `json:"attributes"`
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

func (i *Indonesia) GetDataProvinsi() string {
	url_prov := "https://api.kawalcorona.com/indonesia/provinsi/"

	response_prov, err := http.Get(url_prov)
	if err != nil {
		log.Fatal(err)
	}
	defer response_prov.Body.Close()

	responseData_prov, err := ioutil.ReadAll(response_prov.Body)

	var provinsi []Attributes
	jsonErr_prov := json.Unmarshal([]byte(responseData_prov), &provinsi)
	if jsonErr_prov != nil {
		panic(jsonErr_prov)
	}
	pesan := ""
	for _, v := range provinsi {
		if v.Attributes.Name == "Lampung" {
			name := v.Attributes.Name
			positif := fmt.Sprintf("%d", v.Attributes.Kasus_Posi)
			sembuh := fmt.Sprintf("%d", v.Attributes.Kasus_Semb)
			meninggal := fmt.Sprintf("%d", v.Attributes.Kasus_Meni)
			pesan = "*Data Kasus Covid-19 "+ name +"*\n\nPositif : " + positif + ",\nSembuh : " + sembuh + ",\nMeninggal : " + meninggal + "\n\n#jagakesehatan #dirumahaja"
		}
	}

	return pesan
}