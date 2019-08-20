package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	visited := []int{num}
	for {
		num = sumOfSquare(splitNumber(num))
		if num == 1 {
			fmt.Println("Happy Number")
			return
		} else if contains(visited, num) {
			fmt.Println("Sad Number")
			return
		}
		visited = append(visited, num)
	}
}

func splitNumber(num int) []int {
	var res []int
	for num > 0 {
		remainder := num % 10
		num = num / 10
		res = append(res, remainder)
	}
	return res
}

func sumOfSquare(slice []int) int {
	sum := 0
	for _, n := range slice {
		sum += n * n
	}
	return sum
}

func contains(slice []int, item int) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
