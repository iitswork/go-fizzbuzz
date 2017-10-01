package main

import (
	"fmt"
	"strconv"
	"strings"
)

// DevideBy return true if number can devide by devider
func DevideBy(number int, devider int) bool {
	return number%devider == 0
}

// Contain return true if word contain innumber
func Contain(number int, word int) bool {
	numberSting := strconv.Itoa(number)

	wordString := strconv.Itoa(word)

	return strings.Contains(numberSting, wordString)
}

// FizzBuzz return Fizz if number can devide by three
// FizzBuzz return Buzz if number can devide by five
func FizzBuzz(number int) string {
	if DevideBy(number, 3) || Contain(number, 3) {
		return "Fizz"
	} else if DevideBy(number, 5) || Contain(number, 5) {
		return "Buzz"
	} else {
		return strconv.Itoa(number)
	}
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
