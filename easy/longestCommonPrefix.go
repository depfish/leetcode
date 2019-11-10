package main

import "fmt"

func main() {
	//                0        1      2     3
	//var strs = []string{"flower", "flow", "flight"}
	var strs2 = [] string{"ab", "a"}
	//strs[0]="ggmm"

	//fmt.Printf("str is %v \n",strs)
	fmt.Println(longestCommonPrefix(strs2))
}

// ["dog","racecar","car"]
// ["abab","aba",""]
// ["","b"]
//  ["aa","a"]
// ["aac","acab","aa","abba","aa"]
func longestCommonPrefix(strs []string) string {
	var result string

	var maxLength int
	for n, str := range strs {
		strLen:= len(str)
		if n == 0 {
			maxLength = strLen
			result = str
			continue
		}

		if maxLength == 0 ||strLen == 0 || result[0] != str[0] {
			result = ""
			break
		}

		if len(str) < maxLength {
			maxLength = len(str)
			result=result[:maxLength]
			continue
		}

		for i := 0; i < len(result) && i < strLen; i++ {
			if result[i] != str[i] {
				result = result[:i]
			}
		}
	}
	return result
}
