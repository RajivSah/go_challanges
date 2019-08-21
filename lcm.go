package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	numbers := os.Args[1:]
	res := map[int]int{}
	for _, num := range numbers {
		numInt, _ := strconv.Atoi(num)
		factors := primeFactors(numInt)
		for k, v := range factors {
			if res[k] < v {
				res[k] = v
			}
		}
	}
	result := 1
	for k, v := range res {
		result *= int(math.Pow(float64(k), float64(v)))
	}
	fmt.Println(result)
}

func primeFactors(num int) map[int]int {
	res := map[int]int{}
	for i := 2; i <= int(math.Pow(float64(num), 2.0))+1; i++ {
		if !isPrime(i) {
			continue
		}
		for num%i == 0 {
			res[i]++
			num = num / i
		}
	}
	return res
}

func isPrime(num int) bool {
	if num == 2 {
		return true
	}
	for i := 2; i <= int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
