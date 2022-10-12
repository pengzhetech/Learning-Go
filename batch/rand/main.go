package main

import (
	"fmt"
)

func main() {

	fmt.Println(1 << 0)
	s := LineSubTypeMap[1]
	fmt.Println(s)
}

var LineSubTypeMap = map[int64]string{
	1 << 0: "FL",
	1 << 1: "LM",
}
