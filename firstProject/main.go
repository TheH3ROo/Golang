package main

import (
	helper "firstProject/Helpers"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}
	fmt.Println("")

	siteSerial(urls)
	sitesConcurrent(urls)
}

func returnType(url string) {
	defer helper.FunctionTimer("returnType")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func siteSerial(urls []string) {
	defer helper.FunctionTimer("siteSerial")()
	for _, url := range urls {
		returnType(url)
	}
}

func sitesConcurrent(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}
