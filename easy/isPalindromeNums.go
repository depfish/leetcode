package main

import "fmt"

func main() {
	nums := 121
	// 1221
	// 123
	fmt.Println(isPalindrome(nums))

}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var reverse int = 0
	for x > reverse {
		reverse = reverse * 10 + x%10
		x /= 10
	}
	return x ==  reverse || x == reverse/10
}
