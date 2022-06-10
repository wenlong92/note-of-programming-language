package main
import (
	"log"
	"runtime"
)

func main() {
	// 可以使用runtime或log包中的特殊函数进行调试
	// 包runtime中的函数Caller()提供了相应的信息
	// 因此可以在需要的时候实现一个where()闭包函数来打印函数执行的位置
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d",file,line)
	}
	where()
	where()
	where()
}
