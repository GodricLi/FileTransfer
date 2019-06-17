package main

import (
	"fmt"
	"net"
	"strings"
)

func tcp_server() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	//阻塞，等待用户连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("err= ", err)
		return
	}

	//接收用户请求
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Println("buf=", string(buf[:n]))

	defer listener.Close()
	defer conn.Close()
}

func send_msg(conn net.Conn) { //返回大写字母

	defer conn.Close()

	// 获取客户端连接地址
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "connect  successful!")

	buf := make([]byte, 1024)
	// 读取用户数据
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err= ", err)
			return
		}
		// 用户输入exit判断
		// if string(buf[:n-1]) == "exit" { //将nc工具测试时在末尾增加一个换行符去掉
		if string(buf[:n-2]) == "exit" { //n-2:去掉win系统输入回车时在末尾添加的"\r\n"两个字符
			fmt.Printf("%s:exit!\n", addr)
			return
		}
		//把数据转化为大写发送给客户端
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}

func concurrent_server() { //实现并发服务器

	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()

	// 循环接收用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err=", err)
			return
		}

		//并发调用处理数据函数
		go send_msg(conn)
	}
}

func main() {
	// tcp_server()
	concurrent_server()
}
