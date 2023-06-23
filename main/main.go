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
	amountElementsPerPartX, amountElementsPerPartY, amountElementsPerPartZ, _ := productionCalculation.ProductionElementCalculation(toolsForProductionFullRun)
	amountOfElemetsPerShiftX, amountOfElemetsPerShiftY, amountOfElemetsPerShiftZ, _ := productionCalculation.ElementUseForOneShift(toolsForProductionFullRun, 30)

	fmt.Println("#######################################")
	fmt.Println("It takes " + strconv.Itoa(minutesForProductionOfSlowestEquipment) + " minutes to produce the amount")
	fmt.Println("One Run with the given tools take: " + strconv.Itoa(minutesForAFullProductionRun) + " minutes")
	fmt.Println("One Shift can produce " + strconv.Itoa(amountPerShift) + " parts.")
	fmt.Println("Element X is needed for one part: " + strconv.Itoa(amountElementsPerPartX))
	fmt.Println("Element Y is needed for one part: " + strconv.Itoa(amountElementsPerPartY))
	fmt.Println("Element Z is needed for one part: " + strconv.Itoa(amountElementsPerPartZ))
	fmt.Println("Element X is needed for one shift: " + strconv.Itoa(amountOfElemetsPerShiftX))
	fmt.Println("Element Y is needed for one shift: " + strconv.Itoa(amountOfElemetsPerShiftY))
	fmt.Println("Element Z is needed for one shift: " + strconv.Itoa(amountOfElemetsPerShiftZ))

}
