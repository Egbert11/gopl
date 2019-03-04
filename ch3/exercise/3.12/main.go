package main

import (
	"fmt"
	"os"
	"reflect"
)

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	rmap1 := make(map[rune]int)
	rs := []rune(s1)
	for _, r := range rs {
		rmap1[r] += 1
	}

	rmap2 := make(map[rune]int)
	rs = []rune(s2)
	for _, r := range rs {
		rmap2[r] += 1
	}

	return reflect.DeepEqual(rmap1, rmap2)
}

func main()  {
	if len(os.Args) < 3 {
		fmt.Println("please give two argument strings")
	} else {
		fmt.Println(anagram(os.Args[1], os.Args[2]))
	}

}