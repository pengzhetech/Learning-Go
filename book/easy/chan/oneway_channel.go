package main

import "fmt"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go first(c1)
	go second(c1, c2)

	Cook(c2)
	fmt.Println("洗碗")
}

func first(c chan<- string) {
	c <- "买菜:"
	close(c)
}

func second(c1 <-chan string, c2 chan<- string) {
	r := <-c1
	c2 <- r + "买肉:"
	close(c2)
}

func Cook(c <-chan string) {
	for r := range c {
		fmt.Println(r + "已经准备好,吃顿好的")
	}

}
