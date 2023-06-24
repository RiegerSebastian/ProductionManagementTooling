package productionCalculation

import (
	"example.com/equipmentData"
	"fmt"
	"strconv"
)

func productionTimeCalculation(usedTools []string, productionAmount int) ([]int, error) {
	var timeListForAllEquiptments []int

	for _, tool := range usedTools {
		timeForGivenParts, _ := calculateProductionTimeForTool(tool, productionAmount)
		timeListForAllEquiptments = append(timeListForAllEquiptments, timeForGivenParts)
	}
	return timeListForAllEquiptments, nil
}

// ProductionTimeCalculationForSlowestTool
// Given a list of tools to use in order. How many parts to produce.
// Return days need for production
func ProductionTimeCalculationForSlowestTool(usedTools []string, productionAmount int) (int, error) {
	timeListForAllEquiptments, _ := productionTimeCalculation(usedTools, productionAmount)
	return getMaxValue(timeListForAllEquiptments), nil
}

// ProductionTimeCalculationForAFullRun
// Given a list of tools to use in order. How many parts to produce.
// Return days need for production
func ProductionTimeCalculationForAFullRun(usedTools []string) (int, error) {
	timeListForAllEquiptments, _ := productionTimeCalculation(usedTools, 1)
	return getSumValue(timeListForAllEquiptments), nil
}

func ProductionTimeCalculationForAFullShift(usedTools []string, breakTimeInMinutes int) (int, error) {

	timeForOnePart, _ := ProductionTimeCalculationForAFullRun(usedTools)
	return ((8 * 60) - breakTimeInMinutes) / timeForOnePart, nil
}

func ElementUseForOneShift(usedTools []string, breakTimeInMinutes int) (int, int, int, error) {
	amountElementsPerPartX, amountElementsPerPartY, amountElementsPerPartZ, _ := ProductionResourceCalculation(usedTools)
	productedPartsInOneShift, _ := ProductionTimeCalculationForAFullShift(usedTools, breakTimeInMinutes)
	return amountElementsPerPartX * productedPartsInOneShift, amountElementsPerPartY * productedPartsInOneShift, amountElementsPerPartZ * productedPartsInOneShift, nil
}

// Given: Tool and amount to produce
// Return: Time in days how long production needs
func calculateProductionTimeForTool(tool string, amountOfParts int) (int, error) {
	timeForOnePart, _, _ := equipmentData.GetProductionTimeForEquipment(tool)
	amountOfEquipments, _, _ := equipmentData.GetAmountOfEquipments(tool)
	timeForGivenParts := (timeForOnePart / amountOfEquipments) * amountOfParts

	return timeForGivenParts, nil
}

func getMaxValue(values []int) int {
	max := values[0] //assign the first element equal to max
	for _, number := range values {
		if number > max {
			max = number
		}
	}

	return max
}

func getSumValue(values []int) int {
	sum := values[0]
	for _, number := range values {
		sum = sum + number
	}

	return sum
}

func ProductionResourceCalculation(usedTools []string) (int, int, int, error) {
	var totalElementX int
	var totalElementY int
	var totalElementZ int

	for _, tool := range usedTools {
		elementXForTool, elementYForTool, elementZForTool, _, _ := equipmentData.GetResourceUsageForEquipment(tool)
		totalElementX = totalElementX + elementXForTool
		totalElementY = totalElementY + elementYForTool
		totalElementZ = totalElementZ + elementZForTool
	}
	return totalElementX, totalElementY, totalElementZ, nil
}

func MinimumProductionPerDayForOrder(orders map[string][]string) {
	toolsForPA := []string{"EE", "EF", "EG", "EE"}
	// toolsForPB := []string{"EH", "EF", "EG"}
	// toolsForPC := []string{"EE", "EH", "EG"}

	for _, order := range orders {
		capacityForEquipmentForOrder(order, toolsForPA)
	}
}

func ordersForProduct(orders map[string][]string, product string) {

}

// order: product, amount, time
// toolsForProduct: list of tools to use
func capacityForEquipmentForOrder(order []string, toolsForProduct []string) map[string]int {
	toolsWithLoad := make(map[string]int)
	amount, _ := strconv.Atoi(order[1])
	for _, tool := range toolsForProduct {
		toolsWithLoad[tool] = toolsWithLoad[tool] + amount
	}
	fmt.Println(toolsWithLoad)
	return toolsWithLoad
}
