package main

import (
	"fmt"
	"time"
)

// main.go 脚本

func main() {
	// 启动Server
	go StartServer()

	// TODO 你可以写其他逻辑
	fmt.Println("这是一个Go服务端，实现了Socket消息广播功能")

	// 防止主线程退出
	for {
		time.Sleep(1 * time.Second)
	}
}

// main.go 脚本

func StartServer() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
