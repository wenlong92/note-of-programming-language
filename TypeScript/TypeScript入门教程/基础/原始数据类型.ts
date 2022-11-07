function sayHello(person: string) {
  return "Hello, " + person
}
let user = "Tom";
console.log(sayHello(user))

let isDone: boolean = false;
let createByBoolean: boolean = Boolean(1)

let decLiteral: number = 6
let hexLiteral: number = 0xf00d

let binaryLiteral: number = 0b010

let octalLiteral: number = 0o744

let notANumber: number = NaN
let infinityNumber: number = Infinity;

let myName: string = "Tom";
let myAge: number = 23;

let sentence: string = `Hello, my name is ${myName}. 
I'll be ${myAge + 1} years old next monty.`

function alertName(): void {
  alert('My name is Tom')
}

let u: undefined = undefined
let n: null = null;

