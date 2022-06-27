package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func main() {
	// szoveg := "Ez egy nagyon hosszú szöveg, amiben egy szó többször is előfordul. Többször mint egy, és a kis nagy betűket is figyelembe kell venni"
	//szavakra bont egy stringet, és tömböt ad vissza
	// strings.Fields(szoveg)

	//String konkatenációk
	// fmt.Println("Alap string összerakás")
	// var str1 string
	// // watch := stopwatch.Start()
	// start := time.Now()
	// for i := 0; i < 100000; i++ {
	// 	str1 += strconv.Itoa(i)
	// }
	// stop := time.Now()
	// // watch.Stop()
	// // fmt.Println(str1)
	// fmt.Println(len(str1))
	// // fmt.Printf("Milliseconds elapsed: %s\n", watch.Milliseconds())
	// // fmt.Printf("Milliseconds elapsed: %v\n", watch.String())
	// fmt.Printf("UJ Milliseconds elapsed: %v\n", stop.Sub(start))
	// fmt.Println("\n=============================")
	// fmt.Println("\n8ms teszt")
	// watch2 := stopwatch.Start()
	// start = time.Now()
	// time.Sleep(8 * time.Millisecond)
	// stop = time.Now()
	// watch2.Stop()
	// // fmt.Printf("Milliseconds elapsed: %v\n", watch2.Milliseconds())
	// // fmt.Printf("Milliseconds elapsed: %v\n", watch2.String())
	// fmt.Printf("UJ Milliseconds elapsed: %v\n", stop.Sub(start))

	fmt.Println("\n=============================")
	fmt.Println("\nBytes-os string összerakás")

	var buffer bytes.Buffer
	start := time.Now()
	for i := 0; i < 100000; i++ {
		buffer.WriteString(strconv.Itoa(i))
	}
	stop := time.Now()
	fmt.Printf("UJ Milliseconds elapsed: %v\n", stop.Sub(start))
	fmt.Println(len(buffer.String()))

	// fmt.Println("\n=============================")
	// fmt.Println("\nConcurrent Bytes-os string összerakás")
	// var buffer2 bytes.Buffer
	// var wg sync.WaitGroup
	// start2 := time.Now()
	// for i := 0; i < 100000; i++ {
	// 	wg.Add(1)
	// 	go func(num int) {
	// 		buffer2.WriteString(strconv.Itoa(num))
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()
	// fmt.Printf("UJ Milliseconds elapsed: %v\n", time.Since(start2))
	// fmt.Println(len(buffer2.String()))

	//StringBuilder konkatenáció
	// fmt.Println("\n=============================")
	// fmt.Println("\nStringBuilder-es string összerakás")
	//
	// var sb strings.Builder
	// // watch4 := stopwatch.Start()
	// start = time.Now()
	// for i := 0; i < 100000; i++ {
	// 	sb.WriteString(strconv.Itoa(i))
	// }
	// stop = time.Now()
	// // watch4.Stop()
	// fmt.Println(len(sb.String()))
	// // fmt.Printf("Milliseconds elapsed: %s\n", watch4.Milliseconds())
	// // fmt.Printf("Milliseconds elapsed: %s\n", watch4.String())
	// fmt.Printf("UJ Milliseconds elapsed: %v\n", stop.Sub(start))

	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}
	fmt.Println("")

	start3 := time.Now()
	siteSerial(urls)
	fmt.Println(time.Since(start3))

	start3 = time.Now()
	sitesConcurrent(urls)
	fmt.Println(time.Since(start3))
}

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer resp.Body.Close() //destruktor hívás
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func siteSerial(urls []string) {
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

/*Maps
func maps() {
	stocks := map[string]float64{
		"AMZN": 2087.89,
		"GOOG": 2540.85, //minden sor végére kell a ','(vessző)
		"MSFT": 287.70,
	}
	//elemszám
	fmt.Println(len(stocks))
	// visszaadja a keresett kulcshoz tartozó értéket
	//nullát ad vissza ha nem találja a kulcsot
	fmt.Println(stocks["AMZN"])
	value, ok := stocks["GOOG"]
	if !ok {
		fmt.Println("GOOG not found")
		return
	}
	fmt.Println(value)

	//ha szeretnél értéket beállítani
	stocks["MSFT"] = 822.12
	fmt.Println(stocks)

	//ha törölni szeretnél
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	//ha csak a kulcsok kellenek
	for key := range stocks {
		fmt.Println(key)
	}
	//ha a kulcs-érték párok kellenek
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}
}
*/

/* Maximum keresés
nums := []int{16, 8, 42, 4, 23, 15}
max := nums[0]

for _, value := range nums[1:] {
	if value > max {
		max = value
	}
}
fmt.Println(max)
*/

/* Fizzbuzz
for i := 1; i <= 20; i++ {
	if (i%3 == 0) && (i%5 == 0) {
		fmt.Println("fizz buzz")
	} else if i%3 == 0 {
		fmt.Println("fizz")
	} else if i%5 == 0 {
		fmt.Println("buzz")
	} else {
		fmt.Println(i)
	}
}
*/

/* Int karakterszámláló
n := 123456789
a(n)
func a(n int) {
	last := n % 10
	for n > 10 {
		fmt.Println(n % 10)
		n = n / 10
		fmt.Println(n)
	}
	first := n
fmt.Printf("last: %d , first: %d", last, first)
	fmt.Printf("\nEven-ended? %t\n", last == first)
}
*/

/* BigInt karakterszámláló (nem mukodik)
var limit big.Int
limit.Exp(big.NewInt(10), big.NewInt(99), nil)
func b(n *big.Int) {
	// var oszto big.Int
	fmt.Println(n)
	oszto := big.NewInt(10)
	fmt.Println(oszto)
	last := n.Mod(n, oszto)
	for n.Cmp(oszto) == 1 {
		fmt.Println(n.Mod(n, oszto))
		n = n.Div(n, oszto)
		fmt.Println(n)
	}
	first := n
	fmt.Printf("last: %d , first: %d", last, first)
	fmt.Printf("\nEven-ended? %t\n", last == first)
}
*/

/* Max keresés
for i := 1; i < len(nums)-1; i++ {
	if nums[i] > max {
		max = nums[i]
	}
}
*/
