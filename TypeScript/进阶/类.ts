// ts除了实现所有ES6中的类的功能以外，还添加了一些新的用法

// ts中使用三种访问修饰符：public、private、protected

/*
public：修饰的属性或方法是公有的，可以在任何地方被访问到，默认所有的属性和方法都是public
private：修饰的属性或方法是私有的，不能在声明它的类的外部访问
protected：修饰的属性或方法是受保护的，和private类似，区别在于它在子类中也是允许被访问的
*/

class Animal {
  public name;
  public constructor(name) {
    this.name = name;
  }
}
let a = new Animal('jack')
console.log(a.name)  // jack
a.name = 'Tom'
console.log(a.name)  // Tom

// 如果希望有些属性无法直接存取，可以用private
class Animal1 {
  private name;
  public constructor(name) {
    this.name = name
  }
}
// let a1 = new Animal1('jack')
// console.log(a1.name)  // 报错，但ts编译之后的代码中，并没有限制private属性在外部的可访问性
// 使用private修饰的属性或方法，在子类中不允许访问
class Cat extends Animal1{
  constructor(name) {
    super(name)
    console.log(this.name)  // 报错
  }
}

// 如果是protected修饰，则允许在子类中访问
class Animal2 {
  protected name;
  public constructor(name) {
    this.name = name
  }
}
class Cat1 extends Animal2 {
  constructor(name) {
    super(name)
    console.log(this.name)
  }
}

// 当构造函数修饰为private时，该类不允许被继承或者实例化
class Animal3 {
  public name
  private constructor(name) {
    this.name = name
  }
}
class Cat2 extends Animal3 {  // 报错，不允许继承
  constructor(name) {
    super(name)
  }
}

// 当构造函数修饰为protected时，该类只允许被继承
class Animal4 {
  public name
  protected constructor(name) {
    this.name = name
  }
}
class Cat3 extends Animal4 {
  constructor(name) {
    super(name)
  }
}
let a2 = new Animal4('Jack')  // 报错，该类只允许被继承



// 参数属性：修饰符和readonly还可以使用在构造函数参数中，等同于类中定义该属性同时给该属性赋值，使代码更简洁
class Animal5 {
  public constructor(public name) {
    // do something
  }
}

// readonly：只读属性关键字，只允许出现在属性声明或索引签名或构造函数中
class Animal6 {
  readonly name;
  public constructor(name) {
    this.name = name
  }
}
let a3 = new Animal6('Jack')
console.log(a3.name)  // Jack
a3.name = 'Tom'       // 报错

// 如果readonly和其他访问修饰符同时存在，需要写在其后面
class Animal7 {
  public constructor(public readonly name) {
    this.name = name
  }
}

// 抽象类：abstract用于定义抽象类和其中的抽象方法
// 抽象类不允许被实例化
abstract class Animal8 {
  public name
  public constructor(name) {
    this.name = name
  }
  public abstract sayHi()
}
let a4 = new Animal8('Jack')  // 报错，不能创建抽象类的实例

// 抽象类中的抽象方法必须被子类实现
abstract class Animal9 {
  public name
  public constructor(name) {
    this.name = name
  }
  public abstract sayHi()
}
class Cat4 extends Animal9 {
  public sayHi() {
    console.log(`Meow, My name is ${this.name}`)
  }
}
let cat = new Cat4('Tom')
cat.sayHi()



// 类的类型
// 给类加上ts的类型，与接口类似：
class Animal10 {
  name: string;
  constructor(name: string) {
    this.name = name
  }
  sayHi(): string {
    return `My name is ${this.name}`
  }
}
let a5: Animal10 = new Animal10('Jack')
console.log(a5.sayHi())








