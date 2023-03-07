package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you ")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you ")
}

func main() {
	// 接口：将所有的具有共性的方法定义在一起。任何其他类型只要实现了这些方法就是实现了这个接口
	// 接口可以让我们将不同的类型绑定到一组公共方法上，实现多态和灵活的设计
	// Go语言中的接口时隐式实现的，就是说，如果一个类型实现了一个接口定义的所有方法，name它就自动实现了该接口。因此，可以通过将接口作为参数来实现对不同类型的调用，从而实现多态

	// 例子一
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	// 例子二
	var s Shape
	s = Rectangle{width: 10, height: 5}
	fmt.Printf("矩形面积：%f\n", s.area())

	s = Circle{radius: 3}
	fmt.Printf("原型面积：%f\n", s.area())

}

type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
