package main

import (
	"fmt"
	"math"
)

/**
 * Go 没有类。然而，仍然可以在结构体类型上定义方法。
 * 方法接收者 出现在 func 关键字和方法名之间的参数中。

 * 你可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。
 * 但是，不能对来自其他包的类型或基础类型定义方法。
 */
type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//当 v 是 Vertex 的时候Scale方法没有任何作用。
// `Scale` 修改 `v`。当 v 是一个值（非指针），方法看到的是 Vertex 的副本，并且无法修改原始值。
func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.abs())
}
