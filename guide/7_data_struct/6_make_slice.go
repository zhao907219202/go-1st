package main

import "fmt"

func main() {

	// slice 由函数 make 创建。这会分配一个零长度的数组并且返回一个 slice 指向这个数组
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	// 分片起始索引不为0时cap会随之减少,因为数组指针后移导致容量降低，但是结束为止不会影响cap
	// 如果源slice cap 不足以进行分片，会报越界错误
	// 比如： A cap=4, A[2:5] 就会报错，因为没有index=4不存在

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
