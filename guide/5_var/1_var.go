package main

import "fmt"

// var 可以定义在包或者函数级别 类型在后面
var c, python, java bool = true, false, true
var a, b = 10, "hello"

func main() {
	var i int = 100
	// 短声明不能用于函数外
	j := 1
	fmt.Println(i, j, c, python, java, a, b)
}
