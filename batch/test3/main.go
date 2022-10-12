package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const letterBytes = "0123456789"
const (
	length        = 10                   // 6 bits to represent a letter index
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func RandStringBytesMask() int32 {
	b := make([]byte, length)
	for i := 0; i < length; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	num, _ := strconv.ParseInt(string(b), 10, 32)
	return int32(num)
}

func main() {
	fmt.Println(RandStringBytesMask())
}
