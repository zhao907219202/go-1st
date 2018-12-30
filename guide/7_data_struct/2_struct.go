package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var v = Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	p := &v
	// 和C不同 直接用点访问 不是 ->
	p.X = 1e9
	fmt.Println(v)

}
