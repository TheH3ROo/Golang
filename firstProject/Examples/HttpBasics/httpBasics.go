package httpBasics

import (
	helper "firstProject/Examples/Helper"
	"fmt"
	"net/http"
	"sync"
)

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

func SiteSerial(urls []string) {
	defer helper.FunctionTimer("SiteSerial")()
	for _, url := range urls {
		returnType(url)
	}
}

func SitesConcurrent(urls []string) {
	defer helper.FunctionTimer("SitesConcurrent")()
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
