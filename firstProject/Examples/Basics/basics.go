package basics

import (
	"bytes"
	helper "firstProject/Examples/Helper"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
)

func StringConcatBytes() {
	defer helper.FunctionTimer("StringConcatBytes")
	fmt.Println("\n\nBytes-os string összerakás")

	var buffer bytes.Buffer

	for i := 0; i < 100000; i++ {
		buffer.WriteString(strconv.Itoa(i))
	}

	fmt.Printf("Length of the string: %v\n", len(buffer.String()))
}

func StringConcatBytesConcurrent() {
	defer helper.FunctionTimer("StringConcatBytesConcurrent")
	fmt.Println("\n\nConcurrent Bytes-os string összerakás")
	var buffer bytes.Buffer
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(num int) {
			buffer.WriteString(strconv.Itoa(num))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("Length of the string: %v\n", len(buffer.String()))
}

func StringConcatStringBuilder() {
	defer helper.FunctionTimer("StringConcatStringBuilder")
	fmt.Println("\n\nStringBuilder-es string összerakás")

	var sb strings.Builder
	for i := 0; i < 100000; i++ {
		sb.WriteString(strconv.Itoa(i))
	}
	fmt.Printf("Length of the string: %v\n", len(sb.String()))
}

func Maps() {
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

// FizzBuzz prints Fizz and Buzz according to the value of the number
func FizzBuzz(num int) {
	for i := 1; i <= num; i++ {
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
}

//CALLING PARAMETERS
// n := 123456789
// a(n)

// IntLen receives an int and returns the number of characters it contains.
func IntLen(num int) int {
	last := num % 10
	len := 0
	for num > 10 {
		fmt.Println(num % 10)
		num = num / 10
		fmt.Println(num)
		len++
	}
	first := num
	fmt.Printf("last: %d , first: %d", last, first)
	fmt.Printf("\numEvenum-enumded? %t\num", last == first)
	return len
}

// MaxValue gets an int array and returns the max value
// of the array using a simple for loop
func MaxValue(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// MaxValue2 receives an array and returns the max value
// of the array using range with discard.
// Discard means you will not use the value
func MaxValue2(nums []int) int {
	max := nums[0]
	for _, value := range nums[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// MaxValueIndex receives an array and returns the index
// of the max value of the array
func MaxValueIndex(nums []int) int {
	max := nums[0]
	maxIndex := 0
	for index, value := range nums[1:] {
		if value > max {
			max = value
			maxIndex = index
		}
	}
	return maxIndex
}

// CALLING PARAMETERS
// var limit big.Int
// limit.Exp(big.NewInt(10), big.NewInt(99), nil)

// BigIntLen receives a big.Int and returns the number of characters
// it contains.
// WIP
func BigIntLen(num *big.Int) int {
	// var oszto big.Int
	fmt.Println(num)
	oszto := big.NewInt(10)
	fmt.Println(oszto)
	last := num.Mod(num, oszto)
	length := 0
	for num.Cmp(oszto) == 1 {
		fmt.Println(num.Mod(num, oszto))
		num.Div(num, oszto)
		// num = num.Div(num, oszto)
		fmt.Println(num)
		length++
	}
	first := num
	fmt.Printf("last: %d , first: %d", last, first)
	fmt.Printf("\nEven-ended? %t\n", last == first)
	return length
}
