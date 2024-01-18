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

// example: $capacities = [11, 7, 2]; $targetWater = 100;

func minEmber(capacities []int, targetWater int) string {
	bucketSequence := make(map[int]int)
	emberCount := make(map[int]int)

	for i, capacity := range capacities {
		bucketSequence[capacity] = i + 1 // i want to start from 1
		emberCount[capacity] = 0
	}

	// mengurutkan slice dari yang terkecil ke terbesar
	slices.Sort(capacities)

	highestCapacity := capacities[len(capacities)-1]
	secondHighestCapacity := capacities[len(capacities)-2]
	lowestCapacity := capacities[0]

	// range more than secondHighestCapacity
	for targetWater > secondHighestCapacity {
		emberCount[highestCapacity]++
		targetWater -= highestCapacity
	}

	// range more than lowestCapacity
	for targetWater > lowestCapacity {
		emberCount[secondHighestCapacity]++
		targetWater -= secondHighestCapacity
	}

	// range more than 0
	for targetWater > 0 {
		emberCount[lowestCapacity]++
		targetWater -= lowestCapacity
	}

	// show from the highest capacity
	answers := make([]string, 0)
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
	capacities := []int{5, 7, 11}
	targetWater := 150
	// Bootle 3 = 13, 13 * 11 = 143, 150 - 143 = sisa 7
	// Bootle 2 = 1, 1 * 7 = 7, 7 - 7 = sisa 0
	totalEmber := minEmber(capacities, targetWater)
	fmt.Println(totalEmber)
}
