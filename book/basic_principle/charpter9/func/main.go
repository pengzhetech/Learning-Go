package main

import (
	"errors"
	"fmt"
)

func main() {
	if _, err := dlv(4, 2); err == nil {
		fmt.Println("err:", err)
	}
	fmt.Println("123243243")
}

func dlv(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("b<0")
	}
	return a / b, nil
}

func makeGreeter() func() string {
	return func() string {
		return "hello 哲哥"
	}
}

func visit(numbers []int, callback func(int2 int)) {
	for _, n := range numbers {
		callback(n)
	}
}
