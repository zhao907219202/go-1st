package main

import "math"

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
