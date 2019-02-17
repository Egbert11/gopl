package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	times = 2
)

func main() {
	url := os.Args[1]
	fetchAsync(url) // fetch url by goroutine and channel
	fetchSync(url) // fetch url by loop
}

func fetchAsync(url string) {
	start := time.Now()
	ch := make(chan string)
	for i := 0; i < times; i++ {
		go fetch(url, ch)   // start a goroutine
	}
	for i := 0; i < times; i++ {
		fmt.Println(<-ch)   // receive from channel ch
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("fetch async %.2fs seconds elapsed.\n", secs)
}

func fetchSync(url string) {
	start := time.Now()
	for i := 0; i < times; i++ {
		eachStart := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			fmt.Print(err)
			continue
		}
		nbytes, err := io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("while reading %s: %v\n", url, err)
			continue
		}
		secs := time.Since(eachStart).Seconds()
		fmt.Printf("%.2fs %7d %s\n", secs, nbytes, url)
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("fetch sync %.2fs seconds elapsed.\n", secs)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)   // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}