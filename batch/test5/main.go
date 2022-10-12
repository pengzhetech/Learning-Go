package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//fmt.Println(time.Now().Unix())
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Int31())
}
