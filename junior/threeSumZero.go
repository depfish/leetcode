package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{-1, 0, 1, 2, -1, -4}
	threeSum(nums)
}

func threeSum(nums []int) {
	sort.Ints(nums)
	fmt.Println(nums)
	//res [][]int
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		j := i + 1
		z := len(nums) - 1
		for j < z {
			if nums[i]+nums[j]+nums[z] > 0 {
				z--
			} else if nums[i]+nums[j]+nums[z] < 0 {
				j++
			}

			fmt.Printf("%d %d %d i=%d j=%d \n",nums[i],nums[j],nums[z],i,j)


		}
	}
}
