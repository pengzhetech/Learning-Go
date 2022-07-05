package main

import "fmt"

func main() {
	fmt.Println("test")
	var m = map[string]int{"234": 1}
	i, ok := m["234"]
	fmt.Println(i, ok)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("求和结果:", sum)
}
