package core

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
	"weather-colly/crontab"
	"weather-colly/global"
)

func InitCronTab() {
	config := global.Config

	if !config.Crontab.Open {
		global.Logger.Infoln("cron is closed!")
	}
	// put all works
	crontab.PutWork(crontab.CollyWeather{})

	c := cron.New()
	for _, work := range config.Crontab.Works {
		if _, ok := crontab.WorkMap[work.Name]; !ok {
			global.Logger.Errorln(fmt.Sprintf("%s work is not exist!", work.Name))
			continue
		}
		err := c.AddJob(work.Crontab, crontab.WorkMap[work.Name])
		if err != nil {
			return
		}
	}

	// 在新的 goroutine 中启动定时任务
	go func() {
		c.Start()
		global.Logger.Infoln("crontab is running!")

		t1 := time.NewTimer(time.Second * 10)
		for {
			select {
			case <-t1.C:
				t1.Reset(time.Second * 10)
			}
		}
	}()
}
