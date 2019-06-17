package main

import (
	"fmt"
	"net"
	"os"
)

func test_client() {
	//主动连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err= ", err)
		return
	}

	defer conn.Close()

	//发送数据
	conn.Write([]byte("hello"))

}

func tcp_client() { //持续通信
	// 主动连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("connect err:", err)
		return
	}
	defer conn.Close()

	//将输入键盘的数据发送给server
	go func() {
		data := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(data) //读取输入的内容，放在str里面
			if err != nil {
				fmt.Println("os.Stdin.Read err: ", err)
				return
			}
			//发送数据
			conn.Write(data[:n])
		}
	}()

	// 接收server发送过来的数据
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn read err: ", err)
			return
		}
		//转化成字符串，打印接收到的内容
		fmt.Println(string(buf[:n]))

	}

}

func main() {
	// test_client()
	tcp_client()
}
