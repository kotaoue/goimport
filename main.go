package main

import (
	"log"

	// _ "github.com/kotaoue/goimport/packages/car"
	"github.com/kotaoue/goimport/packages/car"
	"rsc.io/quote"
)

func main() {
	log.Println("Start")
	log.Println(quote.Hello())
	log.Println(car.ignition())
	log.Println("End")
}
