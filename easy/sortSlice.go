package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 :=[]string{"d","c","b","a"}
	s2 :=[]int {5,4,3,2,1}

	sort.Ints(s2)
	sort.Strings(s1)
	fmt.Println(s1)
	fmt.Println(s2)
}
