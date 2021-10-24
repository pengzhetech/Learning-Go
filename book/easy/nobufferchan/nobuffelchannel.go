package main

import "fmt"

/**
无缓冲channel
接受者没有能力保存任何值的channel
无缓冲channel必须要求执行发送操作的goroutine和直接接收操作的goroutine必须同时准备好,才能发送数据
如果两个goroutine没有同时准备好,则会导致先执行发送操作或者接收操作的goroutine处于阻塞状态
无缓冲channel对于发送和接收goroutine是同步的,无缓冲channel也称为同步channel
*/
func main() {
	fmt.Println("主协程开始运行")
	c := make(chan bool)
	go Eat("生蚝", c)
	fmt.Println("主协程运行结束")
	<-c
}

func Eat(foodName string, c chan bool) {
	fmt.Println("我正在吃" + foodName)
	c <- true
}
