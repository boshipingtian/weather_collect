package entity

import (
	"encoding/json"
	"fmt"
	"os"
)

type Area struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
}

var (
	AreaPath = "./assets/area.json"
)

func ReadArea() []Area {
	area, err := os.ReadFile(AreaPath)
	if err != nil {
		fmt.Println("area读取失败", err)
	}
	var areas []Area
	err = json.Unmarshal(area, &areas)
	if err != nil {
		fmt.Println("area转换失败", err)
	}
	return areas
}
