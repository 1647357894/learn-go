package server

// 定义了MathService，用于表示一个远程服务对象；
// Args 结构体用于表示参数；
// Add 这个方法实现了加法的功能，加法的结果通过 replay这个指针变量返回。

// 把一个对象注册为 RPC 服务，可以让客户端远程访问，那么该对象（类型）的方法必须满足如下条件：
// 方法的类型是可导出的（公开的）；
// 方法本身也是可导出的；
// 方法必须有 2 个参数，并且参数类型是可导出或者内建的；
// 方法必须返回一个 error 类型。
//方法的格式
//func (t *T) MethodName(argType T1, replyType *T2) error

type MathService struct {
}

type Args struct {
	A, B int
}

func (m *MathService) Add(args Args, reply *int) error {

	*reply = args.A + args.B

	return nil

}
