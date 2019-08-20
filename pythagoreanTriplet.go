package main

import (
	"fmt"
	"math"
)

func main() {
	for a := 1.0; a <= 1000; a++ {
		for b := a; b <= 1000; b++ {
			if math.Pow(a, 2)+math.Pow(b, 2)-math.Pow(1000-a-b, 2) == 0.0 {
				fmt.Println(a, b, 1000-a-b)
				return
			}
		}
	}
}
