package main

import (
	"fmt"
)

func main() {
	// 8.1.1概念
	// map是引用类型，声明如下：
	// var map1 map[keytype]valuetype
	// var map1 map[string]int

	// 声明时不需知道map长度，map可以动态增长
	// 未初始化的map的值是nil

	// key可以是任意用 == 或!=操作符比较的类型，例如string、int、float
	// 所以数组、切片和结构体不能作为key(含有数组切片的结构体不能作为key，只包含内建类型的struct可作为key)
	// 但是指针和接口类型可以
	// 如果要用结构体作为key，可以提供Key()和Hash()方法，这样可通过结构体的域计算出唯一数字或字符串的key

	// value可以是任意类型，通过使用空接口类型可以存储任意值，但使用这种类型作为值需要先做一次类型断言

	// map传递给函数的代价很小：32位机器上占4个字节，64位机器上占8个字节，无论实际上存储了多少数据
	// 通过key在map中寻找值很快，比线性查找快，但仍比从数组和切片的索引中直接读取慢100倍，如果在乎性能建议用切片

	// map也可用函数作为自己的值，这样可用来做分支结构：key用来选择要执行的函数

	// key1对应的值可通过赋值符号设置val1：map1[key1] = val1

	// v := map1[key1]可将key1对应的值赋值给v，若没有key1存在，v将被赋值为map1的值类型的空值

	// len(map1)方法可获得map中pair数目，这个数目可伸缩，因为map-pairs在运行时可动态添加删除

	var mapList map[string]int
	var mapAssigned map[string]int

	mapList = map[string]int{"one":1, "two":2}  //map可用{key:val}描述方法初始化，就像数组和结构体
	mapCreated := make(map[string]float32)
	mapAssigned = mapList

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is : %d\n", mapList["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapList["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapList["ten"])
	fmt.Println(mapCreated)
	// map是引用类型，内存用make方法分配
	// map初始化：var map1 = make(map[keytype]valuetype)
	// 或简写为：map1 ：= make(map[keytype]valuetype)
	// 上例中mapCreated就是这种方式创建的
	// 相当于mapCreated := map[string]float32{}
	// mapAssigned也是mapList的引用，对mapAssigned的修改会影响到mapList的值

	// 不要使用new，要用make构造map

	// 8.2.1map容量
	// 和数组不同，map可根据新增的key-value对 动态伸缩，不存在固定长度或最大限制
	// 但也可以选择表明map初始容量capacity，像这样 make(map[keytype]valuetype, cap)
	// 例如：map2 := make(map[string]float32, 100)

	// 出于性能的考虑，对于大的map或者会快速扩张的map，即使只是知道大概容量，也最好先表明
	// 例如：
	noteFrequency := map[string]float32 {
		"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
		"G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440}

		// 8.1.3用切片作为map的值
		// 一个key对应多个值时，可以将value定义为[]int类型或其他类型的切片
		mp1 := make(map[int][]int)
		mp2 := make(map[int]*[]int)
}