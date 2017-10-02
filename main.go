package main

import (
	"fmt"
	"strconv"
	"strings"
)

// IsDevideByThree return true if number can devide by three
func IsDevideByThree(number int) bool {
	return number%3 == 0
}

// IsDevideByFive return true if number can devide by five
func IsDevideByFive(number int) bool {
	return number%5 == 0
}

// IsContainThree return true if word contain Three
func IsContainThree(number int) bool {
	numberSting := strconv.Itoa(number)

	return strings.Contains(numberSting, "3")
}

// IsContainFive return true if word contain Five
func IsContainFive(number int) bool {
	numberSting := strconv.Itoa(number)

	return strings.Contains(numberSting, "5")
}

// FizzBuzz return Fizz if number can devide by three or contain a three
// FizzBuzz return Buzz if number can devide by five or contain a five
func FizzBuzz(number int) string {
	if IsDevideByThree(number) || IsContainThree(number) {
		return "Fizz"
	} else if IsDevideByFive(number) || IsContainFive(number) {
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
