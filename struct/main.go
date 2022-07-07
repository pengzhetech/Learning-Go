package main

import "fmt"
import "Learning-Go/struct/test"

func main() {
	/*hello := test.Instance().Hello("1234234")
	fmt.Println(hello)
	*/
	var in test.TestsStruct
	s := in.Hello("6666666")
	fmt.Println(s)
}
