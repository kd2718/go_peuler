package main

import "log"

func main() {

	target := 1000
	var a, b, c int = 998, 1, 1
	var val int
	var found bool

	for a = 998; a > 0; a-- {
		for b = 1; b <= a; b++ {
			for c = target - a - b; c > 0; c-- {
				if a+b+c == target {
					if a*a+b*b == c*c || a*a == b*b+c*c || a*a+c*c == b*b {
						val = a * b * c
						found = true
						break
					}
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}

	log.Print("number: ", val)

}
