package main

import "fmt"

type ChefInterface interface {
	GetHonor() string
}

type Chef struct {
	Name  string
	Age   int
	Honor string
}

func (c Chef) GetHonor() string {
	return c.Honor
}

func (c *Chef) SetHonor(title string) {
	c.Honor = title
}

func main() {
	zhang := Chef{
		Name:  "老张",
		Age:   36,
		Honor: "米其林1星",
	}
	fmt.Println(zhang.GetHonor())
	var ci ChefInterface = zhang
	zhang.SetHonor("我是三星")
	fmt.Println(zhang.GetHonor())
	fmt.Println(ci.GetHonor())
}
