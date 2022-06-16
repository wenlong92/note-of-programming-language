package main
import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool    "An important answer"
	field2 string  "The name of the thing"
	field3 int     "How much there are"
}

func main() {
	// 结构体中的字段除了名字和类型之外，还有一个可选的标签tag
	// tan是一个附属于字段的字符串，可以是文档或其他的重要标记
	// 标签的内容不可以在一般的编程中使用，只有reflect包能获取它。
	// 它可以在运行时自省类型、属性和方法，比如：
	// 在一个变量上调用reflect.TypeOf()可以获取变量的正确类型，
	// 如果变量是一个结构体类型，可通过Field索引结构体的字段，然后就可以使用Tag属性

	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}