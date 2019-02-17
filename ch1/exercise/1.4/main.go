package main

import (
	"bufio"
	"fmt"
	"os"
)

// go build gopl/ch1/exercise/1.4
// ./1.4 gopl/ch1/exercise/1.4/file1 xxx
func main(){
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}


func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan() {
		// Just modify here
		counts[input.Text() + " from " + f.Name()]++
	}
}