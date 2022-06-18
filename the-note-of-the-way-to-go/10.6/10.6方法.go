package main
import (
	"fmt"
	"time"
	"math"
)

type TwoInts struct {
	a int
	b int
}

type IntVector []int

type employee struct {
	salary float32
}

type myTime struct {
	time.Time // anonymous field
}

type B struct {
	thing int
}

type List []int

type Point struct {
	x, y float64
}

type NamedPoint struct {
	Point
	name string
}

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log *Log
}

type Camera struct {}

func (c *Camera) TakeApicture() string {
	return "Click"
}

type Phone struct {}

func (p *Phone) Call() string {
	return "Ring Ring"
}

type CameraPhone struct {
	Camera
	Phone
}

func main() {
	// 10.6.1方法是什么
	// Go语言中，结构体就像是类的一种简化形式，那么类的方法是什么？在Go中有一个概念，它和方法有着同样的名字
	// 且大体意思相同：Go方法是作用在接受者上的一个函数，接受者是某种类型的变量。因此方法是一种特殊类型的函数。
	
	// 接受者类型几乎可以是任何类型，不仅仅是结构体类型：任何类型都可以有方法，甚至是函数类型，可以是int、bool、string或数组的别名类型
	// 但接受者不能是一个接口类型，因为接口类型是一个抽象定义，但是方法确实具体实现，这样做会引发编译错误

	// 最后接受者不能使一个指针类型，但是它可以是任何其他允许类型的指针。

	// 一个类型加上它的方法等价于面向对象中的一个类。
	// 区别在于：Go中，类型的代码和绑定在它上面的方法代码可以不放置在一起，可以存在不同的源文件，唯一的要求是：他们必须是同一个包的。

	// 类型T(或*T)上所有方法的集合叫做类型T(*T)的方法集

	// 因为方法是函数，所以不允许方法重载，即对于一个类型只能有一个给定名称的方法。
	// 但是如果基于接受者类型，是有重载的：具有同样名字的方法可以在2个或多个不同的接受者类型上存在，例如：
	// func (a *deseMatrix) Add(b Matrix) Matrix
	// func (a *sparseMatrix) Add(b Matrix) Matrix

	// 别名类型没有原始类型上已经定义过的方法

	// 定义方法的一般格式如下：
	// func (recv receiver_type) methodName(parameter_list) (return_value_list) {...}

	// 方法名之前，func关键字之后的括号中指定receiver
	// 如果recv是receiver的实例，Method1是它的方法名，那么方法调用遵循传统的object.name
	// 如果recv是一个指针，Go会自动解引用
	// 如果方法不需要使用recv的值，可用_替换，例如：
	// func (_ receiver_type) methodName(parameter_list) (return_value_list) {...}

	// recv像是面向对象语言中的this或self，但Go中并没有这两个关键字
	// 下面是一个结构体上简单方法的例子：
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n",two2.AddThem())

	// 非结构体类型上方法的例子：
	fmt.Println(IntVector{1, 2, 3}.Sum())  //输出为6

	// 定义结构体employee，有一个salary字段，给这个结构体定义一个方法fiveRaise，按照指定的百分比增加薪水
	em1 := employee{1000}
	fmt.Println("Salary raise: %f\n", em1.giveRaise(0.5))

	// 类型和作用在它上面定义的方法必须在同一个包里定义，这就是为什么不能在int、float或类似这些的类型上定义方法
	// 但有一个间接的方式：可以先定义该类型的别名类型，然后再为别名类型定义方法
	// 或者像下面这样将它作为匿名类型嵌入在一个新的结构体中:
	m := myTime{time.Now()}
	// 调用匿名Time上的String方法
	fmt.Println("Full time now:", m.String())
	// 调用myTime.first3Chars
	fmt.Println("First 3 chars:", m.first3Chars())


	// 10.6.2函数和方法的区别
	// 函数将变量作为参数：function1(recv)
	// 方法在变量上被调用：recv.Method1()

	// 在接受者是指针时，方法可以改变接受者的值(或状态)，这点函数也可以做到(当参数作为指针传递，即通过引用调用时，函数也可以改变参数的状态)

	// 接受者必须有一个显示的名字，这个名字必须在方法中被使用
	// receiver_type叫做(接受者)基本类型，这个类型必须在和方法同样的包中被声明

	// Go中，(接受者)类型关联的方法不写在类型结构里面，就像类那样；耦合更加宽松；类型和方法之间的关联由接受者建立
	
	// 方法没有和数据定义(结构体)混在一起：他们是正交的类型；表示(数据)和行为(方法)是独立的。

	
	// 10.6.3指针或值作为接受者
	// 鉴于性能原因，recv最常见的是一个指向receiver_type的指针(因为不想要一个实例的拷贝，如果按值调用的话就会是这样)
	// 特别是在receiver类型是结构体时

	// 如果想要方法改变接受者的数据，就在接受者的指针类型上定义改方法。否则就在普通的值类型上定义方法。
	// 下例做了说明，change()接受一个指向B的指针，并改变它内部的成员；write()通过拷贝接受B的值并只输出B的内容
	// Go为我们做了探测工作，我们自己并没有指出是否在指针上调用方法，Go替我们做了这些事情。
	// b1是值而b2是指针，方法都支持运行了。
	var b1 B     // b1是值
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) //b2是指针
	b2.change()
	fmt.Println(b2.write())

	// 指针方法和值方法都可以在指针或非指针上被调用，如下所示，类型List在值上有一个方法Len()，在指针上有一个方法Append()
	// 但可以看到两个方法都可以在两种类型的变量上被调用
	// 值
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)", lst, lst.Len())
	// 指针
	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v(len:%d)", plst, plst.Len())


	// 10.6.4方法和未导出字段
	// 一个包中的类型被明确导出，但它的字段却没被导出。如何在另一个程序中修改或者读取类型的字段？
	// 在面向对象语言中，提供了getter和setter方法。在Go中setter方法使用Set前缀，对于getter方法只使用成员名

	// 并发访问对象
	// 对象的字段(属性)不应该由2个或以上的不同线程在同一时间去改变。
	// 为安全并发访问，可使用sync包中的方法(9.3节)。14.17节会通过goroutines和channels探索另一种方式。


	// 10.6.5内嵌类型的方法和继承
	// 当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，这在效果上等同于外层类型继承了这些方法：
	// 将父类型在子类型中来实现亚型。
	// 这个机制提供了一种简单的方式模拟经典面向对象语言中的子类和继承相关的效果，类似Ruby中的混入(mixin)，下例：
	n := &NamedPoint{Point{3, 4}, "Pythoagoras"}
	fmt.Println("\n",n.Abs())
	// 可以覆写方法(像字段一样)：和内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法

	// 结构体内嵌和自己在同一个包中的结构体时，可以彼此访问对方所有的字段和方法


	// 10.6.6如何在类型中嵌入功能
	// 主要有两种方法来实现在类型中嵌入功能：
	// A.聚合(或组合)：包含一个所需功能类型的具名字段
	// B.内嵌：内嵌(匿名地)所需功能类型，例如10.6.5演示
	// 例子：为了使这些概念具体化，假设有一个Customer类型，让它通过Log类型来包含日志功能，Log类型只是简单地包含一个累积的消息
	// 如果想染特定类型都具备日志功能，可以实现一个这样的Log类型，然后将它作为特定类型的字段，并提供Log()，它返回这个日志的引用。
	// 方式A实现：
	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	// shorter
	c = &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	c.Log().Add("2 - After me the world will be a better place!")
	fmt.Println(c.Log())
	// 方式B实现：


	// 10.6.7多重继承
	// 多重继承指的是类型获得多个父类型行为的能力，在传统的面向对象语言中通常是不被实现的(C++和Python除外)
	// 在类继承层次中，多重继承会给编译器引入额外的复杂度。
	// 但在Go中，通过在类型中嵌入所有必要的父类型，可很简单的实现多重继承。

	// 下例中，假设有一个类型CameraPhone，通过它可以Call()，也可以TakeApicture()，但是第一个方法属于类型Phone，第二个方法属于类型Camera
	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone exhibits multiple behaviors...")
	fmt.Println("It exhibits behavior of a Camera: ", cp.TakeApicture())
	fmt.Println("It works like a Phone too: ", cp.Call())


	// 10.6.8通用方法和方法命名
	// 在编程中一些基本操作会一遍又一遍出现，例如打开、关闭、读、写、排序等，
	// 并且它们都有一个大致的意思：打开可以作用与一个文件、一个网络连接、一个数据库链接等等。具体实现可能千差万别，但基本概念一致。
	// 在Go中，通过使用接口，标准库广泛的应用这些规则，在标准库中这些通用方法都有一致的名字，例如：Open()/Read()/Write()等
	// 想写规范的Go程序，要遵守这些约定，给方法合适的名字和签名


	// 10.6.9和其他面向对象语言比较Go的类型和方法
	// 在C++、Java、C#、Ruby等面向对象语言中，方法在类的上下文中被定义和继承：
	// 在一个对象上调用方法时，运行时会检测类以及它的超类中是否由此方法的定义，如果没有会导致异常。

	// Go语言中，这样的继承层次是完全没必要的：如果方法在此类型定义了，就可以调用，和其他类型上是否存在这个方法没有关系。

	// Go中，类型就是类(数据和关联的方法)。Go不知道类似面向对象语言的类继承的概念。继承有两个好处：代码复用和多态

	// Go中，代码复用通过组合和委托实现，多态通过接口的使用来实现：有时这也叫组件编程
	
}
func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return s
}

func (em employee) giveRaise(a float32) float32 {
	return em.salary * (1 + a)
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func (b *B) change() { b.thing = 1 }
func (b B) write() string { return fmt.Sprint(b) }

func(l List) Len() int {return len(l)}
func(l *List) Append(val int) { *l = append(*l, val)}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}