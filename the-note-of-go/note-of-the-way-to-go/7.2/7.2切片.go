package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 切片slice是对数组一个连续片段的引用(该数组称之为相关数组，一般为匿名)
	// 因此切片是一个引用类型(更类似于C/C++中的数组类型或Python中的list类型)
	// 这个片段可以是整个数组,或是由起始和中指索引标识的一些项的子集(终止索引标识的项不在切片内)
	// 切片提供了一个相关数组的动态窗口

	// 切片是可索引的，可由len()函数获取长度

	// 给定项的切片索引可能比相关数组的相同元素的索引小。切片长度可在运行时修改。
	// 最小为0，最大为相关数组长度，切片是一个长度可变的数组

	// cap()函数是计算切片容量的函数，测量切片最长可达到多少：它等于切片长度+数组除切片之外的长度
	// 切片s，cap(s)是从s[0]到数组末尾的数组长度。切片长度永远不会超过它的容量，所以 0 <= len(s) <= cap(s)

	// 多个切片如果表示同一个数组的片段，则它们可共享数组；所以一个切片和相关数组的其他切片是共享存储的
	// 相反，不同数组代表不同存储。数组实际上是切片的构建块。

	// 切片的优点是引用，不需要额外的内存并且比使用数组更高效，Go中切片比数组更常见

	// 切片的声明格式: var identifier []type (不需说明长度)
	// 切片未初始化之前默认为nil，长度为0

	// 切片初始化格式：var slice1 []type = arr1[start:end]
	// slice1是数组arr1从start索引到end-1索引之间的元素构成的子集

	// 若 var slice1 []type = arr1[:]
	// 表示slice1等于完整的arr1数组( 这是arr1[0:len(arr1)]的一种缩写 )
	// 另一种表述方式为： slice1 = &arr1

	var arr1 [6]int
	var slice1 []int = arr1[2:5]
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of arr1 is %d\n",len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	slice1 = slice1[0:4]
	for i := 0; i <len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice is %d\n", cap(slice1))

	// 切片只能向后移动，不能向前移动，例如s2 = s2[1:]可以，但s2 = s2[-1:]会报错

	// 不要用指针指向slice，切片本身已是一个引用类型，所以它本身就是一个指针

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b[1:4],"\n")
	fmt.Println(b[:2],"\n")
	fmt.Println(b[2:],"\n")
	fmt.Println(b[:],"\n")

	// 7.2.2将切片传递给函数
	// 如果函数需要对数组做操作，可以将参数声明为切片。调用函数时，将数组分片，创建一个切片引用并传递给函数
	var arr = [5]int{0, 1, 2, 3, 4}
	sum(arr[:])
	fmt.Println(sum(arr[:]),"\n")

	// 7.2.3用make()创建一个切片
	// 当相关数组还没定义时，可用make()函数创建一个切片，
	// 同时创建好相关数组: var slice1 []type = make([]type, len)
	// 简写为：slice1 := make([]type, len)，这里len是数组的长度也是切片的初始长度
	// s2 := make([]int, 10)，则cap(s2) == len(s2) == 10
	// make接受两个参数：元素类型和切片的元素个数

	// 如果创建的slice1不占用整个数组，只占用以len为个数的个项，则 slice1 := make([]type, len, cap)
	// make使用方式是： func make([]T, len, cap)，cap是可选参数

	// 这两种方式都可生成相同的切片：
	// make([]int, 50, 100)
	// new([100]int)[0:50]
	var slice2 []int = make([]int, 10)
	for i := 0; i < len(slice2); i++ {
		slice2[i] = 5 * i
	}
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("Slice2 at %d is %d\n", i, slice2[i])
	}
	fmt.Printf("\nThe length of slice2 is %d\n", len(slice2))
	fmt.Printf("The capacity of slice2 is %d\n", cap(slice2))

	// 字符串是纯粹不可变的字节数组，也可被切分成切片

	// 7.2.4new()和make()区别
	// 都在堆上分配内存，但是它们的行为不同，适用于不同的类型
	// new(T)为每个新的类型T分配一片内存，初始化为0并返回类型为*T的内存地址
	// 这种方法返回一个指向类型为T，值为0的地址的指针，适用于值类型如数组和结构体；相当于 &T{}
	// make(T)返回一个类型为T的初始值，只适用于3中内建的引用类型：切片、map和channel
	s := make([]byte, 5)
	fmt.Printf("len(s)",len(s),"\n")
	fmt.Printf("cap(s)",cap(s),"\n")
	fmt.Printf("2-4 len(s)",len(s[2:4]),"\n")
	fmt.Printf("2-4 cap(s)",cap(s[2:4]),"\n")

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Printf("s1:",s1,"\n")
	fmt.Printf("s2:",s2,"\n")
	s2[1] = 't'
	fmt.Printf("s1:",s1,"\n")
	fmt.Printf("s2:",s2,"\n")

	// 理解new、make、slice、map、channel的关系
	// 1.slice、map、channel都是golang内建的一种引用类型，三者在内存中存在多个组成部分
	// 需要对内存组成部分初始化后才能使用，而make就是对三者进行初始化的一种操作方式
	// 2.new获取的是存储指定变量内存地址的一个变量，对于变量内部结构并不会执行相应的初始化操作
	// 所以slice、map、channel需要make进行初始化并获取对应的内存地址，而非new简单的获取内存地址

	// 7.2.5多维切片
	// 与数组一样，切片也可以由一维组合成高维。
	// 通过分片的分片(或者切片的数组)，长度可以任意动态变化，所以Go语言的多维切片可以任意切分
	// 且内层的切片必须单独分配(通过make函数)

	// 7.2.6bytes包
	// 类型[]byte的切片十分常见，Go语言有一个bytes包专门用来提供这种类型的操作方法

	// bytes包和字符串包十分类似(4.7节)。它还包含一个十分有用的类型Buffer,例如
	// import "bytes"
	// type Buffer struct { ... }
	// 这是一个长度可变的bytes的buffer，提供Read和Write方法，因为读写长度未知的bytes最好使用buffer

	// Buffer定义：var buffer bytes.Buffer
	// 或使用new获得一个指针: var r *bytes.Buffer = new(bytes.Buffer)
	// 或通过函数: func NewBuffer(buf []byte) *Buffer  创建一个Buffer对象并用buf初始化，NewBuffer最好用在从buf读取的时候使用

	// 通过buffer串联字符串，类似Java的StringBuilder类
	// 下例中，创建一个buffer，通过buffer.WriteString(s)方法将字符串s追加到后面，再通过buffer.String()方法转为string
	var buffer bytes.Buffer
	for {
		if s, ok := getNextString(); ok {   //getNextString方法暂未展示
			buffer.WriteString(s)
		} else {
			break
		}
	}
	fmt.Print(buffer.String(), "\n")
	// 这种实现方式比使用 += 更节省内存和CPU，尤其是要串联的字符串数目很多时。
}
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}