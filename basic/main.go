package main

import (
	"fmt"
	"regexp"
	"time"
	"unicode/utf8"
)

func main() {
	fmt.Println("1625220821")
	timeNow := time.Unix(1625220821, 0) //2017-08-30 16:19:19 +0800 CST

	timeString := timeNow.Format("2006-01-02 15:04:05") //2015-06-15 08:52:32
	fmt.Println(timeString)
	fmt.Println(1 << 6)
	fmt.Println(1 << 2)
	fmt.Println(1 << 7)
	fmt.Println(1 << 5)

	ss := "Av.limonita , 1181, "
	fmt.Println(ss)
	sss := trimLastChar(ss)
	fmt.Println(sss)

}

//利用正则表达式压缩字符串，去除空格或制表符
func compressStr(str string) string {
	if str == "" {
		return ""
	}

	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "NA")
}

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size-1]
}
