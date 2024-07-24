package main

import "fmt"

func main() {
	age := 18
	if age < 18 {
		fmt.Printf("kid")
	} else {
		fmt.Printf("Adult")
	}

	//简写形式
	if age1 := 18; age1 < 18 {
		fmt.Printf("Kid")
	} else {
		fmt.Printf("Adult")
	}

}
