// 数组合并了相同类型的对象，元组(Tuple)合并了不同类型的对象。
let tom2: [string, number] = ['Tom', 23]

// 赋值或访问一个已知索引的元素时，可以得到正确的类型
let tom3: [string, number]
tom3[0] = 'Tom'
tom3[1] = 233
tom3[0].slice(1)
tom3[1].toFixed(2)

// 也可以只赋值其中一项
let tom4: [string, number]
tom4[0] = 'Tom'


// 若直接对象元组类型的变量进行初始化或者赋值时，需要提供所有元组类型中指定的项
let tom5: [string, number]
tom5 = ["Tom"]  // 此时会报错



// 越界的元素：添加越界元素时，它的类型会被限制为元祖中每个类型的联合类型
let tom6: [string, number]
tom6 = ['Tom', 4]
tom6.push('male')
tom6.push(true)
