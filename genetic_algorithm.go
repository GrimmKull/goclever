package main

import (
	"fmt"
	"math/rand"
	"time"
)

func oneMax(bitstring []bool) int {
	sum := 0

	for _, val := range bitstring {
		if val {
			sum++
		}
	}

	return sum
}

//generate the population
func randomBitstring(n int) []bool {
	bitstring := make([]bool, n)

	for i, _ := range bitstring {
		if rand.Intn(10) < 5 {
			bitstring[i] = true
		}
	}

	return bitstring
}

func binaryTournament(pop [][]bool, fitness []int) []bool {
	i, j := rand.Intn(len(pop)), rand.Intn(len(pop))

	for i == j {
		j = rand.Intn(len(pop))
	}

	if fitness[i] > fitness[j] {
		return pop[i]
	} else {
		return pop[j]
	}
}

func pointMutation(bitstring []bool, rate float32) []bool {
	child := bitstring

	for i, _ := range bitstring {
		if rand.Float32() < rate {
			child[i] = !bitstring[i]
		}
	}

	return child
}

func crossover(mother, father []bool, rate float32) []bool {
	if rand.Float32() > rate {
		return mother
	}

	point := 1 + rand.Intn(len(mother)-2) //from  1 to n-2 where n = len

	child := append(mother[0:point], father[point:]...)

	return child
}

func reproduce(selected [][]bool, rate_crossover, rate_mutation float32) [][]bool {
	children := make([][]bool, len(selected))
	father := make([]bool, len(selected[0]))

	for i, mother := range selected {
		if i%2 == 0 {
			father = selected[i+1]
		} else {
			father = selected[i-1]
		}

		if i == len(selected)-1 {
			father = selected[0]
		}

		child := crossover(mother, father, rate_crossover)
		child = pointMutation(child, rate_mutation)

		children[i] = child
	}

	return children
}

func selectPopulation(pop [][]bool, fitness []int) [][]bool {
	selected := make([][]bool, len(pop))

	for i, _ := range pop {
		selected[i] = binaryTournament(pop, fitness)
	}

	return selected
}

func search(generations, bits, size int, rate_crossover, rate_mutation float32) []bool {
	pop := make([][]bool, size)
	fitness := make([]int, size)

	best := 0
	fittest := 0

	for i, _ := range pop {
		pop[i] = randomBitstring(bits)

		fitness[i] = oneMax(pop[i])
	}

	for i := 0; i <= generations; i++ {
		selected := selectPopulation(pop, fitness)

		children := reproduce(selected, rate_crossover, rate_mutation)

		for j, child := range children {
			fitness[j] = oneMax(child)

			if fittest < fitness[j] {
				best = j
				fittest = fitness[j]
			}

			if fitness[j] == len(child) {
				fmt.Println("Goal achieved in", i, "generations")
				return child
			}
		}

		pop = children
	}

	return pop[best]
}

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Println(search(2000, 10, 10, 0.5, 0.1))
}
