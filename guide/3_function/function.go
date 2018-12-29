package main

import "fmt"

/**
 * 1.类型在变量名
 * 2.类型相同 除最后一个 前面的都可以省略
 */
func add(x, y int) int {
	return x + y
}

func minus(x int, y int) int {
	return x - y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Print(minus(65, 21))
}
