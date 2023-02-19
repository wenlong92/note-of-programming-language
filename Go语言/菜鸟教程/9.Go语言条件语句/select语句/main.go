package main

import "fmt"
import "time"

func main() {
	// select是Go中一个控制结构，类似于switch语句
	// 只能用于通道操作，每个case必须是一个通道操作，要么发送要么接收
	// select语句会监听所有指定通道上的操作，一旦其中一个通道准备好就会执行相应的代码块
	// 如果多个通道都准备好，会随机选择一个通道执行，如果所有通道都没准备好，会执行default块中的代码

	/*
		语法：
			1.每个case都必须是一个通道
			2.所有channel表达式都会被求值
			3.所有被发送的表达式都会被求值
			4.如果任意某个通道可以进行，它就执行，其他忽略
			5.如果有多个case都可以运行，select随机选出一个执行，其他不会执行
				。如果有default字距，则执行该语句
				。如果没有default字句，select将阻塞，知道某个通道可以运行；Go不会重新对channel或值进行求值
	*/
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// 以下实例中，定义两个通道，并启动两个协程从这两个通道中获取数据
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "from 1"
		}
	}()
	go func() {
		for {
			ch2 <- "from 2"
		}
	}()

	for {
		select {
		case msg10 := <-ch1:
			fmt.Println(msg10)
		case msg20 := <-ch2:
			fmt.Println(msg20)
		default:
			fmt.Println("no message received")
		}
	}

}
