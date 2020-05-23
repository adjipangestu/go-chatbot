package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	Update 	string 	`json:"last_update"`
	Kasus	Data `json:"kasus"`
	Meninggal 	Data `json:"meninggal"`
	Perawatan 	Data `json:"perawatan"`
}

type Data struct {
	Umur   Item   `json:"kelompok_umur"`
	Gejala   Item   `json:"gejala"`
	KondisiPenyerta   Item   `json:"kondisi_penyerta"`
}
type Item struct {
	List []*List	`json:"list_data"`
}

type List struct {
	Key string `json:"key"`
	Doc float64 `json:"doc_count"`
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 0, 64)
}

func main()  {
	url := "https://data.covid19.go.id/public/api/data.json"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(data))
	var responseObject Response
	json.Unmarshal(data, &responseObject)

	pesan := ""
	s := []string{}
	for _, value := range responseObject.Meninggal.KondisiPenyerta.List {
		pesan = value.Key +" ("+ FloatToString(value.Doc)+ "%)"
		s = append(s, pesan)
	}

	datas := strings.Join(s, "\n")
	fmt.Println(datas)



}
