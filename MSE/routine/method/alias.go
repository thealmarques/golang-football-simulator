package method

import (
	"math/rand"
	"time"
)

var probability []float32
var alias []int

func New(input_weights []float32) {
	// temporary
	underfull := []int{}
	overfull := []int{}

	probability = make([]float32, len(input_weights))
	alias = make([]int, len(input_weights))

	// Threshold limit
	threshold := float32(1) / float32(len(input_weights))

	// Fill the temporary tables
	for i, probability := range input_weights {
		if probability >= threshold {
			overfull = append(overfull, i)
		} else {
			underfull = append(underfull, i)
		}
	}

	for len(overfull) > 0 && len(underfull) > 0 {
		underfullIndex := underfull[len(underfull)-1]
		underfull = underfull[:len(underfull)-1]

		overfullIndex := overfull[len(overfull)-1]
		overfull = overfull[:len(overfull)-1]

		// Split the probability in each index
		probability[underfullIndex] = input_weights[underfullIndex] * float32(len(input_weights))
		alias[underfullIndex] = overfullIndex

		// Decrease the probability of the overfull value by the correct ammount
		input_weights[overfullIndex] = (input_weights[overfullIndex] + input_weights[underfullIndex]) - threshold

		if input_weights[overfullIndex] >= threshold {
			overfull = append(overfull, overfullIndex)
		} else {
			underfull = append(underfull, overfullIndex)
		}
	}

	// If the list has values, we should add probability 1
	for len(overfull) != 0 {
		index := overfull[len(overfull)-1]
		overfull = overfull[:len(overfull)-1]
		probability[index] = 1
	}

	// If the list has values, we should add probability 1
	for len(underfull) != 0 {
		index := underfull[len(underfull)-1]
		overfull = underfull[:len(underfull)-1]
		probability[index] = 1
	}
}

func Generate() int {
	rand.Seed(time.Now().UnixNano())

	column := rand.Intn(len(probability))

	// generate biased coin flip
	var coinToss bool
	if rand.Float32() < probability[column] {
		coinToss = true
	} else {
		coinToss = false
	}

	if coinToss {
		return column
	} else {
		return alias[column]
	}
}
