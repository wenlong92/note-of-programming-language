package main

import "fmt"

// 结构体类似其他语言中的 class，可以在结构体中定义多个字段，为结构体实现方法，实例化等
// 定义一个结构体 Student，并添加name、age 字段，并实现 hello()方法
type Student struct {
	name string
	age  int
}

//实现方法和实现函数之间的区别在于 func和函数名之间，加上改方法对应的实例名stu及其类型*Student
func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I am %s", person, stu.name)
}

func main() {
	stu := &Student{
		name: "Tom",
	}
	msg := stu.hello("Jack")
	fmt.Println(msg)

	//除此之外，还可以使用 new 实例化
	stu2 := new(Student)
	fmt.Println(stu2.hello("Alice"))
}
