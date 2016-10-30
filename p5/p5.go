package main


import(
	"log"
)


func get_answer(num int) int {
	out := 1
	done := make(map[int]int)
	for x := 2; x <= num; x ++ {
		done[x] = x
	}
	match := false
	for x := 2; x <= num ; x ++ {
		for j := x; j <= num; j += x {
			match = false
			for k:= j; k <= num; k += x {
				log.Print("k, done k: ", k, done[k])
				if done[k] == 1 {
					log.Print("contineu!!!!!")
					continue
				}
				if done[k] % x == 0 {

					log.Print("current X, j, k: ", x, j, k)
					match = true
					done[k] = done[k] / x
				}
			}
			if match {
				out *= x
			}
			log.Print("current out: ", out)
			log.Print(done)
		}
	}
	return out
}

//func find_largets_prime(num int, sieve []int) {
//	for jdx, x := range sieve {
//		if num % x != 0 {
//			sieve[jdx] = 0
//		}
//	}
//}

func main() {
	num := 20
	total := get_answer(num)
	log.Printf("Final Value: %v", total)

}
