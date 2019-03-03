package main

import "fmt"

func remove_duplicate(strings []string) []string {
	if len(strings) <= 1 {
		return strings
	}
	i := 0
	for _, s := range strings[1:] {
		if s != strings[i] {
			i++
			strings[i] = s
		}
	}
	strings = strings[:i+1]
	return strings
}

func main()  {
	strings := []string{"hello","hello","world","huang","huang","world"}
	strings = remove_duplicate(strings)
	fmt.Println(strings)
}