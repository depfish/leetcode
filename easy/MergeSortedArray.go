package main

func main() {
	var nums1 =make([]int,8)
	nums1[0],nums1[1],nums1[2]=1,2,3
	var nums2 = []int{1, 5, 6,7,8}

	merge(nums1, 3, nums2, 5)

}

//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3

//Output: [1,2,2,3,5,6]
//        [1 2 3 2 5 6]
//        [1 2 3 0 5 6]
//        [1 2 3 2 5 6]
//链接：https://leetcode-cn.com/problems/merge-sorted-array

func merge(nums1 []int, m int, nums2 []int, n int) {
	mPointer:=m
	nPointer:=n

	for mPointer >0 && nPointer >0{
		if nums2[nPointer-1] > nums1[mPointer-1]{
			nums1[mPointer+nPointer-1]=nums2[nPointer-1]
			nPointer--
		}else {
			nums1[mPointer+nPointer-1]=nums1[mPointer-1]
			mPointer--
		}
	}

	if mPointer ==0 && nPointer > 0 {
		for nPointer > 0{
			nums1[mPointer+nPointer-1]=nums2[nPointer-1]
			nPointer--
		}
	}


}
