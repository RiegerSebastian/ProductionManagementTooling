package main

import (
	"example.com/productionCalculation"
	"fmt"
	"log"
	"strconv"
)

func main() {

	minimumProductionCalculation()
}

func minimumProductionCalculation() {

	// order is a map:
	// key: OrderNumber
	// value: prodcut, amount, days to deliver
	order := make(map[string][]string)
	productionRate := []string{"PA", "1000", "100"}
	order["#111"] = productionRate
	// productionRate = []string{"PB", "100", "100"}
	// order["#222"] = productionRate
	// productionRate = []string{"PC", "500", "10"}
	// order["#333"] = productionRate
	productionRate = []string{"PA", "500", "50"}
	order["#444"] = productionRate
	fmt.Println(order)
	productionCalculation.MinimumProductionPerDayForOrder(order)

}

func basicProductionTest() {
	log.SetPrefix("productionCalculation: ")
	log.SetFlags(0)

	toolsForProduction := []string{"EE", "EF", "EG"}
	toolsForProductionFullRun := []string{"EE", "EF", "EG", "EE", "EH", "EE", "EF"}

	minutesForProductionOfSlowestEquipment, _ := productionCalculation.ProductionTimeCalculationForSlowestTool(toolsForProduction, 1000)
	minutesForAFullProductionRun, _ := productionCalculation.ProductionTimeCalculationForAFullRun(toolsForProductionFullRun)
	amountPerShift, _ := productionCalculation.ProductionTimeCalculationForAFullShift(toolsForProductionFullRun, 30)
	amountElementsPerPartX, amountElementsPerPartY, amountElementsPerPartZ, _ := productionCalculation.ProductionResourceCalculation(toolsForProductionFullRun)
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
