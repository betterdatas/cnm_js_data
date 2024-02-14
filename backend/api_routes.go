// backend/api_routes.go

package main

import (
	"fmt"
	"net/http"
)

// 处理Code Notation Manager API请求
func cnmAPIHandler(w http.ResponseWriter, r *http.Request) {
	// 在这里处理与Code Notation Manager相关的API逻辑
	// 例如：读取或修改注释信息，处理前端请求等
	fmt.Fprint(w, "Code Notation Manager API")
}
