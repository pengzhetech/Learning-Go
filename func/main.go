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
	/*x := 2324 % 199
	fmt.Println(x)*/

	var hello Hi
	hello = Hello
	w := hello("3")
	fmt.Printf("%s\n", w)
	hello = Hello4Hunan
	h := hello("8")
	fmt.Printf("%s\n", h)

	var hello2 Hi = Hello
	s := hello2("4")
	fmt.Println(s)

	testFunc, i := TestFunc("1", "2")

	fmt.Println(testFunc, i)
}

func TestFunc(name, types string) (string, int) {
	fmt.Println("接受到的参数", name, types)
	return "1", 100000
}
