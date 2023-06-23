package main

import (
	"example.com/productionCalculation"
	"fmt"
	"log"
	"strconv"
)

func main() {
	log.SetPrefix("productionCalculation: ")
	log.SetFlags(0)

	toolsForProduction := []string{"EE", "EG"}
	minutesForProduction, _ := productionCalculation.ProductionTimeCalculation(toolsForProduction, 1000)

	fmt.Printf("It takes " + strconv.Itoa(minutesForProduction) + " minutes to produce the amount")
}
