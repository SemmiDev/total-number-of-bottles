package main

import (
	"testing"
)

func TestMinEmber(t *testing.T) {
	tests := []struct {
		capacities  []int
		targetWater int
		expectedStr string
	}{
		{[]int{5, 7, 11}, 150, "Bottle 3 = 13 bottles, Bottle 2 = 1 bottle, Bottle 1 = 0 bottle"},
		{[]int{2, 7, 11}, 150, "Bottle 3 = 13 bottles, Bottle 2 = 1 bottle, Bottle 1 = 0 bottle"},
		{[]int{2, 7, 10}, 150, "Bottle 3 = 15 bottles, Bottle 2 = 0 bottle, Bottle 1 = 0 bottle"},
		{[]int{5, 7, 11}, 100, "Bottle 3 = 9 bottles, Bottle 2 = 0 bottle, Bottle 1 = 1 bottle"},
	}

	for _, test := range tests {
		result := minEmber(test.capacities, test.targetWater)
		if result != test.expectedStr {
			t.Errorf("For capacities %v and targetWater %d, expected %s but got %s", test.capacities, test.targetWater, test.expectedStr, result)
		}
	}
}
