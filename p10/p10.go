package main

import "log"

func generate_sieve(num int) []int {
	var out []int
	for x := 3; x < num; x += 2 {
		out = append(out, x)
	}
	return out
}

func get_prime_slice(sieve []int) int {
	sum := 2
	for jdx, x := range sieve {
		if x != 0 {
			sum += x
			for idx := jdx; idx < len(sieve); idx += x {
				sieve[idx] = 0

			}
		}
	}

	return sum
}

func main() {
	num := 2000000
	slice := generate_sieve(num)
	out := get_prime_slice(slice)
	log.Print("anser: ", out)

}
