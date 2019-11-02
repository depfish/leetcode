package main

import "testing"

func BenchmarkCommaSplitNum(b *testing.B){
	num := 123456
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		CommaSplitNum(num)
	}
}