package main

import (
	"log"
)

func main(){
	log.Print("Hello")
	total := 0
	for i := 3; i < 1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			total += i
		}
	}
	log.Print(total)
}

