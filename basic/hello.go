package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	user1 := strings.Join([]string{"123", "23"}, " ")
	fmt.Println(user1)
	user2 := fmt.Sprintf("%s:%s", "pengzhe", "niubi")
	fmt.Println(user2)

	user3 := bytes.Buffer{}
	user3.WriteString("hello")
	user3.WriteString(" world")
	fmt.Println(user3.String())
}
