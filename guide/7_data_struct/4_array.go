package main

import "fmt"

/**
 * 数组的长度是其类型的一部分，因此数组不能改变大小。Go 提供了更加便利的方式来使用数组。
 */
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println(len(a))

}
