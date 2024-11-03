package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"weather-colly/config"
	"weather-colly/global"
)

var configPath = "settings.yaml"

func InitConfig() {
	conf := config.Config{}
	file, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Errorf("get yaml config err %s", err))
	}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		panic(fmt.Errorf("unmarshal yaml config err %s", err))
	}
	log.Println("config yaml file load init success.")
	global.Config = &conf
}
