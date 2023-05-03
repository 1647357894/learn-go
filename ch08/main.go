package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 协程（Goroutine）
// Go 语言中没有线程的概念，只有协程，也称为 goroutine。相比线程来说，协程更加轻量，一个程序可以随意启动成千上万个 goroutine。

//数据流动、传递的场景中要优先使用 channel，它是并发安全的，性能也不错。
// channel 内部使用了互斥锁来保证并发的安全

// chan<- //只写
func counter(out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		out <- i //如果对方不读 会阻塞
	}
}

// <-chan //只读
func printer(in <-chan int) {
	for num := range in {
		time.Sleep(time.Second)
		fmt.Println(num)
	}
}

func downloadFile(chanName string) string {

	//模拟下载文件,可以自己随机time.Sleep点时间试试
	seconds := rand.Intn(3)
	//time.Sleep(time.Second)
	time.Sleep(time.Duration(seconds) * time.Second)
	return chanName + ":filePath"
}

func main() {

	//go 关键字启动协程
	go fmt.Println("hello")
	fmt.Println("我是 main goroutine")
	time.Sleep(time.Second)

	//channel协程之间通信
	//接收：获取 chan 中的值，操作符为 <- chan。
	//发送：向 chan 发送值，把值放在 chan 中，操作符为 chan <-。

	//无缓冲 channel
	//使用 make 创建的 chan 就是一个无缓冲 channel，它的容量是 0，不能存储任何数据。
	//所以无缓冲 channel 只起到传输数据的作用，数据并不会在 channel 中做任何停留。
	//这也意味着，无缓冲 channel 的发送和接收操作是同时进行的，它也可以称为同步 channel。
	ch := make(chan string)
	go func() {
		fmt.Println("go")
		ch <- "goroutine 完成"
		ch <- "goroutine 完成"
	}()
	fmt.Println("main goroutine")
	//同步阻塞获取值
	v := <-ch
	fmt.Println("接收到的chan中的值为：", v)

	//有缓冲 channel
	//有缓冲 channel 的内部有一个缓冲队列；
	//发送操作是向队列的尾部插入元素，如果队列已满，则阻塞等待，直到另一个 goroutine 执行，接收操作释放队列的空间；
	//接收操作是从队列的头部获取元素并把它从队列中删除，如果队列为空，则阻塞等待，直到另一个 goroutine 执行，发送操作插入新的元素。
	cacheCh := make(chan int, 5)
	go func() {
		fmt.Println("go")
		time.Sleep(time.Second)
		cacheCh <- 1
		cacheCh <- 2
		cacheCh <- 3
	}()
	fmt.Println("main goroutine")
	//获取值
	v2 := <-cacheCh
	fmt.Println("接收到的cacheCh中的值为：", v2)
	fmt.Println("cacheCh容量为:", cap(cacheCh), ",元素个数为：", len(cacheCh))

	//	关闭channel
	// channel 被关闭了，就不能向里面发送数据了，如果发送的话，会引起 painc 异常。但是还可以接收 channel 里的数据，如果 channel 里没有数据的话，接收的数据是元素类型的零值。

	close(cacheCh)

	for cacheCh != nil {
		v2 := <-cacheCh
		if v2 == 0 {
			break
		}
		fmt.Println("接收到的cacheCh中的值为：", v2)
	}

	//	单向 channel
	//onlySend := make(chan<- int)
	//onlyReceive := make(<-chan int)
	c := make(chan int) //   chan   //读写

	go counter(c) //生产者
	printer(c)    //消费者

	fmt.Println("done")

	//多路复用
	//多路复用可以简单地理解为，N 个 channel 中，任意一个 channel 有数据产生，select 都可以监听到，然后执行相应的分支，接收数据并处理
	//如果这些 case 中有一个可以执行，select 语句会选择该 case 执行，如果同时有多个 case 可以被执行，则随机选择一个，这样每个 case 都有平等的被执行的机会。
	//如果一个 select 没有任何 case，那么它会一直等待下去。

	//select+channel 示例

	//声明三个存放结果的channel
	firstCh := make(chan string)
	secondCh := make(chan string)
	threeCh := make(chan string)

	//同时开启3个goroutine下载
	go func() {
		firstCh <- downloadFile("firstCh")
	}()

	go func() {
		secondCh <- downloadFile("secondCh")
	}()

	go func() {
		threeCh <- downloadFile("threeCh")
	}()

	//开始select多路复用，哪个channel能获取到值，
	//就说明哪个最先下载好，就用哪个。
	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	case filePath := <-threeCh:
		fmt.Println(filePath)
	}
}
