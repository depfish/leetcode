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

//链接：https://leetcode-cn.com/problems/merge-sorted-array

func merge(nums1 []int, m int, nums2 []int, n int) {
	if nums2[0] > nums1[m-1]{
		for i:=n-1;i>=0;i--{
			nums1[m+i]=nums2[i]
		}

		return
	}

	if nums2[0] < nums1[m-1]{
		for i:=n-1;i>=0;i--{
			if nums2[i] > nums1[m-1]{
				nums1[m+i]=nums2[i]
			}else{
				for k:=0;k<m+n-i;i++{
					if nums2[i] > nums1[k] && nums2[i] < nums1[k+1]{

					}
				}
			}
		}

	}
}
