package main

import "fmt"

/**
 * 返回值可以被命名，命名可以具有一定意义，文档作用
 * 没有参数的 return 返回结果的当前值，影响代码可读性，不建议使用
 */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
