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

	/*
		slice 可以重新切片，创建一个新的 slice 值指向相同的数组。

		表达式

		s[lo:hi]
		表示从 lo 到 hi-1 的 slice 元素，含两端。因此

		s[lo:lo]
		是空的，而

		s[lo:lo+1]
		有一个元素。
	*/
	fmt.Println(p[1:4])
	fmt.Println(p)
	p = p[1:4]
	fmt.Println(p)

	/*
		slice 由函数 make 创建。这会分配一个零长度的数组并且返回一个 slice 指向这个数组：

		a := make([]int, 5)  // len(a)=5
		为了指定容量，可传递第三个参数到 `make`：

		注：cap函数返回分配的空间大小
	*/
	a:=make([]int,5)
	printSlice("a",a)
	b:=make([]int , 0, 5)
	printSlice("b",b)
	c := b[:2]
	printSlice("c",c)

	d:= c[2:5]
	printSlice("d",d)

	/*
		slice 的零值是 `nil`。

		一个 nil 的 slice 的长度和容量是 0。
	*/
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

	/**
	向 slice 添加元素是一种常见的操作，因此 Go 提供了一个内建函数 `append`。 内建函数的文档对 append 有详细介绍。
	
	append 的第一个参数 s 是一个类型为 T 的数组，其余类型为 T 的值将会添加到 slice。

	append 的结果是一个包含原 slice 所有元素加上新添加的元素的 slice。

	如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。 返回的 slice 会指向这个新分配的数组。
	**/
	var f []int
	printSlice("f", f)

	// append works on nil slices.
	f = append(f, 0)
	printSlice("f", f)

	// the slice grows as needed.
	f = append(f, 1)
	printSlice("f", f)

	// we can add more than one element at a time.
	f = append(f, 2, 3, 4)
	printSlice("f", f)

	//for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。
	for i , v := range f {
		println("2 ** %d = %d",i, v)
	}


	/**
	可以通过赋值给 _ 来忽略序号和值。

	如果只需要索引值，去掉“, value”的部分即可。
	**/
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func printSlice(s string,x []int){
	fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)
}