package main

import (
	"fmt"
	"log"
	"net/http"
)

// w 表示response对象，返回给客户端的内容都在对象里处理
// r 表示客户端请求对象，包含了请求头，请求参数等等
func index(w http.ResponseWriter, r *http.Request) {
	// 往w里写内容，就会在浏览器里输出
	fmt.Fprintf(w, "Helll Golang Http!")
}

func main() {
	// 设置路由，如果访问/,则调用index方法
	http.HandleFunc("/", index)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
