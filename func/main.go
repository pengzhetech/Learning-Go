package main

import "fmt"

type Hi func(num string) string

func Hello(num string) string {
	return num + "位客人,欢迎光临"
}
func Hello4Hunan(num string) string {
	return num + "位湖南老板,里面请"
}

func main() {
	var hello Hi
	hello = Hello
	w := hello("3")
	fmt.Printf("%s\n", w)
	hello = Hello4Hunan
	h := hello("8")
	fmt.Printf("%s\n", h)
}
