package main

import (
	"fmt"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	// Go语言不是一种“传统”的面向对象编程语言：它里面没有类和继承的概念
	
	// 但Go语言中有灵活的接口概念，通过它可实现很多面向对象的特性。
	// 接口提供了一种方式来说明对象的行为：如果谁能搞定这件事，他就可以用在这儿

	// 接口定义了一组方法(方法集)，但这些方法不包含(实现)代码：它们没有被实现(它们是抽象的)。接口里也不能包含变量

	// 通过一下格式定义接口：

	// type Namer interface {
	// 	Method1(param_list) return_type
	// 	method2(param_list) return_type
	// }

	// 上面Namer是一个接口类型

	// 按照约定，只包含一个方法的接口的名字由方法名加er后缀组成，例如Printer/Reader/Writer/Logger等
	// 还有一些不常用的方式(当后缀er不合适时)，比如Recoverable，此时接口名以able结尾，或者以 I 开头(像.NET或Java中那样)

	// 接口都很简短，通常会包含0个，最多3个方法

	// 不同于大多数面向对象编程语言，Go语言中接口可以有值，一个接口类型的变量或一个接口值: var ai Namer, 
	// ai 是一个多字数据结构，它的值是nil。它本质上是一个指针，虽然不完全是一回事。
	// 指向接口值的指针是非法的，它们不仅一点用也没有，还会导致代码错误。

	// 类型(比如结构体)可实现某个接口的方法集；这个实现可描述为，该类型的变量上的每一个具体方法所组成的集合
	// 包含了该接口的方法集。实现了Namer接口的类型的变量可以赋值给ai(即receiver的值)，方法表指针就指向了当前的方法实现
	// 当另一个实现了Namer接口的类型的变量被赋给ai，receiver的值喝方法表指针一会相应改变

	// 类型不需要显式声明它实现了某个接口：接口被隐式地实现。 多个类型可以实现同一个接口。

	// 实现某个接口的类型(除了实现接口方法外)可以有其他的方法

	// 一个类型可以实现多个接口

	// 接口类型可以包含一个实例的引用，该实例的类型实现了此接口(接口是动态类型)

	// 即使接口在类型之后裁定已，二者处于不同的包中，被单独编译：只要类型实现了接口中的方法，它就可以实现了此接口
	// 所有这些特性使得接口具有很大的灵活性
	
	// 例子一：
		sq1 := new(Square)
		sq1.side = 5

		var areaIntf Shaper
		areaIntf = sq1
		// shorter,without separate delaration:
		// areaIntf := Shaper(sq1)
		// or even:
		// areaIntf := sq1
		fmt.Printf("The square has area: %f\n", areaIntf.Area())
		// 这是多态的Go版本，多态是面向对象编程中一个广为人知的概念：根据当前的类型选择正确的方法，
		// 或者说：同一种类型在不同的实例上似乎表现出不同的行为

		// 扩展下上例，类型Rectangle也实现了Shaper接口。接着创建一个Shaper类型的数组，迭代他的每一个元素并在上面调用Area()方法，来展示多态行为：
		r := Rectangle{5, 3}
		q := &Square{5}
		// shapes := []Shaper{Shaper(r), Shaper(q)}
		// or shorter:
		shapes := []Shaper{r, q}
		fmt.Println("Looping through shapes for area...")
		for n, _ := range shapes {
			fmt.Println("Shape details: ", shapes[n])
			fmt.Println("Area of this shape is: ", shapes[n].Area())
		}

		// 下面一个更具体的例子：有两个类型stockPosition和car，他们都有一个getValue()方法，
		// 我们可以定义一个具有此方法的接口valuable。
		// 接着定义一个使用valuable类型作为参数的函数showValue()，所有实现了valuable接口的类型都可以用这个函数
		var o valuable = stockPosition{"GOOG", 577.20, 4}
		showValue(o)
		o = car{"BMW", "M3", 66500}
		showValue(o)


		// 有的时候，也会以一种稍微不同的方式来使用接口这个词：从某个类型的角度来看，
		// 它的接口指的是：它的所有导出方法，只不过没有显式地为这些导出方法额外顶一个接口而已。
}
type stockPosition struct {
	ticker string
	sharePrice float32
	count float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}