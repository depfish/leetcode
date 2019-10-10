package main

import (
	"fmt"
	"math"
)

func main() {
	var num int = 15342364690
	fmt.Println(reverse(num))
	fmt.Println(math.MaxInt32)
}
// 321=3*100 + 2*10 +1*1
func reverse(x int) int{
	//      maxInt32= 2147483647
	var maxInt int32= math.MaxInt32
	//                1534236469
	var tmp int
	var res int = 0
	flag :=x
	if x > int(maxInt) || x < -int(maxInt) {
		return 0}
    if x < 0 {
    	x=-x
	}
	for x>0{
		res *=10
		tmp =x%10
		x /=10
		res += tmp
	}

	if flag < 0 {return  -res}

	return res

}
