package main

import "fmt"

// Go内置的运算符有：算术运算符、关系运算符、逻辑运算符、位运算符、赋值运算符、其他运算符

func main() {
	// 算术运算符： / 相除； % 求余；++ 自增； -- 自减；

	// 关系运算符： == 是否相等； != 是否不等；

	// 逻辑运算符： && 两边操作数都是True，则为True； || 两边的操作数有一个为True，则为True； ! 取反；

	// 位运算符：& 当两边都为1时，为1，否则为0； | 当两边都为0时，为0，否则为1； ^ 当两边不同时，为1，否则为0；<< 左移n位就是乘以2的n次方； >> 右移n位除以2的n次方；
	// 当p=0,q=1时 p&q为0，p|q为1，p^q为1
	// 当p=0,q=0时 p&q为0，p|q为0，p^q为0
	// 当p=1,q=0时 p&q为0，p|q为1，p^q为1
	// 当p=1,q=1时 p&q为1，p|q为1，p^q为0
	var a1 uint = 60 // 60 = 0011 1100
	var b1 uint = 13 // 13 = 0000 1101
	var c1 uint = 0
	c1 = a1 & b1
	fmt.Printf("第一行c的值为 %d\n", c1) // 0000 1100，为12

	c1 = a1 | b1
	fmt.Printf("第二行c的值为 %d\n", c1) // 0011 1101，为61

	c1 = a1 ^ b1
	fmt.Printf("第三行c的值为 %d\n", c1) // 0011 0001，为49

	c1 = a1 << 2
	fmt.Printf("第四行c的值为 %d\n", c1) // 1111 0000，为240

	c1 = a1 >> 2
	fmt.Printf("第五行c的值为 %d\n", c1) // 0000 1111，为15

	// 赋值运算符：
	// += 相加后再赋值； -= 相减后再赋值； *= 相乘后再赋值； /= 相除后再赋值； %= 求余后再赋值；
	// <<= 左移后赋值； >>= 右移后赋值； &= 按位与后赋值； ^= 按位异或后赋值； |= 按位或后赋值；
	var a int = 21
	var c int

	c = a
	fmt.Printf("第1行 = 运算符实例，c值为 %d\n", c) // 21

	c += a
	fmt.Printf("第2行 = 运算符实例，c值为 %d\n", c) // 42

	c -= a
	fmt.Printf("第3行 = 运算符实例，c值为 %d\n", c) // 21

	c *= a
	fmt.Printf("第4行 = 运算符实例，c值为 %d\n", c) // 441

	c /= a
	fmt.Printf("第5行 = 运算符实例，c值为 %d\n", c) // 21

	c = 200 // 1100 1000

	c <<= 2
	fmt.Printf("第6行 = 运算符实例，c值为 %d\n", c) // 800

	c >>= 2
	fmt.Printf("第7行 = 运算符实例，c值为 %d\n", c) // 200

	c &= 2
	fmt.Printf("第8行 = 运算符实例，c值为 %d\n", c) // 0

	c ^= 2
	fmt.Printf("第9行 = 运算符实例，c值为 %d\n", c) // 2

	c |= 2
	fmt.Printf("第10行 = 运算符实例，c值为 %d\n", c) // 2

	// 其他运算符： & 返回变量存储地址，&a将给出变量的实际地址； * 指针变量，*a是一个指针变量；
	var a2 int = 4
	var b2 int32
	var c2 float32
	var ptr *int

	fmt.Printf("第1行 a2变量类型为 %T\n", a2)
	fmt.Printf("第2行 b2变量类型为 %T\n", b2)
	fmt.Printf("第3行 c2变量类型为 %T\n", c2)

	ptr = &a2
	fmt.Printf("a2的值 为%d\n", a2)
	fmt.Printf("ptr2的值 为%d\n", ptr)
}
