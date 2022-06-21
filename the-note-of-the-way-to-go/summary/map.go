package main

import "fmt"

func main() {
	// Map是一种无序的键值对的集合。最重要的一点就是通过key快读检索数据，key类似索引，指向数据的值。
	// Map是一种集合，可以像迭代数组和切片那种迭代它，不过无法决定返回顺序。因为Map是使用hash表来实现的。

	// 定义Map
	// 可使用内建函数make，也可使用map关键字定义：
	// var map_variable map[key_data_type]value_data_type
	// map_variable := make(map[key_data_type]value_data_type)

	// 如果不初始化map，会创建一个nil map   nil map不能用来存放键值对

	// 例子：
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	// 查看元素在集合中是否存在
	capital, ok := countryCapitalMap["American"]
	if(ok) {
		fmt.Println("American的首都是", capital)
	} else {
		fmt.Println("这里不存在American的首都")
	}


	// delete()函数
	// delete()函数用于删除集合的元素，参数为map和其对应的key
	countryCapitalMap1 := map[string]string{"France":"Paris","Italy":"Rome","Japan":"Tokyo"}
	delete(countryCapitalMap1, "France")
	fmt.Println(countryCapitalMap1)
}