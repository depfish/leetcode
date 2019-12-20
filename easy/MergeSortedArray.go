package main

func main() {
	var nums1 =make([]int,6)
	nums1[0],nums1[1],nums1[2]=1,2,3
	var nums2 = []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)

}

//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3

//Output: [1,2,2,3,5,6]
//        [1 2 3 2 5 6]
//        [1 2 3 0 5 6]
//        [1 2 3 2 5 6]
//链接：https://leetcode-cn.com/problems/merge-sorted-array

func merge(nums1 []int, m int, nums2 []int, n int) {
	j:=m-1
	k:=n-1
	s:=m+n-1
	for i:=k;i>=0;i--{
		if nums1[j] > nums2[i]{
			nums1[s]=nums1[j]
		}
		if nums1[j] <= nums2[i]{
			nums1[s]=nums2[i]
		}

		//if nums1[j]==nums2[i]{
		//	for p:=s;p>=1;p--{
		//		nums1[p]=nums1[p-1]
		//	}
		//	nums1[j+1]=nums2[i]
		//}

		j--
		s--
	}

	//fmt.Println(nums1)
}
