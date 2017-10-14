package main

import (
	"fmt"
	"runtime"
)

func main(){
	/*
	一个结构体（`struct`）就是一个字段的集合。

除非以 fallthrough 语句结束，否则分支会自动终止。
*/
	fmt.Println("go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.",os)
	}
}