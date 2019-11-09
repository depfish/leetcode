package main

import "fmt"

func main() {
	//IV
	fmt.Println(romanToInt("XXVII"))

}

func romanToInt(s string) int {
	mp := make(map[string]int)
	mp["I"] = 1
	mp["V"] = 5
	mp["X"] = 10
	mp["L"] = 50
	mp["C"] = 100
	mp["D"] = 500
	mp["M"] = 1000
	var sum int
	for i := 0; i < len(s); i++ {
		if i+1 == len(s) {
			sum += mp[string(s[i])]
			break
		}
		if i+1 < len(s) && mp[string(s[i])] >= mp[string(s[i+1])] {
			sum += mp[string(s[i])]
			continue
		}

		sum -= mp[string(s[i])]

	}

	return sum
}
