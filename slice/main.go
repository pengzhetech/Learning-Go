package main

import "fmt"

func main() {
	var foods = []string{"红烧肉", "清蒸鱼", "刘大侠", "蒸螃蟹", "鲍鱼粥"}
	fmt.Println(foods)
	foods = append(foods, "三文鱼")
	fmt.Println(foods)

	fmt.Printf("foods容量:%d\n", cap(foods))
	//为切片扩展五个容量
	foods = append(foods, make([]string, 5)...)
	fmt.Printf("foods长度:%d\n", len(foods))
	fmt.Printf("foods容量:%d\n", cap(foods))
	fmt.Println(foods)
	//在索引为2的位置插入新元素
	foods = append(foods[:2], append([]string{"彭哲"}, foods[2:]...)...)
	fmt.Println(foods)
	//在索引4的位置插入长度为3的新切片
	foods = append(foods[:4], append(make([]string, 3), foods[4:]...)...)
	fmt.Println(foods)
	//在索引4的位置插入切片foods2中的所有元素
	foods2 := []string{"米饭", "面条", "馒头"}
	foods = append(foods[:4], append(foods2, foods[4:]...)...)
	fmt.Println(foods)
	//在末尾添加另外一个切片
	foods = append(foods[:], append(foods2)...)
	fmt.Println(foods)
}
