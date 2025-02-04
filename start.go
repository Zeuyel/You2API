package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	api "you2api/api" // 请替换为您的实际项目名
)

func main() {
	// 定义命令行参数
	proxyPort := flag.String("proxy", "10809", "代理服务器端口")
	serverPort := flag.String("port", "8080", "服务器监听端口")
	flag.Parse()

	// 设置代理端口
	api.SetProxyPort(*proxyPort)

	// 将 /api/main.go 的 Handler 函数注册到根路径
	http.HandleFunc("/", api.Handler)

	// 启动服务器
	fmt.Printf("Server is running on http://localhost:%s\n", *serverPort)
	fmt.Printf("Using proxy at http://127.0.0.1:%s\n", *proxyPort)

	if err := http.ListenAndServe(":"+*serverPort, nil); err != nil {
		log.Fatal("Server error: ", err)
	}
}
