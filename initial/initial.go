package initial

import "weather-colly/core"

func Init() {
	core.InitConfig()
	core.InitLogger()
	core.InitGorm()
}
