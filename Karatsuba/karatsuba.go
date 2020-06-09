package main

import (
	"fmt"
	"math"
)

// Karatsuba implements the karatsuba multiplication algorithm
func Karatsuba(x, y int64) int64 {
	if x < 0 {
		x = -x
	}

	if y < 0 {
		y = -y
	}

	if x == 0 || y == 0 {
		return 0
	}

	// define a base case that allows escape from the recursive call
	if x < 10 || y < 10 {
		return x * y
	}

	// calculating the size of the numbers
	xs := CountDigits(x)
	ys := CountDigits(y)
	size := math.Min(float64(xs), float64(ys))
	mid := uint(math.Floor(size / 2))

	// spliting the numbers based on the size calculated above
	x1, x0 := SplitDigits(x, mid)
	y1, y0 := SplitDigits(y, mid)

	// the three recursive calls to subproblems
	z0 := Karatsuba(x0, y0)
	z1 := Karatsuba((x0 + x1), (y0 + y1))
	z2 := Karatsuba(x1, y1)

	answer := (float64(z2)*math.Pow(10, float64((mid*2))) + (float64((z1-z2-z0))*math.Pow(10, float64(mid)))*float64(z0))

	return int64(answer)
}

// CountDigits returns how many digits are in a number
func CountDigits(i int64) (count int64) {
	num := int(i)
	var result int64

	if num == 0 {
		return 1
	}

	if num < 0 {
		num = -num
	}

	for num > 0 {
		result++
		num = num / 10
	}

	return result
}

// SplitDigits splits a sequence of numbers at the mid
func SplitDigits(num int64, digits uint) (int64, int64) {
	divisor := int64(math.Pow(10, float64(num)))

	if num >= divisor {
		return num / divisor, num % divisor
	}

	return 0, num
}

func main() {
	// arguments := os.Args

	// if len(arguments) != 3 {
	// 	fmt.Println("Invalid amount of arguments, Usage: karatsuba x y")
	// 	return
	// }

	// x, err := strconv.ParseInt(arguments[1], 10, 64)
	// if err != nil {
	// 	fmt.Println("Not a number:", arguments[1])
	// }

	// y, err := strconv.ParseInt(arguments[2], 10, 64)
	// if err != nil {
	// 	fmt.Println("Not a number:", arguments[2])
	// }

	// fmt.Println(Karatsuba(x, y))

	fmt.Println(CountDigits(100))
	fmt.Println(SplitDigits(123, 1))
	fmt.Println(Karatsuba(1234, 5678))
	fmt.Println(Karatsuba(1234, -5678))
	fmt.Println(Karatsuba(100, 100))
}
