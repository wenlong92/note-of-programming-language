package main
import (
	"io/ioutil"
	"regexp"
	"fmt"
)

func main () {
	// 7.6.1从字符串生成字节切片
	// s假设是一个字符串(本质上是一个字节数组)，则可以直接通过 c := []byte(s) 获取一个字节切片c
	// 也可以通过copy函数达到相同目的： copy(dst []byte, src string)

	// 还可以使用for-range获得每个元素
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c \n", i, c)
	}

	// 7.6.2获取字符串的某一部分
	// substr := str[start:end]

	// 7.6.3字符串和切片的内存结构
	// 内存中，一个字符串实际上是一个双字结构，即一个指向实际数据的指针和记录字符串长度的整数
	// 因为指针对用户不可见，我们可以将字符串看做是一个值类型，也就是一个字符数组

	// 7.6.4修改字符串中某个字符
	// Go语言中字符串不可变，因此必须将字符串转换成字节数组，再修改数组中的元素，最后将字节数组转换回字符串格式
	s1 := "hello"
	c1 := []byte(s1)
	c1[0] = 'c'
	s2 := string(c1)
	fmt.Println("s2:",s2,"\n")
	// 所以可以通过操作切片完成对字符串的操作

	// 7.6.5字节数组对比函数
	// 下例中Compare函数会返回两个字节数组字典顺序的整数对比结果

	// 7.6.6搜索及排序切片和数组
	// 标注库提供了sort包实现常见的搜索和排序操作。
	// 可使用sort中的函数 func Ints(a []int)实现对int类型的切片排序，例如：sort.Ints(arri),arri就是需要被升序排序的数组或切片
	// 为检查某个数组是否已被排序，可通过函数 IntsAreSorted(a []int)bool 检查，返回true表示已被排序

	// 可使用函数 func Float64s(a []float64) 排序float64的元素，
	// 或用函数 func Strings(a []string)排序字符串元素

	// 想在数组或切片中搜索一个元素，则该数组或切片必须先被排序(因为标准库中搜索算法使用的是二分法)
	// 然后再使用 func SearchInts(a []int, n int) int进行搜索，并返回对应结果的索引值
	// 类似的函数还有 func SearchFloat64s(a []float64, x float64) int 和 func SearchStrings(a []string, x string) int

	// 7.6.7append函数常见操作
	// 1.将切片b的元素追加到切片a之后， a = append(a, b...)
	// 2.删除位于索引 i 的元素：a = append(a[:i],a[i+1:]...)
	// 3.切除切片a中从索引i至j位置的元素：a = append(a[:i],a[j:])
	// 4.为切片a扩展j个元素长度：a = append(a, make([]T,j)...)
	// 5.在索引 i 的位置插入元素x：a = append(a[:i], append([]T{x},a[i:]...)...)
	// 6.在索引 i 的位置插入长度为j的新切片：a = append(a[:i], append(make([]T, j),a[i:]...)...)
	// 7.在索引 i 的位置插入切片 b 的所有元素：a = append(a[:i], append(b, a[i:]...)...)
	// 8.取出位于切片 a 最末尾的元素x：x, a = a[len(a)-1], a[:len(a)-1]
	// 9.将元素x追加到切片a：a = append(a,x)

	// 因此，可以使用切片和append操作来表示任意可变长度的序列
	// 从数学角度看，切片相当于向量，如果需要的话可以定义一个向量作为切片的别名进行操作

	// 7.6.8切片和垃圾回收
	// 切片的底层指向一个数组，该数组的实际容量可能要大于切片所定义的容量。
	// 只有在没有任何切片指向的是时候，底层的数组内存才会被释放，这种特性有时会导致程序占用多余内存
	// FindDigits函数将一个函数加载到内存，然后搜索其中所有数字并返回一个切片
	// 这段代码可顺利运行，但返回的[]byte指向的底层是整个文件数据，只要该返回的切片不被释放
	// 垃圾回收器就不能释放整个文件所占用的内存。也就是说，一点点有用的数据占用了整个文件的内存，例子：

	// 若要避免这个问题，见FindFileDigits函数
}
var digitRegexp = regexp.MustCompile("[0-9]+")
func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}
func FindFileDigits(filename string) []byte {
	fileBytes, _ := ioutil.ReadFile(filename)
	b := digitRegexp.FindAll(fileBytes, len(fileBytes))
	c := make([]byte, 0)
	for _, bytes := range b {
		c = append(c, bytes...)
	}
	return c
}

func Compare(a, b[]byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0
}