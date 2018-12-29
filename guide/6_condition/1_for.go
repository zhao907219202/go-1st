package main

import "fmt"

// 如果没有循环条件，那就是个死循环 例如: for {}
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	j := 0
	sum = 0
	for j < 10 {
		sum += j
		j++
	}
	fmt.Println(sum)
}
