package main

import (
	"fmt"
	"time"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
		c <- sum //给通道传值
	}
}
func main() {
	//chanBufferTest()
	//rangeClose()
	//selectTest()
	timeOutTest()
}

func timeOutTest() {
	c := make(chan int)
	o := make(chan bool)
	//v:=1
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				o <- true
				break
			}
		}
	}()
	fmt.Println("start")
	//<-o
	//<-c
	//o<-true
	//c<-0
	a := <-o
	fmt.Println(a)
}
func rangeClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func selectTest() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciSelect(c, quit)
}
func fibonacciSelect(c, quit chan int) {
	x, y := 1, 1
	///一直等待
	for {
		//在通道被赋值和取值时调用
		select {
		//赋值c通道
		case c <- x:
			x, y = y, x+y
		//	取值quit
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

//显示关闭
func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	//显示关闭通道，关闭后泽无法再传输数据
	//只能在输入值的时候才能关闭
	close(c)
}

//最后一位是阻塞状态赋值，前几位则是无阻塞
func chanBufferTest() {
	c := make(chan int, 2)
	c <- 2
	c <- 4
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func chanTest() {
	a := []int{1, -2, -55, 333, 55, 92}
	//无缓存通道，相当于一直置于一个位置
	c := make(chan int)
	//sum(a,c)
	go sum(a[len(a)/2:], c)
	go sum(a[:len(a)/2], c)
	//阻塞状态
	x, y := <-c, <-c //取值
	fmt.Println(x, y, x+y)
}
