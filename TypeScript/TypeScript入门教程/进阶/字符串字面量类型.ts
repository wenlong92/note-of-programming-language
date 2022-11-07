// 字符串字面量类型用来约束取值只能是某几个字符串中的一个

type EventNames = 'click' | 'scroll' | 'mousemove';
function handleEvent(ele: Element, event: EventNames) {
  // do something
}

handleEvent(document.getElementById('hello'), 'scroll')
handleEvent(document.getElementById('world'), 'dbclick')  //此时会报错，event不能为 'dbclick'

// 类型别名和字符串字面量类型都是用type进行定义