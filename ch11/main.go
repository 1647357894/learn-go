package main

import (
	"fmt"
	"sync"
	"time"
)

// select timeout 模式的核心在于通过 time.After 函数设置一个超时时间，防止因为异常造成 select 语句的无限等待。
// 如果可以使用 Context 的 WithTimeout 函数超时取消，要优先使用。
func withTimeOutMode() {
	//time.Now().Unix()
	fmt.Println("start ", time.Now().Format(time.DateTime))
	result := make(chan string)
	defer close(result)

	go func() {
		//模拟网络访问
		time.Sleep(8 * time.Second)
		result <- "服务端结果"
	}()

	select {
	case v := <-result:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("网络访问超时了")
	}
	fmt.Println("end ", time.Now().Format(time.DateTime))
}

//	for select 循环模式
//
// 这种模式会一直执行 default 语句中的任务，直到 done 这个 channel 被关闭为止。
func forLoopMode() bool {
	done := make(chan bool)
	defer close(done)
	for {
		select {
		case <-done:
			return true
		default:
			//执行具体的任务
		}
	}
	return false
}

// for range select 有限循环，一般用于把可以迭代的内容发送到 channel上
// 这种模式也会有一个 done channel，用于退出当前的 for 循环，而另外一个 resultCh channel 用于接收 for range 循环的值，这些值通过 resultCh 可以传送给其他的调用者。
func forRangeMode() {
	done := make(chan bool)
	resultCh := make(chan int)
	defer close(done)
	defer close(resultCh)
	for _, s := range []int{} {
		select {
		case <-done:
			return
		case resultCh <- s:
		}
	}
}

//Pipeline 模式
//流水线由一道道工序构成，每道工序通过 channel 把数据传递到下一个工序；
//每道工序一般会对应一个函数，函数里有协程和 channel，协程一般用于处理数据并把它放入 channel 中，整个函数会返回这个 channel 以供下一道工序使用；
//最终要有一个组织者（示例中的 main 函数）把这些工序串起来，这样就形成了一个完整的流水线，对于数据来说就是数据流。

// <-chan只读channel
// chan<- 只写channel
// 普通chan可以转换为只读或只写chan，只读只写不能反向转换
func buy(n int) <-chan string {

	//带缓冲区chan
	//没有缓冲区的chan，大小只有0,只能作为数据传递使用，放不下了则阻塞
	out := make(chan string, n)

	//go func() {
	defer close(out)
	for i := 0; i < n; i++ {
		out <- fmt.Sprint(" <-配件", i)
	}
	//}()
	return out
}

func build(in <-chan string, name string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			out <- fmt.Sprintf(" <-装配(%s) %s", name, s)
		}
	}()
	return out
}

func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			out <- fmt.Sprint("打包", s)
		}
	}()
	return out
}

//扇出和扇入模式
//以工序 1 为中点，三条传递数据的线发散出去，就像一把打开的扇子一样，所以叫扇出。
//以 merge 组件为中点，三条传递数据的线汇聚到 merge 组件，也像一把打开的扇子一样，所以叫扇入。

// 合并多个channel
func merge(cs ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan string) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

//Futures 模式
//Pipeline 流水线模式中的工序是相互依赖的，上一道工序做完，下一道工序才能开始。但是在我们的实际需求中，也有大量的任务之间相互独立、没有依赖，所以为了提高性能，这些独立的任务就可以并发执行。
//Futures 模式可以理解为未来模式，主协程不用等待子协程返回的结果，可以先去做其他事情，等未来需要子协程结果的时候再来取，如果子协程还没有返回结果，就一直等待。

//洗菜

func washVegetables() <-chan string {
	vegetables := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		vegetables <- "洗好的菜"
	}()
	return vegetables
}

//烧水

func boilWater() <-chan string {
	water := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		water <- "烧开的水"
	}()
	return water
}

func main() {

	//if forLoopMode() {
	//	return
	//}

	//forRangeMode()

	//withTimeOutMode()

	//pipeline
	//coms := buy(10)
	//
	////扇出
	//phones1 := build(coms, fmt.Sprint("生产线", 1))
	//phones2 := build(coms, fmt.Sprint("生产线", 2))
	//phones3 := build(coms, fmt.Sprint("生产线", 3))
	//
	////扇入 merge合并
	//phones := merge(phones1, phones2, phones3)
	//packs := pack(phones)
	//for s := range packs {
	//	fmt.Println(s)
	//}

	//Futures
	vegetablesCh := washVegetables() //洗菜

	waterCh := boilWater() //烧水

	fmt.Println("已经安排洗菜和烧水了，我先眯一会")

	time.Sleep(2 * time.Second)

	fmt.Println("要做火锅了，看看菜和水好了吗")

	vegetables := <-vegetablesCh

	water := <-waterCh

	fmt.Println("准备好了，可以做火锅了:", vegetables, water)

}
