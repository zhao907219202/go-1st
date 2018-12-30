package main

import "fmt"

func main() {
	// []T 是一个元素类型为 T 的 slice：不指定size的数组就是一个slice
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p == ", p)
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] ==  %d\n", i, p[i])
	}

	// 取头不取尾
	fmt.Println("p[1:4] ==", p[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("p[:3] ==", p[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("p[4:] ==", p[4:])

}
