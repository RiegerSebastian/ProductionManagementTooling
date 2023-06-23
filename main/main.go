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

	toolsForProduction := []string{"EE", "EF", "EG"}
	toolsForProductionFullRun := []string{"EE", "EF", "EG", "EE", "EH", "EE", "EG"}
	minutesForProductionOfSlowestEquipment, _ := productionCalculation.ProductionTimeCalculationForSlowestTool(toolsForProduction, 1000)
	minutesForAFullProductionRun, _ := productionCalculation.ProductionTimeCalculationForAFullRun(toolsForProductionFullRun)
	amountPerShift, _ := productionCalculation.ProductionTimeCalculationForAFullShift(toolsForProductionFullRun, 30)
	fmt.Println("#######################################")
	fmt.Println("It takes " + strconv.Itoa(minutesForProductionOfSlowestEquipment) + " minutes to produce the amount")
	fmt.Println("One Run with the given tools take: " + strconv.Itoa(minutesForAFullProductionRun) + " minutes")
	fmt.Println("One Shift can produce " + strconv.Itoa(amountPerShift) + " parts.")

}
