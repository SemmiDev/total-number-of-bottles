package main

import (
	"fmt"
	"slices"
	"strings"
)

// Seorang petani memiliki 3 ember dengan kapasitas berbeda untuk menampung air.
// Kapasitas setiap ember adalah bilangan prima antara 0 hingga 30 liter (0 < Kapasitas x < 30).
// Petani ingin menampung X liter air (100 < X < 10_000_000) dengan menggunakan ember-ember tersebut.
// Berapa banyak ember dari setiap kapasitas yang dibutuhkan petani agar total jumlah ember yang digunakan paling sedikit?

// Contoh:
// Ember 1 = 5 liter
// Ember 2 = 7 liter
// Ember 3 = 11 liter
// X = 100 liter
// Jawaban: Ember 3 = 9 buah, Ember 1 = 1 buah, Ember 2 = 0 buah, total = 10 buah
// atau Ember 3 = 9 buah, Ember 1 = 0 buah, Ember 2 = 1 buah, total = 10 buah

func minEmber(capacities []int, targetWater int) string {
	bucketSequence := make(map[int]int)
	emberCount := make(map[int]int)

	for i, capacity := range capacities {
		bucketSequence[capacity] = i + 1 // Start from 1
		emberCount[capacity] = 0
	}

	// Sort capacities in ascending order
	slices.Sort(capacities)

	answers := make([]string, 0)

	// Iterate through capacities in descending order
	for i := len(capacities) - 1; i >= 0; i-- {
		capacity := capacities[i]

		previousCapacity := 0
		if i > 0 {
			previousCapacity = capacities[i-1]
		}

		// Fill the bucket as much as possible
		for targetWater >= capacity && targetWater > previousCapacity {
			emberCount[capacity]++
			targetWater -= capacity
		}
	}

	// Add the remaining water to the smallest bucket
	if targetWater > 0 {
		emberCount[capacities[0]]++
	}

	// Create the result string
	for i := len(capacities) - 1; i >= 0; i-- {
		capacity := capacities[i]
		value := emberCount[capacity]

		if value <= 1 {
			answers = append(answers, fmt.Sprintf("Bottle %d = %d bottle", bucketSequence[capacity], value))
		} else {
			answers = append(answers, fmt.Sprintf("Bottle %d = %d bottles", bucketSequence[capacity], value))
		}
	}

	return strings.Join(answers, ", ")
}

func main() {
	capacities := []int{5, 7, 4}
	// 14 * 7 = 98

	targetWater := 100
	totalEmber := minEmber(capacities, targetWater)
	fmt.Println(totalEmber)
}
