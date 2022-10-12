package main

import "fmt"

func main() {
	actions := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	batchSize := 3
	fmt.Println((len(actions) + batchSize - 1) / batchSize)
	batches := make([][]int, 0, (len(actions)+batchSize-1)/batchSize)
	for batchSize < len(actions) {
		batches = append(batches, actions[0:batchSize])
		actions = actions[batchSize:]
	}
	batches = append(batches, actions)
	fmt.Println("result-------")
	fmt.Println(batches)
}
