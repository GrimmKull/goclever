package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//utils
func randLayerWeightVector(n int) []float64 {
	max := 1.0
	min := -1.0

	w := make([]float64, n)

	for i, _ := range w {
		w[i] = min + (max-min)*rand.Float64()
	}

	return w
}

//transfer functions
func step(activation float64) float64 {
	if activation >= 0 {
		return 1.0
	} else {
		return 0.0
	}
}

//neural nets
type Perceptron struct {
	epochs       int
	weights      []float64
	input        [][]int
	target       []int
	learningRate float64
}

func updateWeights(weights []float64, input []int, target, output, learningRate float64) []float64 {
	for i, val := range input {
		weights[i] += learningRate * (target - output) * float64(val)
	}

	//add bias
	weights[len(input)] += learningRate * (target - output) * 1.0

	return weights
}

func activate(weights []float64, input []int) float64 {
	sum := weights[len(weights)-1] * 1.0
	for i, val := range input {
		sum += weights[i] * float64(val)
	}

	return sum
}

func train(weights []float64, inputs [][]int, targets []int, epochs int, learningRate float64, treshold float64) []float64 {
	for k := 1; k <= epochs; k++ {
		error := 0.0

		for i, input := range inputs {
			output := step(activate(weights, input))

			error += math.Abs(output - float64(targets[i]))
			weights = updateWeights(weights, input, float64(targets[i]), output, learningRate)

			//fmt.Print(output)
		}
		//fmt.Println()
		fmt.Println("epoch:", k, ", error:", error)

		if error <= treshold {
			break
		}
	}

	return weights
}

func getOutputs(weights []float64, inputs [][]int) []float64 {
	outputs := make([]float64, len(inputs))

	for i, input := range inputs {
		output := step(activate(weights, input))

		outputs[i] = output
	}

	return outputs
}

//run
func main() {
	rand.Seed(time.Now().Unix())

	//problem := [][]int{{0,0,0},{0,1,1},{1,0,1},{1,1,1}}
	inputs := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	targets := []int{0, 1, 1, 1}

	epochs := 20
	learningRate := 0.1

	treshold := 0.0

	//fmt.Println(randVector(5))
	weights := randLayerWeightVector(len(inputs[0]) + 1)

	weights = train(weights, inputs, targets, epochs, learningRate, treshold)

	fmt.Println(weights)

	fmt.Println(getOutputs(weights, inputs))
	fmt.Println(targets)
}
