package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// := 是声明变量并复制
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
