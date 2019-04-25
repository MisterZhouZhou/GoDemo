package test

import (
	"time"
	"fmt"
)

func say(s string)  {
	for i:= 0; i< 5; i++ {
		time.Sleep(100*time.Microsecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func testGo()  {
	// go 语句开启一个新的运行期线程
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

/**
  通道类似于栈
 */
func test()  {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int)  {
	x, y := 0, 1
	for i:=0; i<n ;i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func test2()  {
	//interfaceT.TestMan()
	//errorT.TestDiviError()
	//go say("world")
	//say("hello")

	//c := make(chan int, 10)
	//// cap容量
	//go fibonacci(cap(c), c)
	//for i := range c{
	//	fmt.Println(i)
	//}
}