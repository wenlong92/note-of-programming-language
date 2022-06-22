package main

import (
	"fmt"
	"runtime"
)

func main()  {
	// Go开发者不需要写代码来释放程序中不再使用的变量和结构占用的内存
	// Go运行时中有一个独立进程，垃圾收集器(GC)，会处理这些事情，它搜索不再使用的变量然后释放他们的内存
	// 可通过runtime包访问GC进程

	// 调用runtime.GC()函数可显式的触发GC，但只在某些罕见的场景下才有用，比如当内存资源不足时调用runtime.GC()
	// 它会再次函数执行的点上立即释放一大片内存，此时程序可能会有短时的性能下降

	// 如果想知道当前内存状态，可使用：
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
	// 会给出已分配内存的总量

	// 如果需要一个对象obj被从内存移除💰执行一些特殊操作，比如写到日志文件中，可通过如下方式调用函数实现：
	// runtime.SetFinalizer(obj, func(obj *typeObj))

	// func(obj *typeObj)需要一个typeObj 类型的指针参数obj，特殊操作会在它上面执行，func也是个匿名函数

	// 在对象被GC进程选中并从内存中移除之前，SetFinalizer都不会执行，即使程序正常结束或者发生错误。
}