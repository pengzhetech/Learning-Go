package main

import "fmt"

func main() {
	fmt.Println("123")
	var sliceTests = []string{"123", "12"}
	fmt.Println(sliceTests)

	var sliceTests2 = make([]string, 10, 10)
	fmt.Println(sliceTests2)

	sliceTests3 := make([]string, 10, 10)
	fmt.Println(sliceTests3)

	foods := []string{"123", "1234", "134"}
	for inx, item := range foods {
		fmt.Println("index-->", inx)
		fmt.Println("item-->", item)
		fmt.Println("----------------")
	}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

}
