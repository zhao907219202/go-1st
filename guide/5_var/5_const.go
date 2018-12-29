package main

import "fmt"

/**
 * 常量可以是字符、字符串、布尔或数字类型的值。
 * 常量不能使用 := 语法定义。
 */
const PI = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", PI, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
