package main

import "fmt"

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

	// 如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。 返回的 slice 会指向这个新分配的数组
	var a []int
	printSlice1("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice1("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice1("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice1("a", a)

}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
