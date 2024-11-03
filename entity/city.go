package entity

import (
	"encoding/json"
	"fmt"
	"os"
)

type City struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Province string `json:"province"`
	City     string `json:"city"`
}

var (
	CityPath = "./assets/city.json"
)

func ReadCity() []City {
	city, err := os.ReadFile(CityPath)
	if err != nil {
		fmt.Println("city读取失败", err)
	}
	var cities []City
	err = json.Unmarshal(city, &cities)
	if err != nil {
		fmt.Println("city转换失败", err)
	}
	return cities
}
