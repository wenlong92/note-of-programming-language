package main
import (
	"fmt"
	"io"
	"log"
)

func main() {
	// 关键字defer允许我们推迟到函数返回之前(或任意位置执行return语句之后)一刻才执行某个语句或函数
	// (return语句同样可以包含一些操作，而不是单纯的返回某个值)
	function1()

	// 有多个defer行为被注册时，会以逆序执行(类似栈，后进先出)
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}

	// defer可以进行一些函数执行完成后的首尾工作
	// 1.关闭文件流
	// defer file.Close()
	// 2.解锁一个加锁的资源
	// defer mu.Unlock()
	// 3.打印最终报告
	// printHeader()
	// defer printFooter()
	// 4.关闭数据库链接
	// defer disconnectFromDB(),例如：
	doDBOperations()

	// 使用defer实现代码追踪
	// 在进入和离开某个函数打印相关消息

	// 使用defer语句记录函数的参数与返回值
	func1("Go")

}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: deferred until the end of the calling function!\n")
}

func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB()
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returing from function here!")
	return 
}

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}