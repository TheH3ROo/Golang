package main

import (
	httpBasics "firstProject/Examples/HttpBasics"
)

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}
	httpBasics.SiteSerial(urls)
	httpBasics.SitesConcurrent(urls)
}
