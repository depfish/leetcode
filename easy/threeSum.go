package main

import "fmt"

func main() {
	// a + b + c = 0
	//  a + b = -c
	nums := []int{-1, 0, 1, 2, -1, -4}
	threeSum(nums)
}


func threeSum(nums []int) [][]int {
	var res [][]int
	lenNums :=len(nums)
	pointerB :=lenNums -1
	for i:=0;i<lenNums;i++{
		for pointerA:=0;pointerA<lenNums-1;pointerA++{
			if pointerA !=pointerB{
				nums[i]= -(nums[pointerA] + nums[pointerB])
				fmt.Printf("nums[i]=%d nums[pointerA]=%d nums[pointerB]=%d\n",nums[i],nums[pointerA],nums[pointerB])
			}
			pointerB--
		}
	}


	return res


}