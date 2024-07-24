package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("filename.txt")
	if err != nil {
		fmt.Println(err)
	}

	if err := hello(""); err != nil {
		fmt.Println(err)
	}

	fmt.Println(get(5))
	fmt.Println("finished")
}

// 可以通过 errors.New返回自定义错误
func hello(name string) error {
	if len(name) == 0 {
		return errors.New("error: name is null")
	}
	fmt.Println("hello,", name)
	return nil
}

// 当发生不可预知错误时，可能会导致程序非正常退出，Go 语言中称之为 panic
// Go语言中提供了 defer 和 recover
func get(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some error happened!", r)
			ret = -1
		}
	}()
	arr := [3]int{2, 3, 4}
	ret = arr[index]
	return
}
