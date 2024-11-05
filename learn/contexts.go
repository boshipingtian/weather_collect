package learn

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

// TimeoutTest Timeout 超时控制
func TimeoutTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	go SlowOperation(ctx)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("goroutines: ", runtime.NumGoroutine())
		}
	}()
	time.Sleep(6 * time.Second)
}

func SlowOperation(ctx context.Context) {
	done := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		done <- 1
	}()

	select {
	case <-ctx.Done():
		fmt.Println("SlowOperation timeout:", ctx.Err())
	case <-done:
		fmt.Println("Work done:")
	}
}

// QueryFrameworkStats
// 根据github仓库统计信息接口查询某个仓库信息
func QueryFrameworkStats(ctx context.Context, framework string) <-chan string {
	stats := make(chan string)
	go func() {
		repos := "https://api.github.com/repos/" + framework
		req, err := http.NewRequest("GET", repos, nil)
		if err != nil {
			return
		}
		req = req.WithContext(ctx)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		defer func(Body io.ReadCloser) {
			if err := Body.Close(); err != nil {
				panic(err)
			}
		}(resp.Body)
		stats <- string(data)
	}()

	return stats
}

func QueryGithubStat() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	framework := "gin-gonic/gin"
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case statsInfo := <-QueryFrameworkStats(ctx, framework):
		fmt.Println(framework, " fork and start info : ", statsInfo)
	}
}
