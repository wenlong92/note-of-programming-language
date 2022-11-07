// 枚举类型用于取值被限定在一定范围内的场景，例如一周只能有七天，颜色限定为红绿蓝等
// 枚举使用enum关键字来定义
enum Days {Sun, Mon, Tue, Wed, Thu, Fri, Sat}

// 枚举成员会被赋值为从 0 开始递增的数字，同时也会对枚举值到枚举名进行反向映射
console.log(Days['Sun'] === 0)  // true
console.log(Days['Mon'] === 1)  // true
console.log(Days['Tue'] === 2)  // true
console.log(Days['Wed'] === 3)  // true
console.log(Days['Thu'] === 4)  // true
console.log(Days['Fri'] === 5)  // true
console.log(Days['Sat'] === 6)  // true

console.log(Days[0] === 'Sun')  // true


// 也可以给枚举项手动赋值
enum Days1 {Sun = 7, Mon = 1, Tue, Wed, Thu, Fri, Sat}
console.log(Days1['Sun'] === 7)  // true
console.log(Days1['Mon'] === 1)  // true
console.log(Days1['Tue'] === 2)  // true
console.log(Days1['Wed'] === 3)  // true
// 未手动赋值的枚举项会接着上一个枚举项递增


// 如果未手动赋值的枚举项与手动赋值的重复，ts是不会察觉的，最新的会覆盖前面的值



// 手动枚举的项可以不是数字，此时需要使用类型断言来让tsc无视类型检查
enum Days2 {Sun = 7, Mon, Tue, Wed, Thu, Fri, Sat = <any>"S"}

// 手动赋值的枚举值也可以为小数或负数，后续未手动赋值的项递增步长仍为1
enum Days3 {Sun = 7, Mon = 1.5, Tue, Wed, Thu, Fri, Sat}
console.log(Days3['Tue'] === 2.5) // true

enum Days4 {Sun = 7, Mon = -1, Tue, Wed, Thu, Fri, Sat}
console.log(Days4['Tue'] === 0) // true




// 常数项和计算所得项
// 枚举项有两种类型：常数项和计算所得项
// 上面的例子是常数项，下面举一个计算所得项的例子
enum Color {Red, Green, Blue = 'blue'.length}   // 'blue'.length就是一个计算所得项
console.log(Color['Blue'])  // 4

// 如果紧接在计算所得项后面的是未手动赋值的项，那么它就会因为无法获得初始值而报错
// enum Color1 {Red = 'red'.length, Green, Blue}   // 报错


// 常数枚举：使用const enum定义的枚举类型
const enum Directions {
  Up,
  Down,
  Left,
  Right
}
let directions = [Directions.Up, Directions.Down, Directions.Left, Directions.Right]
// 常数枚举与普通枚举的区别在于，它会在编译阶段被删除，并且不能包含计算成员



// 外部枚举：使用declare enum定义的枚举类型
declare enum Directions1 {
  Up,
  Down,
  Left,
  Right
}
let directions1 = [Directions1.Up, Directions1.Down, Directions1.Left, Directions1.Right]