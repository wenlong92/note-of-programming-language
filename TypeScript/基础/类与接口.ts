// 接口的另一个用途，对类的一部分行为进行抽象
// 一般来讲，一个类只能继承另一个类，有时不同类之间可以有一些共有的特性，这时就可以把特性提取成接口，用implements关键字实现
interface Alarm {
  alert(): void;
}
class Door{

}
class SecurityDoor extends Door implements Alarm {
  alert() {
    console.log("SecurityDoor alert")
  }
}
class Car implements Alarm {
  alert(){
    console.log("Car alert")
  }
}



// 一个类可以实现多个接口
interface Alarm1 {
  alert(): void;
}
interface Light {
  lightOn(): void;
  lightOff(): void;
}
class Car1 implements Alarm1, Light {
  alert() {
    console.log('Car alert')
  }
  lightOn() {
    console.log("Car light on")
  }
  lightOff() {
    console.log("Car light off")
  }
}


// 接口与接口之间可以是继承关系
interface Alarm2 {
  alert(): void;
}
interface LightableAlarm extends Alarm2 {
  lightOn(): void;
  lightOff(): void;
}
// LightalbeAlarm继承Alarm2，除了拥有alert方法，还拥有两个新方法


// 在常见的面向对象语言中，接口是不能继承类的，但在ts中可以
class Point {
  x: number;
  y: number;
  constructor(x: number, y: number) {
    this.x = x;
    this.y = y;
  }
}
interface Point3d extends Point {
  z: number;
}
let point3d: Point3d = {x: 1, y: 2, z: 3}