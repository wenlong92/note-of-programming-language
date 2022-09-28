interface Person {
  name: string;
  age: number;
}
let tom: Person = {
  name: 'Tom',
  age: 23
}
// 定义一个接口Person，变量tom的类型时Person，这就约束了tom的形状必须和接口Person一致
// 接口一般首字母大写
// 变量比接口少一些属性是不允许的，多一些属性也不允许


// 如果不希望完全匹配，可以用可选属性，但仍不允许添加未定义属性
interface PersonOne {
  name: string;
  age?: number;
}
let jack: PersonOne = {
  name: 'jack'
}

// 如果希望接口允许有任意属性
interface PersonTwo {
  name: string;
  age?: number;
  [propName: string]: string  // 使用[propName: string]定义了任意属性取string类型的值
}
// 此时，一旦定义任意属性，则确定属性和可选属性必须是它的类型的自己，所以PersonTwo接口会报错

interface PersonThree {
  name: string;
  age?: number;
  [propName: string]: string | number  //一个接口中只能有一个任意属性，如果接口中有多个类型的属性，可以在任意属性中使用联合类型
}

// 如果希望对象中某些字段只能在创建时被赋值，则可以使用readonly定义只读属性
interface PersonFour {
  readonly id: number;
  name: string;
  age?: number;
  [propName: string]: any
}
let rose: PersonFour = {
  id: 12345,
  name: 'rose',
  gender: 'female'
}
rose.id = 345  //此时会报错