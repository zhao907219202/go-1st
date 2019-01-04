package main

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

//如果顶级的类型只有类型名的话，可以在文法的元素中省略键名。
var mm = map[string]Vertex2{
	"Google":    Vertex2{37.42202, -122.08408},
	"Bell Labs": {40.68433, -74.39967},
}

/**
 * map 在使用之前必须用 make 而不是 new 来创建；
 * 值为 nil 的 map 是空的，并且不能赋值
 */
func main() {
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{40.68433, -74.39967}

	fmt.Println(m["Bell Labs"])

	fmt.Println(mm)

}
