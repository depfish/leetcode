package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := 123
	// 1221
	// 123
	fmt.Println(isPalindrome(nums))

}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x !=0) {
		return false
	}
	str := strconv.Itoa(x)
	var rev string
	for _, v := range str {
		rev =  string(v) + rev
	}
	return str == rev
}
