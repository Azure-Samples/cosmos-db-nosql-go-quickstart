package main

import (
	"log"
)

func startCosmos(callback func(msg string)) {
	log.Println("You called Cosmos!")
	callback("Thank you!")	
	callback("...you're welcome...")
}