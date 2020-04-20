package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)



type Provinsi struct {
	Name    	string    `json:"name"`
	Positif     int32   `json:"positif"`
	Sembuh    	int32    `json:"sembuh"`
	Meninggal   int32    `json:"meninggal"`
}

type Data struct {
	Provonsi   []Provinsi `json:"data"`
}

func main() {
	//url_prov := "http://api.sekelik.com/corona/provinsi"
	//response_prov, err := http.Get(url_prov)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer response_prov.Body.Close()
	//
	//responseData_prov, err := ioutil.ReadAll(response_prov.Body)
	//
	//var provinsi Data
	//jsonErr_prov := json.Unmarshal([]byte(responseData_prov), &provinsi)
	//if jsonErr_prov != nil {
	//	panic(jsonErr_prov)
	//}
	//
	//pesan := ""
	//for _, v := range provinsi.Provonsi {
	//	if v.Positif == "Lampung" {
	//		name := v.Name
	//		positif := fmt.Sprintf("%d", v.Positif)
	//		sembuh := fmt.Sprintf("%d", v.Sembuh)
	//		meninggal := fmt.Sprintf("%d", v.Meninggal)
	//		pesan = "*Data Kasus Covid-19 "+ name +"*\n\nPositif : " + positif + ",\nSembuh : " + sembuh + ",\nMeninggal : " + meninggal + "\n\n#jagakesehatan #dirumahaja"
	//	}
	//}

	// Open our jsonFile
	jsonFile, err := os.Open("data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Successfully Opened data.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	responseData_prov, _ := ioutil.ReadAll(jsonFile)

	var provinsi Data
	jsonErr_prov := json.Unmarshal([]byte(responseData_prov), &provinsi)
	if jsonErr_prov != nil {
		panic(jsonErr_prov)
	}

	msg := "covid:java"
	result := strings.Replace(msg, "covid:", "", 1)
	fmt.Println(result)
	//re := regexp.MustCompile(result)
	//for _, v := range provinsi.Provonsi {
	//	pesan := ""
	//	results := re.MatchString(strings.ToLower(v.Name))
	//
	//	if results == true {
	//		name := v.Name
	//		positif := fmt.Sprintf("%d", v.Positif)
	//		sembuh := fmt.Sprintf("%d", v.Sembuh)
	//		meninggal := fmt.Sprintf("%d", v.Meninggal)
	//		pesan = "*Data Kasus Covid-19 " + name + "*\n\nPositif : " + positif + ",\nSembuh : " + sembuh + ",\nMeninggal : " + meninggal + "\n\n#jagakesehatan #dirumahaja\n"
	//	}
	//
	//	fmt.Println(pesan)
	//}
}