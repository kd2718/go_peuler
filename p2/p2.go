package main

import "log"

func main() {
	total := 0
	var i, j int
	for i, j = 1, 2; j < 4000000; i, j = j, i + j {
		if j % 2 == 0 {
			total += j
		}
	}
	log.Print("Final")
	defer log.Print(total)
	log.Print("OK")
	//log.Print(j)

}
