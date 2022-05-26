package main
import (
	"fmt"
	_ "strings"
	_ "strconv"
	_ "time"
	_ "sort"
)
// func test(n int) {
// 	if n > 2 {
// 		n--
// 		test(n)
// 	}
// 	fmt.Println("n = ",n)
// }
// func demo(n1 int, n2 int) (sum int, sub int) {
// 	sub = n1 - n2
// 	sum = n1 + n2
// 	return sum, sub
// }
// var age = test()
// func test() int {
// 	fmt.Println("test")
// 	return 90
// }
// func init() {
// 	fmt.Println("init")
// }
// func makeSuffix(suffix string) func(string) string {
// 	return func (name string) string {
// 		if !strings.HasSuffix(name, suffix) {
// 			return name + suffix
// 		}
// 		return name
// 	}
// }

// func sum(n1 int, n2 int) int {
// 	defer fmt.Println("n1 = ",n1)
// 	defer fmt.Println("n2 = ",n2)

// 	res := n1 + n2
// 	fmt.Println("3 res = ",res)
// 	return res
// }
// type Person struct {
// 	Name string
// }
// func (p Person) test() {
// 	fmt.Println("test() name =", p.Name)
// }

type Circle struct {
	radius float64
}
func (c Circle) area() float64 {
	res := 3.14 * c.radius * c.radius
	fmt.Println(res)
	return res
}
func main() {
	var c1 Circle
	c1.radius = 3.45
	c1.area()



	// var p1 Person
	// p1.Name = "tom"
	// p1.test()


	// type Person struct {
	// 	Name string
	// 	Color string
	// }
	// // person := Person{"white","bill"}
	// // fmt.Println(person.Name)

	// var person *Person = new(Person)
	// fmt.Println(person)


	// type Cat struct {
	// 	Name string
	// 	Age int
	// 	Color string
	// 	Hobby string
	// }
	// var cat1 Cat
	// fmt.Printf("%p\n",&cat1)
	// cat1.Name = "bill"
	// cat1.Age = 12
	// cat1.Color = "white"
	// cat1.Hobby = "play"
	// fmt.Println(cat1)



	// listMap := make(map[int]string)
	// listMap[10] = "90"
	// listMap[100] = "190"
	// listMap[3] = "56"
	// listMap[1] = "900"

	// var keys []int
	// for k,_ := range listMap {
	// 	keys = append(keys, k)
	// }
	// sort.Ints(keys)
	// fmt.Println(keys)

	// for _, k := range keys {
	// 	fmt.Println(listMap[k])
	// }


	// var monsters = make([]map[string]string, 3)
	// if monsters[0] == nil {
	// 	monsters[0] = make(map[string]string,2)
	// 	monsters[0]["name"] = "牛魔王"
	// 	monsters[0]["age"] = "500"
	// }
	// fmt.Println(monsters)


	// cities := map[string]string {
	// 	"no1" : "peking",
	// 	"no2" : "shanghai",
	// 	"no3" : "shenzhen",
	// }
	// for k, v := range cities {
	// 	fmt.Println(k, v)
	// }

	// stuMap := make(map[string]map[string]string)
	// stuMap["stu1"] = make(map[string]string,3)
	// stuMap["stu1"]["name"] = "tom"
	// stuMap["stu1"]["sex"] = "male"
	// stuMap["stu1"]["age"] = "16"

	// stuMap["stu2"] = make(map[string]string,3)
	// stuMap["stu2"]["name"] = "mary"
	// stuMap["stu2"]["sex"] = "female"
	// stuMap["stu2"]["age"] = "15"
	// fmt.Println(len(stuMap))

	// for k1, v1 := range stuMap {
	// 	fmt.Println(k1, v1)
	// 	for k2, v2 := range v1 {
	// 		fmt.Println(k2, v2)
	// 	}
	// }


	// cities := make(map[string]string)
	// cities["no1"] = "peking"
	// cities["no2"] = "shenzhen"
	// cities["no3"] = "guangzhou"
	// fmt.Println(cities["no1"])

	// cities["no1"] = "shanghai"
	// fmt.Println(cities["no1"])

	// delete(cities, "no1")
	// fmt.Println(cities)

	// // cities = make(map[string]string)
	// // fmt.Println(cities)

	// val, ok := cities["no2"]
	// fmt.Println("ok",ok)
	// if ok {
	// 	fmt.Println(val)
	// } else {
	// 	fmt.Println("no")
	// }


	// info := make(map[string]map[string]string) 
	// info["stu1"] = make(map[string]string, 3)
	// info["stu1"]["name"] = "tom"
	// info["stu1"]["sex"] = "male"
	// info["stu1"]["location"] = "peking"

	// info["stu2"] = make(map[string]string, 3)
	// info["stu2"]["name"] = "bill"
	// info["stu2"]["sex"] = "male"
	// info["stu2"]["location"] = "shanghai"

	// info["stu3"] = make(map[string]string, 3)
	// info["stu3"]["name"] = "mary"
	// info["stu3"]["sex"] = "female"
	// info["stu3"]["location"] = "shenzhen"
	// fmt.Println(info)
	// fmt.Println(info["stu1"])
	// fmt.Println(info["stu1"]["name"])


	// cities := make(map[string]string)
	// cities["no1"] = "shenyang"
	// cities["no2"] = "shenzhen"
	// fmt.Println(cities)

	// heros := map[string]string {
	// 	"hero1" : "bill",
	// 	"hero2" : "bill2",
	// }
	// fmt.Println(heros)

	// var a map[string]string
	// a = make(map[string]string, 1)
	// a["no1"] = "123"
	// a["no2"] = "1234"
	// a["no3"] = "1235"
	// a["no4"] = "1235"
	// fmt.Println(a)
	// fmt.Println(len(a))


	// var slice []float64 = make([]float64, 5, 10)
	// slice[1] = 10
	// slice[3] = 20
	// fmt.Println(slice)
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	// persons := [...]string{"bill", "jack", "mary"}
	// for index, value := range persons {
	// 	fmt.Printf("%v, %v\n",index, value)
	// 	fmt.Println(index, persons[index])
	// }

	// var numArr1 [3]int = [3]int{1,2,3}
	// fmt.Println("numArr1 = ", numArr1)

	// var numArr2  = [3]int{1,2,3}
	// fmt.Println("numArr2 = ",numArr2)

	// var numArr3 = [...]int{3,4,5,5}
	// fmt.Println("numArr3 = ",numArr3)

	// var numArr4 = [...]int{0: 200,2: 300,1: 400}
	// fmt.Println("numArr4 = ",numArr4)

	// var intArr [3]int
	// fmt.Println(intArr)
	// intArr[0] = 10
	// intArr[1] = 20
	// intArr[2] = 30
	// fmt.Println(intArr)
	// fmt.Printf("地址%p  [0]地址%p  [1]地址%p  [2]地址%p",&intArr, &intArr[0], &intArr[1], &intArr[2])
	// num1 := 100
	// fmt.Printf("num1类型%T, num1值%v, num1的地址%v\n", num1, num1, &num1)

	// num2 := new(int)
	// // *num2 = 100
	// fmt.Printf("num2类型%T, num2值%v, num2的地址%v\n, num2指向的值%v", num2, num2, &num2, *num2)
	// now := time.Now()
	// fmt.Printf("%v\n",now.Minute())

	// n, err := strconv.Atoi("12")
	// if err != nil {
	// 	fmt.Println("err: ",err)
	// } else {
	// 	fmt.Println(n)
	// }

	// str := strconv.Itoa(12333)
	// fmt.Println(str)

	// bytes := []byte("hello go")
	// fmt.Printf("bytes: %v\n",bytes)

	// str := "hello北京"
	// r := []rune(str)
	// for i := 0; i < len(r); i++ {
	// 	fmt.Printf("%c\n",r[i])
	// }
	// str := "hello被"
	// fmt.Println("len = ",len(str))

	// res := sum(10, 20)
	// fmt.Println("res = ", res)

	// f2 := makeSuffix(".jpg1")
	// fmt.Println(f2("winter"))
	// fmt.Println("main",age)
	// a1, b1 := demo(2, 1)
	// fmt.Println(a1, b1)

	// test(4)

	// for i := 1; i <= 20; i++ {
	// 	if i == 3 {
	// 		return
	// 	}
	// 	fmt.Println("i ", i)
	// }
	// fmt.Println("hello")

	// sum := 0
	// for i := 1; i <= 100; i++ {
	// 	sum += i
	// 	if sum > 20 {
	// 		fmt.Println("i ", i)
	// 		break
	// 	}
	// }

	// lable2:
	// for i := 2; i < 4; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if j == 2 {
	// 			break lable2
	// 		}
	// 		fmt.Println("j = ",j)
	// 	}
	// }

	// var n1 int32 = 5
	// var n2 int32 = 20
	// switch n1 {
	// 	case n2, 10, 5 :
	// 		fmt.Println("ok")
	// 		fallthrough
	// 	case 30 :
	// 		fmt.Println("123")
	// 		// fallthrough
	// 	default :
	// 		fmt.Println("wu")
	// }

	// var a int = 1 >> 2
	// var b int = -2 >> 3
	// var c int = 1 << 2
	// var d int = -1 << 2
	// fmt.Println("a = ",a)
	// fmt.Println("b = ",b)
	// fmt.Println("c = ",c)
	// fmt.Println("d = ",d)

	// var name string
	// var age byte
	// var sal float32
	// var isPass bool

	// fmt.Println("请输入姓名：")
	// fmt.Scanln(&name)

	// fmt.Println("请输入年龄：")
	// fmt.Scanln(&age)

	// fmt.Println("请输入薪水：")
	// fmt.Scanln(&sal)

	// fmt.Println("请输入是否通过考试：")
	// fmt.Scanln(&isPass)

	// fmt.Printf("名字：%v\n 年龄：%v\n 薪水： %v\n 是否通过： %v", name, age, sal, isPass)
}