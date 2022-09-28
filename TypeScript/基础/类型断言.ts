// 类型断言：手动指定一个值的类型
// 语法：值 as 类型   或者 <类型>值     tsx语法中必须使用前者
// 建议统一使用前者

// 应用场景：将一个联合类型断言为其中一个类型
interface Cat {
  name: string;
  run(): void;
}
interface Fish {
  name: string;
  swim: void;
}
function getName(animal: Cat | Fish) {
  if(typeof animal.swim === "function") {    // 由于Cat中没有swim方法，所以会报错
    return true
  }
  return false
}
// 此时可以用断言
function getName1(animal: Cat | Fish) {
  if(typeof (animal as Fish) === "function") {  // 将animal断言成Fish可以解决报错，但不能避免运行时的错误，所以使用时要慎重
    return true
  }
  return false
}



// 应用场景：将一个父类断言为更加具体的子类
class ApiError extends Error {
  code: number = 0;
}
class HttpError extends Error {
  statusCode: number = 200
}
function isApiError(error: Error) {
  if(typeof (error as ApiError).code === "number") {
    return true
  }
  return false
}
// 函数isApiError用来判断传入的参数是否为ApiError类型



// 应用场景：将任何一个类型断言为any
// 当我们确定这段代码不会出错但ts编译时会报错，这是我们可以使用  as any 临时将其断言为any类型，例如
// window.foo = 1   //在window上添加一个foo属性，但ts编译时报错，我们可以使用下面方法：
(window as any).foo = 1   //在any类型变量上，访问任何属性都是允许的
// 如果不是非常确定，不要使用 as any



// 应用场景：将any断言为一个具体的类型
// 在日常开发中，由于各种原因，不可避免的要处理any类型变量
// 例如，历史遗留代码中有个getCacheData函数，返回值是any
function getCacheData(key: string): any {
  return (window as any).cache[key]
}
// 在使用它时，最好能够将调用它之后的返回值断言成一个精确的类型，方便后续操作
interface Cat {
  name: string;
  run(): void;
}
const tom1 = getCacheData('tom') as Cat;
tom1.run()
console.log('tom1:',tom1)