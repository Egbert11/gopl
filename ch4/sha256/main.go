package main

import (
	"crypto/sha256"
	"fmt"
)

func main()  {
	s1 := []byte("x")
	s2 := []byte("X")
	c1 := sha256.Sum256(s1)
	c2 := sha256.Sum256(s2)
	fmt.Printf("%d\n%d\n%x\n%x\n%t\n%T\n", s1, s2, c1, c2, c1 == c2, c1)
}
