package main

import (
	"fmt"
	"hash/crc32"
)

func main() {

	hash := CRC32("35634654654654")
	fmt.Println(hash & (128 - 1))
}
func CRC32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}
