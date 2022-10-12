package main

import "fmt"

func main() {
	//https://github.com/golang/go/wiki/SliceTricks
	actions := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	batchSize := 3
	batches := make([][]int, 0, (len(actions)+batchSize-1)/batchSize)

	for batchSize < len(actions) {
		actions, batches = actions[batchSize:], append(batches, actions[0:batchSize:batchSize])
	}
	batches = append(batches, actions)
	for _, value := range batches {
		fmt.Println(value)
	}
}
