/*
If we list all the natural numbers below 10 that are multiples of 3 or
5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/
package main

import "fmt"

func main() {
	result := SumOfMultiples(1000)
	fmt.Println(result)
}

// SumOfMultiples finds the sum of all the multiple of 3 or 5 below the
// given number.
func SumOfMultiples(below_num int) (sum int) {
	for i := 1; i < below_num; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

// TriangleNumber returns the nth triangle number.
// For more info see: https://en.wikipedia.org/wiki/Triangular_number
func TriangleNumber(nth_number int) int {
	return (nth_number * (nth_number + 1)) / 2
}
