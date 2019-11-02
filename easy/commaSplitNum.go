package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num:=0
	//1000
	//   12,345
	//   1,000,000
	fmt.Println(CommaSplitNum(num))
}

func CommaSplitNum(num int) string{
	var isNegtive bool
	if num < 0 {
		isNegtive = true
		num = - num
	}
	if num < 1000 {
		return strconv.Itoa(num)
	}

	var res string
	for  num > 0 {
		if num < 1000 {
			res = strconv.Itoa(num) + "," + res
			break
		}
			tmp := num % 1000
			num /= 1000
			res = strconv.Itoa(tmp) + "," + res
	}

	if isNegtive {
		return strings.TrimRight("-" + res,",")
	}

	return strings.TrimRight(res,",")
}