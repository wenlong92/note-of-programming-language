package main
import (
	"fmt"
)

func main() {
	// 10.2.1结构体工厂
	// Go语言不支持面向对象编程语言中那样的构造子方法，但可以很容易在Go中是实现“构造子工厂”方法
	// 为了方便通常会为类型定义一个工厂，按惯例，工厂名字以new或New开头，例如File结构体：
	// 调用工厂方法
	f := NewFile(10, "./test.txt")
	// Go语言中经常像上面这样在工厂方法里使用初始化来简便的实现构造函数。
	// 如果File是一个结构体类型，则表达式new(File)和&File{}等价

	// 10.2.2 map 和 struct VS new()和make()
}
type File struct {
	fd int       // 文件描述符
	name string  // 文件名
}
// 下例是File结构体类型对应的工厂方法，它返回一个指向结构体实例的指针：
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}