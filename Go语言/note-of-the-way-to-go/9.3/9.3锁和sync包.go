package main

import (
	"sync"
	"fmt"
)

func main() {
	// 在一些复杂程序中，一般通过不同线程执行不同应用是实现程序的并发。
	// 当不同线程使用同一变量时，会出现一个问题：无法预知变量被不同线程修改的顺序(资源竞争，指不同线程对同一变量使用的竞争)

	// 经典做法是一次只让一个线程对共享变量进行操作。当变量被一个线程改变时，为它上锁，知道这个线程执行完并解锁后，其他线程才能访问
	// 之前学习map类型时不存在锁的机制，所以map类型时非线程安全的。当并行访问一个共享map类型数据时，map数据会出错

	// Go中锁的机制是通过sync包中Mutex实现的。sync.Mutex是一个互斥锁，作用是守护在临界区入口以确保同一时间只能有一个线程进入临界区
	// 假设info是一个需要上锁的放在共享内存中的变量，实现例子如下：
	type Info struct {
		mu sync.Mutex
	}
	// 如果函数想要改变info变量：
	func Update(info *Info) {
		info.mu.Lock()
		info.Str = new value
		info.mu.Unlock()
	}

	// 还有一个有用的例子是通过Mutex实现一个可以上锁的共享缓冲器
	tyep SyncedBuffer struct {
		lock sync.Mutex
		buffer bytes.Buffer
	}

	// sync包中还有一个RWMutex锁：通过RLock()允许同一时间多个线程对变量进行读操作，但只能一个线程进行写操作
	// 如果使用Lock()将和普通的Mutex作用相同
	// 包中还有一个Once类型变量的方法 once.Do(call) 确保被调用函数只能被调用一次

	// 相对简单的情况下，可使用sync包，如果这种方式导致程序明显变慢或者引起其他问题
	// 可通过goroutines和channels来解决问题
}