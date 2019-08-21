package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	res := []int64{1, 5, 6}
	suffix := map[int]int64{0: 5, 1: 6}
	for i := 1.0; i <= 10.0; i++ {
		for k, v := range suffix {
			for j := int64(1); j <= 10; j++ {
				checkNum := j*int64(math.Pow(10.0, i)) + v
				if isCurious(checkNum) {
					res = append(res, checkNum)
					suffix[k] = checkNum
				}
			}
		}
	}
	fmt.Println(res)
}

func isCurious(num int64) bool {
	numLength := len(strconv.FormatInt(num, 10))
	if (num*num)%int64(math.Pow10(numLength)) == num {
		return true
	}
	return false
}
