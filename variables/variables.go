package main

import (
	"fmt"
	"math/cmplx"
)

var c, python, java bool

//ToBe 与导入语句一样，变量的定义“打包”在一个语法块中
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	var i int
	//变量定义可以包含初始值，每个变量对应一个。

	//如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。
	var cPlusPlus, goLang = true, "yes！"
	fmt.Println(i, c, python, java, cPlusPlus, goLang)

	//在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。
	// 函数外的每个语句都必须以关键字开始（`var`、`func`、等等），`:=` 结构不能使用在函数外。
	php, ruby := true, "no!"

	fmt.Println(php, ruby)

	/*

			Go 的基本类型有Basic types

		bool

		string

		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte // uint8 的别名

		rune // int32 的别名
		     // 代表一个Unicode码

		float32 float64

		complex64 complex128
	*/
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
