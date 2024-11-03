package collect

import (
	"github.com/gocolly/colly/v2"
	"strconv"
	"strings"
	"weather-colly/global"
)

type CollyCityWeatherCode struct {
	Name string // 城市名
	Code int    // 气象代码
}

var (
	targetUrl = "https://www.cnblogs.com/trigger-cn/p/17452078.html"
)

func CollyWeatherCode() []CollyCityWeatherCode {
	var result []CollyCityWeatherCode
	collector := colly.NewCollector()
	// 设置请求头，模拟浏览器
	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36")
	})
	// 处理响应
	collector.OnHTML("#cnblogs_post_body > ul", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, element *colly.HTMLElement) {
			text := element.Text
			split := strings.Split(text, ",")
			code, _ := strconv.Atoi(split[1])
			cityWeatherCode := CollyCityWeatherCode{Code: code, Name: split[0]}
			result = append(result, cityWeatherCode)
		})
	})
	// 爬取页面
	err := collector.Visit(targetUrl)
	if err != nil {
		global.Logger.Errorln(err.Error())
	}
	return result
}
