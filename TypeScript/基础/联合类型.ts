let myFavoriteNumberOne: string | number;
myFavoriteNumberOne = 'seven'
myFavoriteNumberOne = 7

// 当不确定一个联合类型的变量到底是哪个类型时，只能访问此联合类型的所有类型里公有的属性和方法
function getLength(something: string | number): number {
  return something.length  //因为length不是string和number的共有属性，所以报错
}
function getLength1(something: string | number): string {
  return something.toString()  // 此时不会报错
}