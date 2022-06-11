package main

import (
	"fmt"
)

func main() {
	// 7.1.1概念
	// 数组是具有相同唯一类型的一组已编号且长度固定的数据项序列(这是一种同构的数据结构)
	// 这种类型可以是任意的原始类型(整型、字符串)或自定义类型。
	// 数组长度必须是一个常量表达式，且是一个非负整数。
	// 数组长度也是数组类型的一部分，所以[5]int 和[10]int是属于不同类型的
	// 数组编译时，值初始化是按照数组顺序完成的

	// 若想让数组元素类型为任意类型的话，可使用空接口作为类型(11章)
	// 当使用值时，必须先做一个类型判断(11章 )

	// 数组元素可以通过索引读取或修改。
	// 元素的数目(也叫做数组长度或大小)必须是固定且声明时就给出(编译时须知道数组长度以分配内存)
	// 数组长度最大为2GB

	// 声明格式如下：
	// var identifier [len]type
	// 例如： var arr1 [5]int

	// 每个元素是一个整型值，声明数组时所有元素自动初始化为默认值0

	// 索引项为i的数组元素赋值可这么操作：arr[i] = value 因此数组时可变的

	// 只有有效的索引才能被使用
	// 使用大于等于len(arr1)的索引时：会提示或报错

	// 由于索引的存在，遍历数组的方法是使用for结构
	var arr1 [5]int
	
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at index %d is %d\n",i, arr1[i])
	}
	// for循环中条件 i < len(arr1) 若写成 i <= len(arr1)会产生越界错误

	// 也可以使用for-range的生成方式
	// for i,_ := range arr1 {}
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	// Go语言中数组是一种值类型(不同于C/C++中是指向首元素的指针)
	// 所以可以通过new()来创建 var arr1 = new([5]int)
	// 这种方式和 var arr2 [5]int 的区别
	// arr1的类型时 *[5]int 而arr2的类型是 [5]int
	// 这样的结果就是把一个数组赋值给另一个时，需要再做一次数组内存拷贝，例：
	// arr2 := *arr1
	// arr2[2] = 100
	// 这样两个数组就有了不同的值，在赋值后修改arr2不会对arr1生效
	// 所以在函数中数组作为参数传入时，如func1(arr2)，会产生一次数组拷贝，func1不会修改原始数组arr2
	// 若想修改原数组，则arr2须通过&操作符以引用的方式传过来，如：func1(&arr2)
	var ar [3]int
	f(ar)
	fp(&ar)

	// 7.1.2数组常量
	// 如果数组值提前知道，可以通过数组常量的方法初始化数组，而不用依次使用 [] = 方法(所有的组成元素都有相同的常量语法)
	var arrkeyValue = [5]string{3: "Chris", 4: "Ron"}
	for i := 0; i < len(arrkeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrkeyValue[i]) //索引值0、1、2的value为空，索引值3、4的value分别为Chris和Ron
	}

	// 7.1.3多维数组
	// 可以用一维数组组装成多维数组 [3][5]int  [2][2][2]float64
	// 内部数组总是长度相同。Go语言的多维数组是矩形式(例外是切片的数组)

	// 7.1.4将数组传递给函数
	// 将大数组传递给函数会消耗很多内存，两种方法避免
	// 1.传递数组的指针(Go中不常用)
	// 2.使用数组的切片
	// 传递数组指针的例子：
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array)
	fmt.Printf("The sum of the array is: %f", x)
}
func f(a [3]int) {
	fmt.Println(a)
}
func fp(a *[3]int) {
	fmt.Println(a)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}