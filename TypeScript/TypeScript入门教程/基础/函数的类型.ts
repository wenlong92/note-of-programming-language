// JavaScript中有两种常见的定义函数的方式：函数声明和函数表达式

// 函数声明
function sum(x, y) {
  return x + y
}
// 函数表达式
let mySum = function(x, y) {
  return x + y
}

// 在ts中，要对输入和输出进行约束
function sum1(x: number, y: number): number {
  return x + y
}

let mySum1:(x: number, y: number) => number = function(x: number, y: number): number {
  return x + y
}
// ts中 =>和 ES6中不同，ts中=>用来表示函数的定义，左边是输入类型，需要括号括起来，右边是输出类型
// ES6中=>叫做箭头函数


// 也可以用接口定义函数的形状
interface SearchFunc {
  (source: string, subString: string): boolean;
}
let mySearch: SearchFunc;
mySearch = function(source: string, subString: string) {
  return source.search(subString) !== -1
}
// 采用函数表达式或接口定义函数方式时，对等号左侧进行类型限制，保证以后对函数名赋值时保证参数个数、参数类型、返回值类型不变


// 可选参数,与接口中可选参数类似，用?表示可选参数
// 可选参数接在必须参数后，也就是说可选参数后不能出现必须参数
function buildName(firstName: string, lastName?: string) {
  if(lastName) {
    return firstName + ' ' + lastName
  } else {
    return firstName
  }
}
let tomcat = buildName("Tom", 'Cat')
let tom = buildName("Tom")


// 参数默认值
function buildName1(firstName: string, lastName: string = 'Cat') {
  return firstName + " " + lastName
}


// 剩余参数
// ES6中使用 ...rest  方式获取函数中剩余参数
function push(array, ...items) {
  items.forEach(function(item) {
    array.push(item)
  })
}
let a: any[] = [];
push(a, 1, 2, 3)
// ts中
function push1(array: any[], ...items: any[]) {
  items.forEach(function(item) {
    array.push(item)
  })
}



// 重载：允许一个函数接受不同数量或类型的参数时，做出不同的处理
// 例如，接受数字123时，输出321，输入'hello'时，输出'olleh'
// function reverse(x: number | string): number | string | void {
//   if(typeof x === 'number') {
//     return Number(x.toString().split("").reverse().join(""))
//   } else if(typeof x === 'string') {
//     return x.split("").reverse().join("")
//   }
// }
// console.log(reverse(123))

// 然而这样有一个缺点，就是不能够精确的表达，输入为数字的时候，输出也应该为数字，输入为字符串的时候，输出也应该为字符串。
// 这时，我们可以使用重载定义多个 reverse 的函数类型：
function reverse(x: number): number;
function reverse(x: string): string;
function reverse(x: number | string): number | string | void {
    if (typeof x === 'number') {
        return Number(x.toString().split('').reverse().join(''));
    } else if (typeof x === 'string') {
        return x.split('').reverse().join('');
    }
}
// 上例中，我们重复定义了多次函数 reverse，前几次都是函数定义，最后一次是函数实现。在编辑器的代码提示中，可以正确的看到前两个提示。
// 注意，TypeScript 会优先从最前面的函数定义开始匹配，所以多个函数定义如果有包含关系，需要优先把精确的定义写在前面。