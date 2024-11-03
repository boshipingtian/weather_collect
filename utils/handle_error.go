package utils

import (
	"weather-colly/global"
)

func ExecuteWithErrorHandling(fn func() error) {
	if err := fn(); err != nil {
		global.Logger.Errorln("执行操作时出错: %v", err)
	}
}
