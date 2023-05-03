package main

import "testing"

// go test -v ./ch18 执行该目录下的所有单元测试

//单元测试覆盖率
//go test -v --coverprofile=ch18.cover ./ch18

//查看单元测试覆盖率报告
//go tool cover -html=ch18.cover -o=ch18.html

// 单元测试规则
// 含有单元测试代码的 go 文件必须以 _test.go 结尾，Go 语言测试工具只认符合这个规则的文件。
// 单元测试文件名 _test.go 前面的部分最好是被测试的函数所在的 go 文件的文件名，比如以上示例中单元测试文件叫 main_test.go，因为测试的 Fibonacci 函数在 main.go 文件里。
// 单元测试的函数名必须以 Test 开头，是可导出的、公开的函数。
// 测试函数的签名必须接收一个指向 testing.T 类型的指针，并且不能返回任何值。
// 函数名最好是 Test + 要测试的函数名，比如例子中是 TestFibonacci，表示测试的是 Fibonacci 这个函数。

//单元测试是保证代码质量的好方法，但单元测试也不是万能的，使用它可以降低 Bug 率，但也不要完全依赖。除了单元测试外，还可以辅以 Code Review、人工测试等手段更好地保证代码质量。

func TestFibonacci(t *testing.T) {

	//预先定义的一组斐波那契数列作为测试用例

	fsMap := map[int]int{}
	fsMap[-1] = 0

	fsMap[0] = 0

	fsMap[1] = 1

	fsMap[2] = 1

	fsMap[3] = 2

	fsMap[4] = 3

	fsMap[5] = 5

	fsMap[6] = 8

	fsMap[7] = 13

	fsMap[8] = 21

	fsMap[9] = 34

	for k, v := range fsMap {

		fib := Fibonacci(k)

		if v == fib {

			t.Logf("结果正确:n为%d,值为%d", k, fib)

		} else {

			t.Errorf("结果错误：期望%d,但是计算的值是%d", v, fib)

		}
	}
}

// 基准测试
// 基准测试（Benchmark）是一项用于测量和评估软件性能指标的方法，主要用于评估你写的代码的性能。

// 基准测试函数必须以 Benchmark 开头，必须是可导出的；
// 函数的签名必须接收一个指向 testing.B 类型的指针，并且不能返回任何值；
// 最后的 for 循环很重要，被测试的代码要放到循环里；
// b.N 是基准测试框架提供的，表示循环的次数，因为需要反复调用测试的代码，才可以评估性能。

//go test  -benchmem -bench=. -benchtime=3s ./ch18

//-benchmem  内存统计
//-benchtime 执行时间

//goos: windows
//goarch: amd64
//pkg: iswang.jie.com/m/v2/ch18
//cpu: 13th Gen Intel(R) Core(TM) i5-13600KF
//BenchmarkFibonacci
//BenchmarkFibonacci-20(20线程)            5884945(循环次数)               210.6 ns/op (每次执行时间 这里是210.6 纳秒)
//PASS

func BenchmarkFibonacci(b *testing.B) {

	n := 10
	//进行基准测试之前会做一些准备，比如构建测试数据等，这些准备也需要消耗时间，所以需要把这部分时间排除在外。这就需要通过 ResetTimer 方法重置计时器
	b.ResetTimer() //重置计时器
	//StartTimer 和 StopTimer 方法，控制什么时候开始计时、什么时候停止计时。

	//统计每次操作分配内存的次数，以及每次操作分配的字节数，这两个指标可以作为优化代码的参考
	b.ReportAllocs() //开启内存统计

	for i := 0; i < b.N; i++ {
		Fibonacci(n)
	}
}

//并发基准测试
//测试在多个 goroutine 并发下代码的性能

func BenchmarkFibonacciRunParallel(b *testing.B) {

	n := 10
	//Go 语言通过 RunParallel 方法运行并发基准测试。RunParallel 方法会创建多个 goroutine，并将 b.N 分配给这些 goroutine 执行。
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Fibonacci(n)
		}
	})

}
