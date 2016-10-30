package main

import (
	"log"
	. "strconv"
)

func digit_comp(x int) bool {
	test := []rune(Itoa(x))
	out := true
	for i, j := 0, len(test)-1; i < j; i, j = i+1, j-1 {
		if test[j] != test[i] {
			out = false
			break
		}

	}
	return out
}

func main() {

	x := 999
	value := 0
	var temp, tempi, tempj int
	for i := x; i > 0; i--{
		for j := i; j > i/2; j-- {
			temp = i * j
			if temp > value {
				if digit_comp(temp) == true {
					value = temp
					tempi = i
					tempj = j
				}
			}
		}
	}
	log.Printf("Final: %v from %v * %v", value, tempi, tempj)

}
