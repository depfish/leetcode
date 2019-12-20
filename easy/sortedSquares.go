package main

import "fmt"

func main() {
	A :=[]int{-4,-1,0,3,10}
	b:=sortedSquares(A)
	fmt.Println(b)
}


//Input: [-7,-3,2,3,11] --> [-7,-3]  [2,3,11]
//Output: [4,9,9,49,121]    [49,9]   [4,9,121]
//                          [9,49]   [4,9,121]
//[0 0 0 0 0 4 9 9 49]
func sortedSquares(A []int) []int {
	lenA:=len(A)
	sliceB :=make([]int,lenA)
	var flag int
for i:=0;i<lenA;i++{
	if A[i] <0 && A[i+1] >0 && i+1 <=lenA{
		flag=i
		fmt.Printf("max neg index is : %v\n",i)
		break
	}
}

for k,v := range A{
	A[k]=v*v
}

fmt.Println(A)

pointerA :=flag
pointerB :=flag+1
for pointerA >=0 && pointerB <=lenA {
	if A[pointerA] > A[pointerB]{
		sliceB=append(sliceB,A[pointerB])
		pointerB++
	}else {
		sliceB=append(sliceB,A[pointerA])
		pointerA--
	}
	}

if pointerA<0 && pointerB < lenA {
	for pointerB < lenA{
		sliceB =append(sliceB,A[pointerB])
		pointerB++
	}
}


if pointerB == lenA-1 && pointerA >=0 {
	for pointerA >=0 {
		sliceB=append(sliceB,A[pointerA])
		pointerA--
	}
}

return sliceB
}