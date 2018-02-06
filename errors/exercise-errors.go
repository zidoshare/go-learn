package main

import "fmt"

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}

func main() {
	result,err :=Sqrt(2)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}

	result,err = Sqrt(-2)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}

}
