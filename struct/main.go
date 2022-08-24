package main

import (
	"fmt"
	"time"
)
import "Learning-Go/struct/test"

func main() {
	/*hello := test.Instance().Hello("1234234")
	fmt.Println(hello)
	*/
	var in test.TestsStruct
	s := in.Hello("6666666")
	fmt.Println(s)

	fmt.Println(time.Now().Unix() + 7776000)
	fmt.Println(time.Now().Unix() - 7776000)

}
