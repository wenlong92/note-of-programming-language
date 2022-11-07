// 「类型+方括号」是最简单的表示数组方法
let fibonacci: number[] = [1, 1, 2, 3, 5]
// 此时数组中不允许出现其他类型，数组的方法中例如push方法也不允许添加其他类型

// 接口也可以描述数组
interface NumberArray {
  [index: number]: number;   //只要索引类型是数字时，则值的类型必须是数字
}
let fibonacci1: NumberArray = [1, 1, 2, 3, 5]
// 一般不会用接口描述数组，太复杂，不过接口经常用来表示类数组


// 类数组不是数组类型
function sum() {
  let args: number[] = arguments  // arguments实际上是一个类数组，不能用普通数组的方式表示，而应该用接口
}
function sum1() {
  let args: {
    [index: number]: number;
    length: number;
    callee: Function;
  } = arguments
}


// 用any表示数组中允许出现任意类型
let list: any[] = ["123", 12, {website: "www.baidu.com"}]