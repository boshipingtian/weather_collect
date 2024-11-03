package entity

import (
	"encoding/json"
	"fmt"
	"os"
)

type Town struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Town     string `json:"town"`
}

var (
	TownPath = "./assets/town.json"
)

func ReadTown() []Town {
	town, err := os.ReadFile(TownPath)
	if err != nil {
		fmt.Println("town读取失败", err)
	}
	var towns []Town
	err = json.Unmarshal(town, &towns)
	if err != nil {
		fmt.Println("town转换失败", err)
	}
	return towns
}
