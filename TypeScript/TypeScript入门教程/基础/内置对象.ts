// JavaScript中的内置对象，可以直接在ts中当做定义好的类型


// ECMAScript的内置对象：Boolean、Error、Date、RegExp等
// 可以在ts中将变量定义为这些类型
let b: Boolean = new Boolean(1)
let e: Error = new Error('Error occurred')
let d: Date = new Date();
let r: RegExp = /[a-z]/


// DOM和BOM的内置对象：Document、HTMLElement、Event、NodeList等
let body: HTMLElement = document.body;
let allDiv: NodeList = document.querySelectorAll('div')
document.addEventListener('click', function(e: MouseEvent) {
  // Do something
})


// ts核心库的定义文件：定义了所有浏览器环境需要用到的类型，且预置在ts中
