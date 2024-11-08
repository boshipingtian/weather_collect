package learn

import (
	"fmt"
	"sync"
	"time"
)

// goroutines
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// channel
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

// buffer channel
func bufferChannel() {
	ch := make(chan int, 2)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("ch <-", i)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(5 * time.Second)
			x := <-ch
			fmt.Println(x, "<- ch")
		}
	}()
	time.Sleep(26 * time.Second)
}

// close And range
func fibonacciForClose(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func rangeChannelForRange() {
	c := make(chan int, 10)
	go fibonacciForClose(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

// select
//
// select 语句使一个 Go 程可以等待多个通信操作。
//
// select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。
func fibonacciForSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func fibonacciRun() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciForSelect(c, quit)
}

// select default
//
// 当 select 中的其它分支都没有准备好时，default 分支就会执行。
//
// 为了在尝试发送或者接收时不发生阻塞，可使用 default 分支
func selectDefault() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// Tree 不同二叉树的叶节点上可以保存相同的值序列。例如，以下两个二叉树都保存了序列 `1，1，2，3，5，8，13`。
//
// 在大多数语言中，检查两个二叉树是否保存了相同序列的函数都相当复杂。 我们将使用 Go 的并发和信道来编写一个简单的解法。
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk 遍历树 t，并树中所有的值发送到信道 ch。
func Walk(t *Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same 判断 t1 和 t2 是否包含相同的值。
func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	v1 := <-ch1
	v2 := <-ch2
	return v1 == v2
}

func binaryTreeRun() {
	tree1 := Tree{
		Left:  &Tree{Value: 1},
		Value: 2,
		Right: &Tree{Value: 3},
	}
	tree2 := Tree{
		Left:  &Tree{Value: 1},
		Value: 2,
		Right: &Tree{Value: 3},
	}
	fmt.Println(Same(&tree1, &tree2))
}

// 我们已经看到信道非常适合在各个 Go 程间进行通信。
//
// 但是如果我们并不需要通信呢？比如说，若我们只是想保证每次只有一个 Go 程能够访问一个共享的变量，从而避免冲突？
//
// 这里涉及的概念叫做 *互斥（mutual*exclusion）* ，我们通常使用 *互斥锁（Mutex）* 这一数据结构来提供这种机制。
//
// Go 标准库中提供了 sync.Mutex 互斥锁类型及其两个方法：
//
// Lock
// Unlock
// 我们可以通过在代码前调用 Lock 方法，在代码后调用 Unlock 方法来保证一段代码的互斥执行。参见 Inc 方法。
//
// 我们也可以用 defer 语句来保证互斥锁一定会被解锁。参见 Value 方法。
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}
