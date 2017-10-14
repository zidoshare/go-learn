package main

import (
	"fmt"
)

//Vert 坐标
type Vert struct {
	x int
	y int
}

var (
	v1 = Vert{10,11}
	v2 = Vert{x:20}
	v3 = Vert{}
	v4 = &Vert{30,40}
)
func main(){
	av := Vert{1,2}
	/*
结构体字段可以通过结构体指针来访问。

通过指针间接的访问是透明的。
	*/
	p := &av
	p.x = 1e9
	fmt.Println(av)
	fmt.Println(av.x)

	fmt.Println(v1,v2,v3,v4,*v4)
}