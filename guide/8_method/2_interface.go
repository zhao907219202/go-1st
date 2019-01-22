package main

import (
	"fmt"
	"math"
)

/**
 * 接口中定义方法集合
 */
type Abser interface {
	Abs() float64
}

type MyFloat1 float64

type Vertex1 struct {
	X, Y float64
}

func (f MyFloat1) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat1(-math.Sqrt2)
	v := Vertex1{3, 4}
	a = f //MyFloat1 实现了Abs 即实现了Abser
	fmt.Println(a.Abs())
	a = &v //Vertex1 实现了Abs 即实现了Abser
	fmt.Println(a.Abs())

	//a = v // Vertex1 没有实现，所以执行报错
	//fmt.Println(a.Abs())

}
