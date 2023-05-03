package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"iswang.jie.com/m/v2/ch22/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//func main() {
//	// tcp
//	//client, err := rpc.Dial("tcp", "localhost:1234")
//	//http
//	//client, err := rpc.DialHTTP("tcp", "localhost:1234")
//	// tcp  json
//	client, err := jsonrpc.Dial("tcp", "localhost:1234")
//	if err != nil {
//		log.Fatal("dialing:", err)
//	}
//	args := server.Args{A: 7, B: 8}
//	var reply int
//	//调用的远程方法的名字，这里是MathService.Add，点前面的部分是注册的服务的名称，点后面的部分是该服务的方法；
//	//客户端为了调用远程方法提供的参数，示例中是args；
//	//为了接收远程方法返回的结果，必须是一个指针，也就是示例中的& replay，这样客户端就可以获得服务端返回的结果了。
//	err = client.Call("MathService.Add", args, &reply)
//
//	if err != nil {
//		log.Fatal("MathService.Add error:", err)
//	}
//	fmt.Printf("MathService.Add: %d+%d=%d", args.A, args.B, reply)
//}

// http json
func main() {

	client, err := DialHTTP("tcp", "localhost:1234")

	if err != nil {

		log.Fatal("dialing:", err)

	}

	args := server.Args{A: 7, B: 8}

	var reply int

	err = client.Call("MathService.Add", args, &reply)

	if err != nil {

		log.Fatal("MathService.Add error:", err)

	}

	fmt.Printf("MathService.Add: %d+%d=%d", args.A, args.B, reply)

}

// DialHTTP connects to an HTTP RPC server at the specified network address

// listening on the default HTTP RPC path.

func DialHTTP(network, address string) (*rpc.Client, error) {

	return DialHTTPPath(network, address, rpc.DefaultRPCPath)

}

// DialHTTPPath connects to an HTTP RPC server

// at the specified network address and path.

func DialHTTPPath(network, address, path string) (*rpc.Client, error) {

	var err error

	conn, err := net.Dial(network, address)

	if err != nil {

		return nil, err

	}

	io.WriteString(conn, "GET "+path+" HTTP/1.0\n\n")

	// Require successful HTTP response

	// before switching to RPC protocol.

	resp, err := http.ReadResponse(bufio.NewReader(conn), &http.Request{Method: "GET"})

	connected := "200 Connected to JSON RPC"

	if err == nil && resp.Status == connected {

		return jsonrpc.NewClient(conn), nil

	}

	if err == nil {

		err = errors.New("unexpected HTTP response: " + resp.Status)

	}

	conn.Close()

	return nil, &net.OpError{

		Op: "dial-http",

		Net: network + " " + address,

		Addr: nil,

		Err: err,
	}

}
