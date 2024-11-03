package collect

import "time"

type CollyWeather struct {
	City            string     // 城市
	Datetime        *time.Time // 时间
	Temperature     float32    // 温度
	Humidity        float32    // 紫外线
	BodyTemperature float32    // 体感温度
	Visibility      float32    // 能见度
	Precipitation   float32    // 降雨量
	Pressure        float32    // 大气压
	Wind            float32    // 风速
	WindDirection   float32    // 风向
}

func CollectWeather() {

}
