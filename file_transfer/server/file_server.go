package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func Server_file() {
	//1.监听
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net listen err:", err)
		return
	}
	fmt.Println("connect successful")
	defer listener.Close()

	//2.阻塞，等待用户连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listen accept err: ", err)
		return
	}
	defer conn.Close()

	// 3.接收用户请求
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn read err:", err)
		return
	}
	//3.1获取接收到的文件名
	fileName := string(buf[:n])
	fmt.Println(fileName)
	// 3.2回复确认，“ok”
	conn.Write([]byte("ok"))

	//4.接收文件内容
	RecvFile(fileName, conn)

}

func RecvFile(fileName string, conn net.Conn) {
	//1.新建文件
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("file create err :", err)
		return
	}

	//2.接收文件内容，接收一次写一次
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("file recv complete")
				return
			} else {
				fmt.Println("file recv err :", err)
				return
			}
			//文件发送完成，收到数据为0
			// if n == 0 {
			// 	fmt.Println("n=0，file recv complete")
			// 	break
			// }
		}
		//往文件写入内容
		file.Write(buf[:n])
	}
}

func main() {
	Server_file()
}
