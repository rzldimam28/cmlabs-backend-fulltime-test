package fizzbuzz_test

import (
	"fmt"
	"testing"
)

func FizzBuzz(num int) {
	for i := 1; i <= num; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	FizzBuzz(100)
}