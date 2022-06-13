package main
import (
	"fmt"
	"time"
)

func main() {
	// 当前时间time.Now()
	t := time.Now()
	fmt.Println(t)

	// 获取时间的一部分
	fmt.Printf("%02d.%02d.%4d\n",t.Day(),t.Month(),t.Year())

	// UTC表示通用协调世界时间
	t1 := time.Now().UTC()
	fmt.Println(t1)
}