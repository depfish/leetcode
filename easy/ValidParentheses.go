package main

import "fmt"

func main() {
	//str := "{}{"
	str2 :="{[]}"
	fmt.Println(isValid(str2))
}

func isValid(s string) bool {
	strLen := len(s)
	if strLen == 0 {
		return true
	}

	if strLen %2 !=0{
		return false
	}

	brackets := map[rune]rune{ ')' : '(', '}' : '{', ']' : '[' }

	var stack []rune
	for _,char:= range s{
		if char ==brackets[')'] || char == brackets['}'] || char==brackets[']']{
			stack=append(stack,char)
		}else if len(stack) >0 && brackets[char]==stack[len(stack)-1]{
			stack=stack[:len(stack)-1]
		}else {
			return false
		}
	}

	return len(stack) == 0

}
