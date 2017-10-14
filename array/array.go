package main

import (
	"fmt"
)

func main(){
	var array [2]string
	array[0] = "dw"
	array[1] = "dwfef"

	fmt.Println(array)

	p := []int{1,2,4,6,8}
	fmt.Println(p)
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}