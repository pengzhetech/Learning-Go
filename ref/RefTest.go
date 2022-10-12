package main

import "fmt"

func main() {

	fmt.Println("test")
	user := User{Name: "彭哲", Age: 10}
	fmt.Printf("--------函数调用前¬------,%v", user)
	passBasic(user)
	fmt.Printf("--------函数调用后¬------,%v", user)
}

func passBasic(user User) {
	fmt.Println("函数内")
	user.Name = "内部彭哲"
	fmt.Println(user)
	fmt.Println("函数内结束")
}

type User struct {
	Name string
	Age  int
}
