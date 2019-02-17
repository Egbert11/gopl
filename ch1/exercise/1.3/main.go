package main

import (
	"fmt"
	"strings"
)

var arr = [...]string{"hello", "world", "go"}
var args = arr[:]

func DiyJoin() string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

func StringJoin() string {
	return strings.Join(args, " ")
}


func main(){
	fmt.Println("DiyJoin:", DiyJoin())
	fmt.Println("StringJoin:", StringJoin())
}

/*
$ cd ch1/exercise/1.3
$ go test -bench=BenchmarkDiyJoin
goos: darwin
goarch: amd64
pkg: gopl/ch1/exercise/1.3
BenchmarkDiyJoin-4      10000000               116 ns/op
PASS
ok      gopl/ch1/exercise/1.3   1.302s

$ go test -bench=BenchmarkStringJoin
goos: darwin
goarch: amd64
pkg: gopl/ch1/exercise/1.3
BenchmarkStringJoin-4           20000000                71.7 ns/op
PASS
ok      gopl/ch1/exercise/1.3   1.521s

  */
