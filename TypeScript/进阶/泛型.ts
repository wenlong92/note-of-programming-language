// 泛型：在定义函数、接口或类的时候，不预先指定具体的类型，而在使用时再指定类型的一种特性
function createArray(length: number, value: any): Array<any> {
  let result = [];
  for (let i = 0; i < length; i++) {
    result[i] = value;
  }
  return result;
}
createArray(3, 'x')   // ['x', 'x', 'x']
// 上面代码编译不会报错，但缺陷是没有准确的定义返回值的类型 Array<any>允许数组每一项为任意类型，但预期是每一项都应该是输入的value类型
// 以下用泛型重写
function createArray1<T>(length: number, value: T): Array<T> {
  let result: T[] = [];
  for (let i = 0; i < length; i++) {
    result[i] = value;
  }
  return result;
}
createArray1<string>(3, 'x')   // ['x', 'x', 'x']
createArray1(3, 'x')   // 也可以不用手写，让类型推论自动推算出来



// 多个类型参数
// 定义泛型时，可以一次定义多个类型参数
function swap<T, U>(tuple: [T, U]): [U, T] {
  return [tuple[1], tuple[0]]
}
swap([7, 'seven'])  // ['seven', 7]


// 泛型约束
// 函数内部使用泛型变量时，由于事先不知道是哪种类型，所以不能随意操作它的属性或方法，例如：
function loggingIdentity<T>(arg: T): T {
  console.log(arg.length)
  return arg
}
// 上例中泛型T不一定包含属性length，所以编译时报错

// 这时可以对泛型进行约束，只允许这个函数传入那些包含length属性的变量，这就是泛型约束
interface Lengthwise {
  length: number;
}
function loggingIdentity1<T extends Lengthwise>(arg: T): T {
  console.log(arg.length)
  return arg
}
loggingIdentity1(7)  // 调用时，arg不包含length属性，则会报错


// 多个类型参数之间也可以相互约束：
function copyFields<T extends U, U>(target: T, source: U): T {
  for (let id in source) {
    target[id] = (<T>source)[id]
  }
  return target;
}
let x = { a: 1, b: 2, c: 3, d: 4 }
console.log(copyFields(x, {b: 10, d: 20}))   // 其中要求T继承U，保证了U上不会出现T中不存在的字段


// 泛型接口
// 之前章节中，可以使用接口的方式定义一个函数需要符合的形状，例如：
interface SearchFunc {
  (source: string, subString: string): boolean;
}
let mySearch: SearchFunc;
mySearch = function(source: string, subString: string) {
  return source.search(subString) !== -1
}
// 也可以使用含有泛型的接口定义函数的形状
interface CreateArrayFunc {
  <T>(length: number, value: T): Array<T>;
}
let createArray2: CreateArrayFunc;
// let createArray2: CreateArrayFunc<any>;   // 也可以把泛型参数提前到接口名上
createArray2 = function<T>(length: number, value: T): Array<T> {
  let result: T[] = [];
  for (let i = 0; i < length; i++) {
    result[i] = value;
  }
  return result;
}
createArray2(3, 'x')


// 泛型类：与泛型接口类似，泛型也可以用于类的类型定义中：
class GenericNumber<T> {
  zeroValue: T;
  add: (x: T, y: T) => T;
}
let myGenericNumber = new GenericNumber<number>()
myGenericNumber.zeroValue = 0;
myGenericNumber.add = function(x, y) { return x + y; };


// 泛型参数的默认类型
// 在ts2.3以后，可以为泛型中的类型参数指定默认类型。
// 当使用泛型时，没有在代码中直接指定类型参数，从实际值参数中无法推测出时， 这个默认类型就会其作用
function createArray3<T = string>(length: number, value: T): Array<T> {
  let result: T[] = [];
  for (let i = 0; i < length; i++) {
    result[i] = value;
  }
  return  result
}


