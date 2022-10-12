package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(CreateRandomNumber())
}

func CreateRandomNumber() string {
	var numbers = []byte{1, 2, 3, 4, 5, 7, 8, 9}
	var container string
	length := bytes.NewReader(numbers).Len()
	for i := 1; i <= 12; i++ {
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(length)))
		container += fmt.Sprintf("%d", numbers[random.Int64()])
	}
	return container
}
