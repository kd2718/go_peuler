package main

import "log"

func main() {
	sum_of_squars, square_of_sum := 0, 0
	num := 100
	for x := 1; x <= num; x++{
		sum_of_squars, square_of_sum = sum_of_squars + x * x, square_of_sum + x
	}
	out := square_of_sum * square_of_sum - sum_of_squars
	log.Print("anser: ", out)
}
