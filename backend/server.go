// backend/server.go

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 设定服务器端口
	port := 8080
	fmt.Printf("Server is running on :%d\n", port)

	// 设置路由
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/cnm", cnmAPIHandler)

	// 启动服务器
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}

// 处理主页请求
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Code Notation Manager")
}

// 处理Code Notation Manager API请求
func cnmAPIHandler(w http.ResponseWriter, r *http.Request) {
	// 在这里处理与Code Notation Manager相关的API逻辑
	// 例如：读取或修改注释信息，处理前端请求等
	fmt.Fprint(w, "Code Notation Manager API")
}
