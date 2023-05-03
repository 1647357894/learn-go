package main

import (
	"io"
	"iswang.jie.com/m/v2/ch22/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// RPC 服务
// RPC，也就是远程过程调用，是分布式系统中不同节点调用的方式（进程间通信），属于 C/S 模式。RPC 由客户端发起，调用服务端的方法进行通信，然后服务端把结果返回给客户端。
// RPC的核心有两个：通信协议和序列化。在 HTTP 2 之前，一般采用自定义 TCP 协议的方式进行通信，HTTP 2 出来后，也有采用该协议的，比如流行的gRPC。
// 序列化和反序列化是一种把传输内容编码和解码的方式，常见的编解码方式有 JSON、Protobuf 等。
// 在大多数 RPC的架构设计中，都有Client、Client Stub、Server、Server Stub这四个组件，Client 和 Server 之间通过 Socket 进行通信。RPC 架构如下图所示：
// https://learn.lianglianglee.com/%E4%B8%93%E6%A0%8F/22%20%E8%AE%B2%E9%80%9A%E5%85%B3%20Go%20%E8%AF%AD%E8%A8%80-%E5%AE%8C/assets/CgqCHl_8K6eADlRHAAFxSlJHXWc596.png

// RPC 调用的流程：
// 客户端（Client）调用客户端存根（Client Stub），同时把参数传给客户端存根；
// 客户端存根将参数打包编码，并通过系统调用发送到服务端；
// 客户端本地系统发送信息到服务器；
// 服务器系统将信息发送到服务端存根（Server Stub）；
// 服务端存根解析信息，也就是解码；
// 服务端存根调用真正的服务端程序（Sever）；
// 服务端（Server）处理后，通过同样的方式，把结果再返回给客户端（Client）。

//func main() {
//
//	//通过 RegisterName 函数注册了一个服务对象，该函数接收两个参数：
//	//服务名称（MathService）；
//	//具体的服务对象，也就是我刚刚定义好的MathService 这个结构体。
//	rpc.RegisterName("MathService", new(server.MathService))
//	//改用http
//	rpc.HandleHTTP()
//	l, e := net.Listen("tcp", ":1234")
//	if e != nil {
//		log.Fatal("listen error:", e)
//	}
//	//rpc.Accept(l)
//	//换成http的服务
//	http.Serve(l, nil)
//
//	//	run
//	//	go run ch22/server_main.go
//	//  go run ch22/client_main.go
//
//	//Go 语言 net/rpc 包提供的 HTTP 协议的 RPC调试URL  http://localhost:1234/debug/rpc
//}

// tcp json rpc
//func main() {
//	rpc.RegisterName("MathService", new(server.MathService))
//	l, e := net.Listen("tcp", ":1234")
//	if e != nil {
//		log.Fatal("listen error:", e)
//	}
//	for {
//		conn, err := l.Accept()
//		if err != nil {
//			log.Println("jsonrpc.Serve: accept:", err.Error())
//			return
//		}
//		//json rpc
//		go jsonrpc.ServeConn(conn)
//	}
//}

// http json
func main() {

	rpc.RegisterName("MathService", new(server.MathService))

	//注册一个path，用于提供基于http的json rpc服务

	http.HandleFunc(rpc.DefaultRPCPath, func(rw http.ResponseWriter, r *http.Request) {

		conn, _, err := rw.(http.Hijacker).Hijack()

		if err != nil {

			log.Print("rpc hijacking ", r.RemoteAddr, ": ", err.Error())

			return

		}

		var connected = "200 Connected to JSON RPC"

		io.WriteString(conn, "HTTP/1.0 "+connected+"\n\n")

		jsonrpc.ServeConn(conn)

	})

	l, e := net.Listen("tcp", ":1234")

	if e != nil {

		log.Fatal("listen error:", e)

	}

	http.Serve(l, nil) //换成http的服务

}

//TODO grpc protobuf
