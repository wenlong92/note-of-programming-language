package main
import (
	"fmt"
	"strings"
)

func main() {
	// 每种语言都有一些对字符串的预定义处理函数。Go中使用strings包对字符串进行主要操作

	// 前缀和后缀 strings.HasPrefix()  strings.HasSuffix()
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Dose the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("\n%t\n",strings.HasPrefix(str, "Th"))

	// strings.Contains(s, substr string)bool 判断是否包含substr
	var str1 string = "This is the demo of  strings.Contains"
	fmt.Printf("\n%t",strings.Contains(str1, "demo"))

	// 判断子字符串或字符在父字符串中出现的位置索引,
	// strings.Index(s, str string) int (str的第一个字符索引) ,   strings.LastIndex(s, str string) int (str的第一个字符索引)
	var str2 string = "This is the demo of string index"
	fmt.Printf("\n%d",strings.Index(str2, "is"))
	fmt.Printf("\n%d",strings.Index(str2, "h"))
	fmt.Printf("\n%d",strings.LastIndex(str2, "in"))
	fmt.Printf("\n%d",strings.LastIndex(str2, "d"))

	// 字符串替换 strings.Replace(str, old, new string, n int) string
	var str3 string = "replace string string string string"
	fmt.Printf("\n%f",strings.Replace(str3, "string", "int", 3)) //替换前三个string为int
	fmt.Printf("\n%f",strings.Replace(str3, "string", "all", -1)) //替换所有string为all

	// 统计字符串str在字符串s中出现的非重叠次数  strings.Count(s,str string) int
	var str4 string = "Hello, how is it going, Hugo?"
	var manyG string = "gggggggggg"
	fmt.Printf("\n%d", strings.Count(str4, "H"))
	fmt.Printf("\n%d", strings.Count(manyG, "gg"))

	// 重复 count 次字符串 s 并返回一个新的字符串 strings.Repeat(s, count int) string
	var str5 string = "Hi there! "
	str6 := strings.Repeat(str5, 3)
	fmt.Printf("\n%f", str6)

	// 修改字符串大小写  strings.ToLower(s) string   strings.ToUpper(s) string
	var str7 string = "Hey, how are you George?"
	fmt.Printf("\n%f", strings.ToLower(str7))
	fmt.Printf("\n%f", strings.ToUpper(str7))

	// 修剪字符串 strings.TrimSpace(s)剔除字符串开头结尾的空白符号   strings.Trim(s,"cut")剔除字符串开头结尾的cut
	// TrimLeft剔除开头字符串  TrimRight剔除结尾字符串
	var str8 string = "  Hello world  "
	str9 := "hey bill hey"
	str10 := "hi jack hi"
	fmt.Printf("\n%f", strings.TrimSpace(str8))
	fmt.Printf("\n%f", strings.Trim(str9, "hey"))
	fmt.Printf("\n%f", strings.TrimLeft(str10, "hi"))
	fmt.Printf("\n%f", strings.TrimRight(str10, "hi"))

	// 分割字符串
	// strings.Fields(s)利用一个或多个空白符作为动态长度的分隔符将字符串分割成若干小块，返回一个slice
	// 如果字符串只包含空白符号，返回一个长度为0的slice
	// strings.Split(s, sep)自定义分割符号对指定字符串进行分割，同样返回slice
	// 因为这两个函数都返回slice，所以习惯使用for-range循环对其进行处理
	var str11  = "space space space space"
	str11sub := strings.Fields(str11)
	fmt.Printf("\n%f",str11sub)

	var str12 = "sss1sss1sss1sss1"
	str12sub := strings.Split(str12,"1")
	fmt.Printf("\n%f",str12sub)

	// 拼接slice到字符串 strings.Join(sl []string, sep string)string
	str13 := strings.Join(str12sub, ";")
	fmt.Printf("\n%f", str13)

	// strings.NewReader(str),用于生成一个Reader并读取字符串中的内容，返回指向改Reader的指针
	// Read()从[]byte中读取内容
	// ReadByte()和ReadRune()从字符串中读取下一个byte或者rune
	var str14 = "hello world"
	str15 := strings.NewReader(str14)
	fmt.Printf("\n%f",str15)

	// 与字符串相关的类型转换都是通过strconv包实现的



}