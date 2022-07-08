package main

import (
	basics "firstProject/Examples/Basics"
)

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}
	basics.SiteSerial(urls)
	basics.SitesConcurrent(urls)
}
