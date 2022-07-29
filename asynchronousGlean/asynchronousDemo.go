package asynchronousGlean

import (
	"fmt"
	"time"
)

type AsynchronousDemo struct {
}

func (data *AsynchronousDemo) Run() error {
	fmt.Println("asynchronous start")
	c := make(chan int)
	go sum(24, 18, c)
	go another(c)
	fmt.Println("asynchronous running")
	time.Sleep(10 * time.Second)
	fmt.Println("asynchronous complete")
	return nil
}

func (data *AsynchronousDemo) Stop() error {
	return nil
}

func (data *AsynchronousDemo) Close() error {
	return nil
}

func sum(x, y int, c chan int) {
	fmt.Println("sum start")
	time.Sleep(5 * time.Second)
	c <- x + y
}

func another(c chan int) {
	fmt.Println("another start")
	fmt.Println("sum data: ", <-c) //管道有数据了直接继续执行，相当于异步通知
	fmt.Println("another complete")
}
