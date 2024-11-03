package entity

import (
	"encoding/json"
	"fmt"
	"os"
)

type Province struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Province string `json:"province"`
}

var (
	ProvincePath = "./assets/province.json"
)

func ReadProvince() []Province {
	province, err := os.ReadFile(ProvincePath)
	if err != nil {
		fmt.Println("province读取失败", err)
	}
	var provinces []Province
	err = json.Unmarshal(province, &provinces)
	if err != nil {
		fmt.Println("province转换失败", err)
	}
	return provinces
}
