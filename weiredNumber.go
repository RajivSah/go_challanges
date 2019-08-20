package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	nLower, _ := strconv.Atoi(os.Args[1])
	nUpper, _ := strconv.Atoi(os.Args[2])
	for i := nLower; i <= nUpper; i++ {
		divisors := divisor(i)
		if sumSlice(divisors) > i && !subsetSum(divisors, len(divisors), i) {
			fmt.Println(i)
		}
	}
}

func divisor(num int) []int {
	res := []int{1}
	for i := 2; i <= int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			res = append(res, i, num/i)
		}
	}
	return res
}

func subsetSum(slice []int, length int, sum int) bool {
	if sum == 0 {
		return true
	} else if length == 0 || sum < 0 {
		return false
	}
	return subsetSum(slice, length-1, sum) || subsetSum(slice, length-1, sum-slice[length-1])
}

func sumSlice(slice []int) int {
	sum := 0
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	return sum
}
