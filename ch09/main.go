package main

import (
	"fmt"
	"sync"
	"time"
)

//sync

// 共享的资源
//var sum = 0
//
//func add(i int) {
//	sum += i
//}

//sync.Mutex 互斥锁
// Lock 和 Unlock，代表加锁和解锁。当一个协程获得 Mutex 锁后，其他协程只能等到 Mutex 锁释放后才能再次获得锁

var (
	sum int
	//mutex sync.Mutex
	mutex sync.RWMutex
)

//func add(i int) {
//	mutex.Lock()
//	defer mutex.Unlock()
//	sum += i
//}

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}

// 增加了一个读取sum的函数
// 多个 goroutine 可以同时读数据，不再相互等待。
func readSum() int {
	//mutex.Lock()
	//defer mutex.Unlock()
	//只获取读锁
	mutex.RLock()
	defer mutex.RUnlock()
	b := sum
	return b
}

func doOnce() {
	//sync.Once 来保证代码只执行一次
	//适用于创建某个对象的单例、只加载一次的资源等只执行一次的场景。
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}

	//用于等待协程执行完毕
	var wg sync.WaitGroup
	wg.Add(10)

	//启动10个协程执行once.Do(onceBody)
	for i := 0; i < 10; i++ {
		go func() {
			//把要执行的函数(方法)作为参数传给once.Do方法即可
			once.Do(onceBody)
			//onceBody()
			wg.Done()
		}()
	}

	wg.Wait()
}

// 共享资源，需要锁保护
var done = false

// 10个人赛跑，1个裁判发号施令
func race() {

	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)

	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()

			fmt.Println(num, "号已经就位")
			cond.L.Lock()
			defer cond.L.Unlock()

			//伪唤醒问题，loop 处理
			for !done {
				//等待
				//使用时需要加锁
				cond.Wait()
			}
			fmt.Println(num, "号开始跑……")

		}(i)
	}

	//等待所有goroutine都进入wait状态
	time.Sleep(2 * time.Second)

	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		done = true
		cond.Broadcast() //唤醒所有等待的协程
		//cond.Signal() //随机唤醒一个协程
	}()
	//防止函数提前返回退出
	wg.Wait()
}

func run() {

	var wg sync.WaitGroup
	//因为要监控1010个协程，所以设置计数器为1010
	wg.Add(1010)

	//开启100个协程让sum+10
	for i := 0; i < 1000; i++ {
		go func() {
			add(1)
			//计数器值减1
			wg.Done()
		}()
	}

	//写的时候不能同时读，因为这个时候读取的话可能读到脏数据（不正确的数据）；
	//读的时候不能同时写，因为也可能产生不可预料的结果；
	//读的时候可以同时读，因为数据不会改变，所以不管多少个 goroutine 读都是并发安全的。

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("和为:", readSum())
			//计数器值减1
			wg.Done()
		}()
	}
	//防止提前退出
	//time.Sleep(2 * time.Second)

	//一直等待，只要计数器值为0
	wg.Wait()
	fmt.Println("======================")
	fmt.Println("和为:", sum)
}

func main() {

	//run()
	//doOnce()
	//race()

	var scene sync.Map

	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
