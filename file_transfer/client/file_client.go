package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func filename() (filename string) {
	//获取文件信息，文件名,终端输入带路径的文件名
	list := os.Args
	if len(list) != 2 {
		fmt.Println("useage :canton find ")
		return
	}

	file_name := list[1]
	info, err := os.Stat(file_name)
	if err != nil {
		fmt.Println("os.Stat err : ", err)
		return
	}
	return info.Name()
	// fmt.Println("name:", info.Name()) //获取文件名
	// fmt.Println("size:", info.Size()) //获取文件大小
}

// 发送文件致服务端
func cilent_file_transfer() {
	// 1.主动连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}
	defer conn.Close()

	//2.先给接收方发送文件名
	//2.1获取文件名
	fmt.Println("请输入文件名：")
	var path string
	fmt.Scan(&path) //拿到文件路径

	// 2.2获取文件名
	info, err_f := os.Stat(path)
	if err_f != nil {
		fmt.Println("os stat err: ", err)
		return
	}

	file_name := info.Name()
	_, err = conn.Write([]byte(file_name))
	if err != nil {
		fmt.Println("conn write err:", err)
		return
	}

	//3.接收对方的确认，如果回复“ok”，说明对方准备好了，可以发送文件
	var n int
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn read err:", err)
		return
	}
	if "ok" == string(buf[:n]) {
		//4.发送文件内容
		send_file(path, conn)
	}

}

func send_file(path string, conn net.Conn) {
	//1.以只读方式打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("os open file err: ", err)
		return
	}
	defer conn.Close()

	// 2.读取文件内容并发送,循环读取文件内容，
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf) // 读取内容
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完成！")
				return
			} else {
				fmt.Println("file read err: ", err)
				return
			}
		}
		// 将每次循环读取的内容发送
		conn.Write(buf[:n])
	}

}

func main() {
	// filename()
	cilent_file_transfer()

}
