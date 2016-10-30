package main

import (
	"log"
)

func generate_sieve(num uint64) []uint64 {
	var out []uint64
	ncap := num
	var x uint64 = 3
	for ; x < ncap; x += 2 {
		out = append(out, x)
	}
	return out
}

func find_largets_prime(num, maxp uint64) uint64 {
	var item uint64 = 0
	sieve := generate_sieve(num)
	var count uint64 = 1
	var x, idx uint64
	var jdx int
	for jdx, x = range sieve {
		if x > item {
			item = x
			count++
			if count == maxp {
				item = x
				break
			}
		}
		if x != 0 {
			for idx = uint64(jdx); idx < uint64(len(sieve)); idx += x {
				if sieve[idx] != 0 {
					sieve[idx] = 0
				}

			}
		}
	}
	log.Print("Nth prime: ", count)

	return item
}

func main() {
	var num, maxp uint64 = 1844674, 10001

	out := find_largets_prime(num, maxp)

	log.Print("anser: ", out)

}
