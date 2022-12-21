package main

import (
	"fmt"
	"hash/crc32"
	"strings"
)

func main() {

	hash := CRC32("35634654654654")
	fmt.Println(hash & (128 - 1))

	i := xiaoshu("123.0")
	fmt.Println(i)
}
func CRC32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

func xiaoshu(a string) int {
	numstr := fmt.Sprint(a)
	tmp := strings.Split(numstr, ".")
	if len(tmp) <= 1 {
		return 0
	}
	return len(tmp[1])
}
