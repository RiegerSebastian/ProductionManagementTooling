package productionCalculation

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
	amountElementsPerPartX, amountElementsPerPartY, amountElementsPerPartZ, _ := ProductionElementCalculation(usedTools)
	productedPartsInOneShift, _ := ProductionTimeCalculationForAFullShift(usedTools, breakTimeInMinutes)
	return amountElementsPerPartX * productedPartsInOneShift, amountElementsPerPartY * productedPartsInOneShift, amountElementsPerPartZ * productedPartsInOneShift, nil
}

// Given: Tool and amount to produce
// Return: Time in days how long production needs
func calculateProductionTimeForTool(tool string, amountOfParts int) (int, error) {
	timeForOnePart, _, _ := getProductionTimeForEquipment(tool)
	amountOfEquipments, _, _ := getAmountOfEquipments(tool)
	timeForGivenParts := (timeForOnePart / amountOfEquipments) * amountOfParts

	return timeForGivenParts, nil
}

// Given: Tool
// Return: Time need for one part in minutes
func getProductionTimeForEquipment(tool string) (int, bool, error) {
	toolThroughputConfig := map[string]int{
		"EE": 3,
		"EF": 12,
		"EG": 1,
		"EH": 2,
	}
	productionTimeInMinutes, ok := toolThroughputConfig[tool]
	return productionTimeInMinutes, ok, nil
}

// Given: Tool
// Return: Time need for one part in minutes
func getElementUsageForEquipment(tool string) (int, int, int, bool, error) {
	toolProductUseConfig := map[string][]int{
		// Equipment: X, Y, Z
		"EE": {2, 0, 0},
		"EF": {1, 1, 0},
		"EG": {0, 0, 2},
		"EH": {2, 1, 1},
	}
	elementUse, ok := toolProductUseConfig[tool]
	return elementUse[0], elementUse[1], elementUse[2], ok, nil
}

func getAmountOfEquipments(tool string) (int, bool, error) {
	toolAmountConfig := map[string]int{
		"EE": 1,
		"EF": 3,
		"EG": 1,
		"EH": 1,
	}
	productionEquipmentsAmount, ok := toolAmountConfig[tool]
	return productionEquipmentsAmount, ok, nil
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

func ProductionElementCalculation(usedTools []string) (int, int, int, error) {
	var totalElementX int
	var totalElementY int
	var totalElementZ int

	for _, tool := range usedTools {
		elementXForTool, elementYForTool, elementZForTool, _, _ := getElementUsageForEquipment(tool)
		totalElementX = totalElementX + elementXForTool
		totalElementY = totalElementY + elementYForTool
		totalElementZ = totalElementZ + elementZForTool
	}
	return totalElementX, totalElementY, totalElementZ, nil
}
