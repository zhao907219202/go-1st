package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16}

func main() {
	// for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow = make([]int, 5)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
