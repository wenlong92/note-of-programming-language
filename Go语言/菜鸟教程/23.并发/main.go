package main

import (
	"fmt"
	"time"
)

func main() {
	// Go语言支持并发，只需要通过go关键字来开启goroutine即可
	// goroutine是轻量级线程，调度时有Golang运行时进行管理

	go say("world")
	say("Hello ")
	// 以上代码，输出的hello和world是没有固定先后顺序的，因为两个是goroutine在执行

	// channel 通道：用来传递数据的一个数据结构
	// 通道可用于两个goroutine之间通过传递一个指定类型的值来同步运行和通讯。操作符<-用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道

	// 声明一个通道
	ch := make(chan int)
	fmt.Println(ch)

	// ch <- v   把v发送到通道ch
	// v := <- ch 从ch接收数据并赋值给v

	// 默认情况下，通道不带缓冲区。发送端发送数据，同时必须有接收端相应的接收数据
	s := []int{7, 2, 8, -9, 5, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	// 带缓冲区的通道允许发送端对的数据发送和接收端的数据获取处于异步状态，就是说发送单发送的数据可以放在缓冲区里，等待接收端去获取数据，而不是理科需要接收端去获取数据
	// 缓冲区大小是有限制的，缓冲区一满，发送端就无法再发送数据
	// 如果通道不带缓冲，发送方会阻塞知道接收方从同道中人接收值。如果通道带缓冲，发送方则会阻塞知道发送的值被拷贝到缓冲区内；如果缓冲区满，则需要等待知道某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞

	ch1 := make(chan int, 2) // 缓冲区大小为2
	ch1 <- 1
	ch1 <- 2
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	// Go 通过range关键字实现遍历读取到的数据，类似于数组或切片
	c1 := make(chan int, 10)
	go fibonacci(cap(c1), c1)
	for i := range c1 {
		fmt.Println(i)
	}

}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, y+x
	}
	close(c)
}
