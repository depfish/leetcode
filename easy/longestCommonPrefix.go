package main

import "fmt"

func main() {
	//                0        1      2     3
	//var strs = []string{"flower", "flow", "flight"}
	var strs2 = [] string{"","a"}
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

	var maxIndex int
	for n, str := range strs {
		strLen := len(str)
		if n == 0  && strLen > 0{
			maxIndex = strLen - 1
			result = str
			continue
		}
		if strLen == 0  || (result[0] != str[0]) {
			result = ""
			break
		}

		if strLen -1 < maxIndex {
			maxIndex = strLen -1
			result = result[:strLen]
		}

		for i := 0; i <= maxIndex && i < strLen; i++ {
			if result[i] != str[i] {
				maxIndex=i -1
				result = result[:i]
			}
		}
	}
	return result
}
