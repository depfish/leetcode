package main

func main() {
	//var nums = []int{2, 7, 11, 15}
	//var nums2 = []int{3, 2, 4}
	//var nums3= []int{3,3,3}
	var nums4=[]int{2,5,5,1,1}
	twoSum(nums4, 10)

}

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	var slice []int
	for i := 0; i < len(nums); i++ {
		for j := 1; j<len(nums); j++ {
			if i != j {
				if target-nums[i] == nums[j] {
					k, v := mp[i]
					if len(mp) == 0 {
						mp[nums[i]] = nums[j]
						slice = append(slice, i, j)
					}
					if v && k != nums[i] && k != nums[j] {
						mp[nums[i]] = nums[j]
						slice = append(slice, i, j)
					}

				}
			}


		}
	}

	return slice

}
