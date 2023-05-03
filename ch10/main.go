package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 协程如何退出
//Context
//一个任务会有很多个协程协作完成，一次 HTTP 请求也会触发很多个协程的启动，而这些协程有可能会启动更多的子协程，并且无法预知有多少层协程、每一层有多少个协程。
//如果因为某些原因导致任务终止了，HTTP 请求取消了，那么它们启动的协程怎么办？该如何取消呢？因为取消这些协程可以节约内存，提升性能，同时避免不可预料的 Bug。
//Context 就是用来简化解决这些问题的，并且是并发安全的。Context 是一个接口，它具备手动、定时、超时发出取消信号、传值等功能，主要用于控制多个协程之间的协作，尤其是取消操作。
//一旦取消指令下达，那么被 Context 跟踪的这些协程都会收到取消信号，就可以做清理和退出操作。

//Deadline 方法可以获取设置的截止时间，第一个返回值 deadline 是截止时间，到了这个时间点，Context 会自动发起取消请求，第二个返回值 ok 代表是否设置了截止时间。
//Done 方法返回一个只读的 channel，类型为 struct{}。在协程中，如果该方法返回的 chan 可以读取，则意味着 Context 已经发起了取消信号。通过 Done 方法收到这个信号后，就可以做清理操作，然后退出协程，释放资源。
//Err 方法返回取消的错误原因，即因为什么原因 Context 被取消。
//Value 方法获取该 Context 上绑定的值，是一个键值对，所以要通过一个 key 才可以获取对应的值。

//Context 不要放在结构体中，要以参数的方式传递。
//Context 作为函数的参数时，要放在第一位，也就是第一个参数。
//要使用 context.Background 函数生成根节点的 Context，也就是最顶层的 Context。
//Context 传值要传递必须的值，而且要尽可能地少，不要什么都传。
//Context 多协程安全，可以在多个协程中放心使用。

//              <- ctx4
//      <- ctx2 <- ctx5
// ctx1
//      <- ctx3 <-ctx6
//
// ctx2 cancel，结果 ctx2,ctx4,ctx5 都会cancel，其他的不会cancel

//func main() {
//
//	var wg sync.WaitGroup
//
//	wg.Add(1)
//
//	stopCh := make(chan bool) //用来停止监控狗
//
//	go func() {
//
//		defer wg.Done()
//
//		watchDog(stopCh, "【监控狗1】")
//
//	}()
//
//	time.Sleep(5 * time.Second) //先让监控狗监控5秒
//
//	stopCh <- true //发停止指令
//
//	wg.Wait()
//
//}
//
//func watchDog(stopCh chan bool, name string) {
//
//	//开启for select循环，一直后台监控
//
//	for {
//
//		select {
//
//		case <-stopCh:
//
//			fmt.Println(name, "停止指令已收到，马上停止")
//
//			return
//
//		default:
//
//			fmt.Println(name, "正在监控……")
//
//		}
//
//		time.Sleep(1 * time.Second)
//
//	}
//
//}

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	background := context.Background()

	//生成一个可取消的 Context。
	ctx, stop := context.WithCancel(background)

	//生成一个可定时取消的 Context，参数 d 为定时取消的具体时间。
	//context.WithDeadline(background, time.Now())
	//生成一个可超时取消的 Context，参数 timeout 用于设置多久后取消
	//context.WithTimeout(background, time.Second)
	//生成一个可携带 key-value 键值对的 Context。
	//context.WithValue(background, "k", "v")

	//以上四个生成 Context 的函数中，前三个都属于可取消的 Context，它们是一类函数，最后一个是值 Context，用于存储一个 key-value 键值对。

	go func() {
		defer wg.Done()
		watchDog(ctx, "【监控狗1】")
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "【监控狗2】")
	}()

	valCtx := context.WithValue(ctx, "userId", 2)
	//valCtx := context.WithValue(background, "userId", 2)

	go func() {

		defer wg.Done()

		getUser(valCtx)

	}()

	time.Sleep(5 * time.Second) //先让监控狗监控5秒

	stop() //发停止指令

	wg.Wait()

	//通过 Context 实现日志跟踪
	//在用户请求的入口点生成 TraceID。
	//通过 context.WithValue 保存 TraceID。
	//然后这个保存着 TraceID 的 Context 就可以作为参数在各个协程或者函数间传递。
	//在需要记录日志的地方，通过 Context 的 Value 方法获取保存的 TraceID，然后把它和其他日志信息记录下来。
	//这样具备同样 TraceID 的日志就可以被串联起来，达到日志跟踪的目的。
	//以上思路实现的核心是 Context 的传值功能。
}

func watchDog(ctx context.Context, name string) {
	//开启for select循环，一直后台监控
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控……")
		}
		time.Sleep(1 * time.Second)
	}
}

func getUser(ctx context.Context) {

	for {

		select {

		case <-ctx.Done():

			fmt.Println("【获取用户】", "协程退出")

			return

		default:

			userId := ctx.Value("userId")

			fmt.Println("【获取用户】", "用户ID为：", userId)

			time.Sleep(1 * time.Second)

		}

	}

}
