package main

import (
	"testing"
)

func BenchmarkDiyJoin(b *testing.B){
	for i := 0; i < b.N; i++ {
		DiyJoin()
	}
}

func BenchmarkStringJoin(b *testing.B){
	for i := 0; i < b.N; i++ {
		StringJoin()
	}
}