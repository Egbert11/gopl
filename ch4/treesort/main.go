package main

import "fmt"

type tree struct {
	value int
	left, right *tree
}

func Sort(values []int) {
	var t *tree
	for _, val := range values {
		t = add(t, val)
	}
	appendValue(values[:0], t)
}

func appendValue(values []int, t *tree) []int {
	if t != nil {
		values = appendValue(values, t.left)
		values = append(values, t.value)
		values = appendValue(values, t.right)
	}
	return values
}


func add(t *tree, val int) *tree  {
	if t == nil {
		t = new(tree)
		t.value = val
		return t
	}
	if  t.value >= val {
		left := add(t.left, val)
		t.left = left
	} else {
		right := add(t.right, val)
		t.right = right
	}
	return t
}


func main() {
	values := []int{10,2,34,2,11,45,22,52,12}
	Sort(values)
	fmt.Println(values)
}