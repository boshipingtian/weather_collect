package entity

import (
	"encoding/json"
	"fmt"
	"os"
)

type Country struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	CnName   string `json:"cnname"`
	FullName string `json:"fullname"`
}

var (
	CountryPath = "./assets/country.json"
)

// ReadCountry Read data from country.json and write it to models.Country{}
func ReadCountry() []Country {
	country, err := os.ReadFile(CountryPath)
	if err != nil {
		fmt.Println("country读取失败", err)
	}
	var countries []Country
	err = json.Unmarshal(country, &countries)
	if err != nil {
		fmt.Println("country转换失败", err)
	}
	return countries
}
