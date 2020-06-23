package main

import (
	"log"

	_ "github.com/kotaoue/goimport/car"
	"rsc.io/quote"
)

func main() {
	log.Println("Start")
	log.Println(quote.Hello())
	// log.Println(car.run())
	log.Println("End")
}
