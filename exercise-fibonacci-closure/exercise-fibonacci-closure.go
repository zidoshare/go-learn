package main 

import (
	"fmt"
)

func fibonacci() func()(int,int) {
	v0,v1,v2 := 0,1,0
	num :=0
	return func()(int,int) {
		num++
		if num == 1 {
			return v0,num
		} else if num == 2 {
			return v1,num
		} else {
			v2 = v0 + v1
			v0 = v1
			v1 = v2
			return v0 + v1,num
		}
	}
}

func main(){
	f := fibonacci()

	for i:=0; i<10; i++ {
		fmt.Println(f())
	}
}