package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15}
	//var nums2 = []int{3, 2, 4}
	//var nums3= []int{3,3,3}
	//var nums4=[]int{2,5,5,1,1}
	fmt.Println(twoSum(nums, 9))

}



// 2 , 7 ,11 ,15
// 3 , 3 , 3
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	res := make([]int, 0, 2)
	var b int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			mp[nums[i]] = i
			continue
		}
		b = target - nums[i]
		_, exist := mp[b]

		if exist {
			res = append(res, mp[b], i)
			break
		}
		mp[nums[i]] = i

	}

	return res

}
