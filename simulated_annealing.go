package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func euc_2d(coord1, coord2 City) int {
	diffx := coord1.x - coord2.x
	diffy := coord1.y - coord2.y
	return int(math.Sqrt(float64(diffx*diffx + diffy*diffy)))
}

func cost(permutation []int, cities []City) int {
	distance := 0
	city2 := 0

	for index, city1 := range permutation {
		if index == len(permutation)-1 {
			city2 = permutation[0]
		} else {
			city2 = permutation[index+1]
		}

		distance += euc_2d(cities[city1], cities[city2])
	}

	return distance
}

func randomPermutation(cities []City) []int {
	perm := make([]int, len(cities))

	for i, _ := range perm {
		perm[i] = i
	}

	for i, _ := range perm {
		r := rand.Intn(len(perm)-i) + i
		perm[r], perm[i] = perm[i], perm[r]
	}

	return perm
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func indexOf(s []int, x int) int {
	for i, val := range s {
		if val == x {
			return i
		}
	}

	return -1
}

func stochasticTwoOpt(permutation []int) []int {
	city1, city2 := rand.Intn(len(permutation)), rand.Intn(len(permutation))

	exclude := []int{city1}

	if city1 == 0 {
		exclude = append(exclude, len(permutation)-1)
	} else {
		exclude = append(exclude, city1-1)
	}

	if city1 == len(permutation)-1 {
		exclude = append(exclude, 0)
	} else {
		exclude = append(exclude, city1+1)
	}

	for {
		if indexOf(exclude, city2) == -1 {
			break
		}

		city2 = rand.Intn(len(permutation))
	}

	if city2 < city1 {
		city1, city2 = city2, city1
	}

	// reverse function will be applied to the permutation slice imidiately no need to assign the reversed
	// portion back to the slice
	reverse(permutation[city1:city2])

	return permutation
}

type Candidate struct {
	vector []int
	cost   int
}

func createNeighbor(current Candidate, cities []City) Candidate {
	candidate := Candidate{}

	// create a new epmty vector and copy the current vector into it since assigning this vector in go
	// would apply changes to both vectors (both are treated like a pointer to the same vector)
	candidate.vector = make([]int, len(current.vector))
	copy(candidate.vector, current.vector)
	candidate.vector = stochasticTwoOpt(candidate.vector)
	candidate.cost = cost(candidate.vector, cities)

	return candidate
}

func shouldAccept(candidate, current Candidate, temp float64) bool {
	if candidate.cost <= current.cost {
		return true
	}

	res := math.Exp(float64(current.cost-candidate.cost) / temp)
	rnd := rand.Float64()

	return res > rnd
}

func search(cities []City, iterations int, max_temp, temp_change float64) Candidate {
	current := Candidate{}
	current.vector = randomPermutation(cities)
	current.cost = cost(current.vector, cities)

	temp, best := max_temp, current

	for i := 0; i < iterations; i++ {
		candidate := createNeighbor(current, cities)
		temp = temp * temp_change

		if shouldAccept(candidate, current, temp) {
			current = candidate
		}

		if candidate.cost < best.cost {
			best = candidate
		}

		if (i+1)%100 == 0 {
			fmt.Println("Iteration", i+1, "temp=", temp, "best=", best)
		}
	}

	return best
}

type City struct {
	x int
	y int
}

func main() {
	rand.Seed(time.Now().Unix())

	cities := []City{{565, 575}, {25, 185}, {345, 750}, {945, 685}, {845, 655},
		{880, 660}, {25, 230}, {525, 1000}, {580, 1175}, {650, 1130}, {1605, 620},
		{1220, 580}, {1465, 200}, {1530, 5}, {845, 680}, {725, 370}, {145, 665},
		{415, 635}, {510, 875}, {560, 365}, {300, 465}, {520, 585}, {480, 415},
		{835, 625}, {975, 580}, {1215, 245}, {1320, 315}, {1250, 400}, {660, 180},
		{410, 250}, {420, 555}, {575, 665}, {1150, 1160}, {700, 580}, {685, 595},
		{685, 610}, {770, 610}, {795, 645}, {720, 635}, {760, 650}, {475, 960},
		{95, 260}, {875, 920}, {700, 500}, {555, 815}, {830, 485}, {1170, 65},
		{830, 610}, {605, 625}, {595, 360}, {1340, 725}, {1740, 245}}

	iterations := 2000
	max_temp := 100000.0
	temp_change := 0.98

	best := search(cities, iterations, max_temp, temp_change)

	fmt.Println(best.cost)

	fmt.Println("Result should be 7542.")
}
