package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
* 包 http 通过任何实现了 http.Handler 的值来响应 HTTP 请求：
* package http
* type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
* }
* 在这个例子中，类型 Hello 实现了 `http.Handler`。
*/

type Hello struct {
}

func (h Hello) ServeHTTP(
	w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	// 处理指定path
	http.Handle("/path", h)
	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
