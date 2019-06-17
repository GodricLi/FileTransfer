package main

import (
	"fmt"
	"net/http"
)

// 使用net/http编写服务端。
// 服务端编写业务逻辑处理器，w给客户端回复数据，r读取客户端发送的数据
func Handler(w http.ResponseWriter, r *http.Request) {
	// 读取客户端数据
	fmt.Println("Method:", r.Method)
	fmt.Println("URL:", r.URL)
	fmt.Println("Herder:", r.Header)
	fmt.Println("Body:", r.Body)

	// 给客户端发送数据
	// fmt.Fprintln(w, "hello")	//相当于w.write
	w.Write([]byte("hello go"))
}

func WebServer() {
	// 注册处理函数，用户连接，自动调用指定的函数处理
	http.HandleFunc("/", Handler)
	// 在指定的地址进行监听，开启一个HTTP
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func main() {
	WebServer()
}
