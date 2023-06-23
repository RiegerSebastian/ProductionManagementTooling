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
