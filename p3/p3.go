package main


import(
	"log"
	"math"
)


func generate_sieve(num int) []int {
	var out []int
	ncap := int(math.Sqrt(float64(num)))
	for x := 3; x < ncap ; x += 2{
		out = append(out, x)
	}
	return out
}

func find_largets_prime(num int, sieve []int) int {
	item := 0
	for jdx, x := range sieve {
		if x > item && num % x == 0{
			item = x
		}
		if x != 0 {
			for idx := jdx; idx < len(sieve); idx += x {
				if item % x == 0 {
					sieve[idx] = 0
				}

			}
		}
	}

	return item
}

func main() {
	num := 600851475143
	sieve := generate_sieve(num)
	value := find_largets_prime(num, sieve)
	log.Printf("Final Value: %v", value)

}
