package main

func main() {
	// 字符串是一种值类型，值不可变，字符串是字节的定长数组
	// 字符串中的字符根据需要占用1-4个字节，与Java、Python、C++不同，这样可以减少内存和硬盘空间占用，同时也不用像其他语言那样对使用UTF-8字符集文本进行编解码

	// Go支持 解释字符串 和 非解释字符串 两种形式字面值
	// .解释字符串 \n  \r  \t  \u(\U)  \\
	// .非解释字符串 `this is a raw string \n`，此时\n\被原样输出

	// 拼接字符串高效的方法 bytes.Buffer > strings.Join() > +

}