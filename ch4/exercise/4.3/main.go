package main

import "fmt"

func reverse(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}


// not by pointer, the result will not reverse, except return a new arry`
func reverse2(s [5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}


func main() {
	x := [5]int{1,2,3,4,5}
	reverse(&x)
	fmt.Println(x)
	reverse2(x)
	fmt.Println(x)
}