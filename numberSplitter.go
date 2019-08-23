//  go run  numberSplitter.go number parts
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	parts, _ := strconv.Atoi(os.Args[2])
	quotient := num / parts
	remainder := num % parts

	res := []int{}
	for i := 0; i < parts; i++ {
		res = append(res, quotient)
	}
	res[0] += remainder
	fmt.Println(res)
}
