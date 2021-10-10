package main

import "fmt"

func main() {
	fmt.Println("test")
	var m map[string]int = map[string]int{"234": 1}
	i, ok := m["234"]
	fmt.Println(i, ok)
}
