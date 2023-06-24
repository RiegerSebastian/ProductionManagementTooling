package equipmentData

// Given: Tool
// Return: Time need for one part in minutes
func GetProductionTimeForEquipment(tool string) (int, bool, error) {
	toolThroughputConfig := map[string]int{
		"EE": 3,
		"EF": 12,
		"EG": 1,
		"EH": 2,
	}
	productionTimeInMinutes, ok := toolThroughputConfig[tool]
	return productionTimeInMinutes, ok, nil
}

func GetResourceUsageForEquipment(tool string) (int, int, int, bool, error) {
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

func GetAmountOfEquipments(tool string) (int, bool, error) {
	toolAmountConfig := map[string]int{
		"EE": 1,
		"EF": 1,
		"EG": 1,
		"EH": 1,
	}
	productionEquipmentsAmount, ok := toolAmountConfig[tool]
	return productionEquipmentsAmount, ok, nil
}
