package main

import (
	"github.com/go-programming-tour-book/blog-service/internal/routers"
)

func main() {
	//创建默认engine实例
	r := routers.NewRouter()
	// 监听端口，启动服务
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
