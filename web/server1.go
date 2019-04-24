package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	// fmt.Println(r.Form) // 这些信息是输出到服务器端到打印信息
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])
	var keys string
	var values string
	for k, v := range r.Form {
		keys = keys + "," + k
		values = values + strings.Join(v, ",")
	}
	// fmt.Println(w, "Hello astaxie!") // 这个写入到w的是输出到客户端的
	// 使用Fprintln否则在浏览器中看不到信息
	// fmt.Fprintln(w, "Hello astaxie!")
	fmt.Fprintln(w, "dddd")
	// fmt.Fprintln(w, values)
}

func main() {
	http.HandleFunc("/", sayHelloName)       // 设置访问的路由
	err := http.ListenAndServe(":9999", nil) // 设置监听端口
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
