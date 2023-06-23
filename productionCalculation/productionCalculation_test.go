package productionCalculation

import (
	"testing"
)

func TestGetProductionTimeForEquipment(t *testing.T) {
	timeInMinutesForEE, ok, _ := getProductionTimeForEquipment("EE")
	if ok != true {
		t.Fatalf("Equipment EE not found ")
	}
	timeInMinutesForEF, ok, _ := getProductionTimeForEquipment("EF")
	timeInMinutesForEG, ok, _ := getProductionTimeForEquipment("EG")
	timeInMinutesForEH, ok, _ := getProductionTimeForEquipment("EH")
	if timeInMinutesForEE == 0 {
		t.Fatalf("Time for equipment EE is wrong")
	}
	if timeInMinutesForEF == 0 {
		t.Fatalf("Time for equipment EF is wrong")
	}
	if timeInMinutesForEG == 0 {
		t.Fatalf("Time for equipment EG is wrong")
	}
	if timeInMinutesForEH == 0 {
		t.Fatalf("Time for equipment EH is wrong")
	}
}

func TestGetMaxValue(t *testing.T) {
	maxValueTestSet := []int{1, 2, 17, 13, 28, 44, 37, 68}
	maxValue := getMaxValue(maxValueTestSet)

	if maxValue != 68 {
		t.Fatalf("Max Value is wrong")
	}
}
