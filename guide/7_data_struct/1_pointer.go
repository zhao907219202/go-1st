package main

import "fmt"

func main() {
	// 没有指针运算 零值为nil
	i, j := 41, 2701

	p := &i
	fmt.Println(*p)

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	var pNil *int
	fmt.Println(pNil)

}
