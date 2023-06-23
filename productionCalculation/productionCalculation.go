package productionCalculation

// ProductionTimeCalculation
// Given a list of tools to use in order. How many parts to produce.
// Return days need for production
func ProductionTimeCalculation(usedTools []string, productionAmount int) (int, error) {
	var timeListForAllEquiptments []int

	for _, tool := range usedTools {
		timeForGivenParts, _ := calculateProductionTimeForTool(tool, 1000)
		timeListForAllEquiptments = append(timeListForAllEquiptments, timeForGivenParts)
	}
	return getMaxValue(timeListForAllEquiptments), nil
}

// Given: Tool and amount to produce
// Return: Time in days how long production needs
func calculateProductionTimeForTool(tool string, amount int) (int, error) {
	timeForOnePart, _, _ := getProductionTimeForEquipment(tool)
	timeForGivenParts := timeForOnePart * amount

	return timeForGivenParts, nil
}

// Given: Tool
// Return: Time need for one part in minutes
func getProductionTimeForEquipment(tool string) (int, bool, error) {
	toolConfig := map[string]int{
		"EE": 3,
		"EF": 12,
		"EG": 1,
		"EH": 2,
	}
	productionTimeInMinutes, ok := toolConfig[tool]
	return productionTimeInMinutes, ok, nil
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
