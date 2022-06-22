package main
import (
	"time"
	"fmt"
)

func main() {
	// 最简单的计算执行消耗时间的办法就是，在开始之前设置起始时间，再记录结束时的结束时间
	// 最后计算差值
	// 可以使用time包中的Now()和Sub函数
	start := time.Now()
	fmt.Printf("test")
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("time: %s\n", delta)
}