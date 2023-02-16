package main

import "fmt"

func main() {
	// 字符串连接
	fmt.Println("Google" + "Runoob")

	// 格式化字符串
	// Sprintf 根据格式化参数生成格式化的字符串并返回该字符串。
	var stockcode = 123
	var enddate = "2020-12-02"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	// Printf 根据格式化参数生成格式化的字符串并写入标准输出。
	fmt.Printf(url, stockcode, enddate)
}
