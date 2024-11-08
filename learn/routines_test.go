package learn

import (
	"fmt"
	"testing"
	"time"
)

func Test_say(t *testing.T) {
	go say("world")
	say("hello")
}

func Test_sum(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func Test_bufferChannel(t *testing.T) {
	bufferChannel()
}

func Test_rangeChannel(t *testing.T) {
	rangeChannelForRange()
}

func Test_fibonacciRun(t *testing.T) {
	fibonacciRun()
}

func Test_selectDefault(t *testing.T) {
	selectDefault()
}

func Test_binaryTreeRun(t *testing.T) {
	binaryTreeRun()
}

func TestSafeCounter(t *testing.T) {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
