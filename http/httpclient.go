package main

import (
	"fmt"
	"net"
	"net/http"
)

func TestResponse() {
	// 主动连接服务
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("connect err:", err)
		return
	}
	defer conn.Close()
	// 发送请求，才有响应
	requestBuf := "GET /go HTTP/1.1 "
	conn.Write([]byte(requestBuf))

	// 接收响应
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if n == 0 {
		fmt.Println("read err:", err)
		return
	}
	//打印响应信息
	fmt.Printf("%#v", string(buf[:n]))
}

// 使用net/http编写客户端访问服务器
func HttpClient() {
	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("htttp get err:", err)
		return
	}
	defer response.Body.Close()

	fmt.Println("Status:", response.Status)
	fmt.Println("StatusCode:", response.StatusCode)
	fmt.Println("Header:", response.Header)
	// fmt.Println("Body:", response.Body)	是一个数据流Body: &{0xc042160080 <nil> <nil>}

	// 读取Body
	buf := make([]byte, 1024)
	var temp string
	for {
		n, err := response.Body.Read(buf)
		if err != nil {
			fmt.Println("body read err:", err)
			break
		}
		temp += string(buf[:n])
	}
	fmt.Println("Body:", temp)

}

func main() {
	// TestResponse()
	HttpClient()
}
