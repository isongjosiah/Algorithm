package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Karatsuba implements the karatsuba multiplication algorithm
func Karatsuba(x, y float64) float64 {
	// define a base case that allows escape from the recursive call
	if x < 10 || y < 10 {
		return x * y
	}

	// calculating the size of the numbers
	xs := CountDigits(x)
	ys := CountDigits(y)
	size := math.Min(xs, ys)
	mid := math.Floor(size / 2)

	// spliting the numbers based on the size calculated above
	x1, x0 := SplitDigits(x, mid)
	y1, y0 := SplitDigits(y, mid)

	// the three recursive calls to subproblems
	z0 := Karatsuba(x0, y0)
	z1 := Karatsuba((x0 + x1), (y0 + y1))
	z2 := Karatsuba(x1, y1)

	answer := (z2*(math.Pow(10, (mid*2))+((z1-z2-z0)*math.Pow(10, mid))) + z0)

	return answer
}

// CountDigits returns how many digits are in a number
func CountDigits(i float64) (count float64) {
	// i = len(strconv.Itoa(i))
	v := len(strconv.FormatFloat(i, 'f', -1, 64))
	return float64(v)
}

// SplitDigits splits a sequence of numbers at the mid
func SplitDigits(digit, mid float64) (x, y float64) {
	a := fmt.Sprintf("%f", digit)
	mid2 := int(mid)
	// figure out how the midpoint works after testing it out
	a1 := a[:mid2]
	a0 := a[mid2:]
	x, _ = strconv.ParseFloat(a1, 64)
	y, _ = strconv.ParseFloat(a0, 64)
	return x, y
}

func main() {
	arguments := os.Args

	if len(arguments) != 3 {
		fmt.Println("Invalid amount of arguments, Usage: karatsuba x y")
		return
	}

	x, err := strconv.ParseFloat(arguments[1], 64)
	if err != nil {
		fmt.Println("Not a number:", arguments[1])
	}

	y, err := strconv.ParseFloat(arguments[2], 64)
	if err != nil {
		fmt.Println("Not a number:", arguments[2])
	}

	fmt.Println(Karatsuba(x, y))

	// fmt.Println(SplitDigits(123, 1))
	// fmt.Println(Karatsuba(1234, 5678))
	// fmt.Println(Karatsuba(1234, -5678))
	// fmt.Println(Karatsuba(100, 100))
}
