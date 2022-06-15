package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	// 在字符串中对正则表达式模式进行匹配
	// 简单模式用Match方法即可：
	// ok, _ := regexp.Match(pat, []byte(searchIn))
	// 变量ok返回true或false，也可用MatchString：
	// ok, _ := regexp.MatchString(pat,searchIn)

	// 更多方法中，须先将正则模式通过Compile方法返回一个Regexp对象，然后将掌握一些匹配、查找、替换相关功能
	searchIn := "John: 2578.34 William: 45567.23 steve: 5632.18"
	pat := "[0-9]+.[0-9]+"
	
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat,[]byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}

	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)

	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)

	// Compile函数也可能返回一个错误
	// 当用户输入或从数据中获取正则表达式时，要检验正确性
	// 也可用MustCompile方法，和Compile一样可以检验正则的有效性，但是当正则不合法时程序将panic
}