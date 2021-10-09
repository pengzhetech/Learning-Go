package main

import "fmt"

func main() {
	var foods = []string{"红烧肉", "清蒸鱼", "刘大侠", "蒸螃蟹", "鲍鱼粥"}

	for i, item := range foods {
		res := fmt.Sprintf("%d--%q", i, item)
		fmt.Println(res)
	}

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

	/**
	从切片中删除元素
	1:找到待删除元素前面的切片
	2:找到待删除元素后面的切片
	3:把这两个切片组合起来,即可到达删除元素的目的
	*/

	var deleteFoods = []string{"1", "2", "3", "4"}
	fmt.Println(deleteFoods)
	deleteFoods2 := append(deleteFoods[:2], deleteFoods[3:]...)
	fmt.Println(deleteFoods2)

	testFoods := []string{"红烧肉", "清蒸鱼", "熘大虾", "蒸螃蟹", "鲍鱼粥"}
	testFoods2 := make([]*string, len(testFoods))
	for i, value := range testFoods {
		testFoods2[i] = &value
	}
	fmt.Println(testFoods[0], testFoods[1], testFoods[2], testFoods[3], testFoods[4])
	fmt.Println(*testFoods2[0], *testFoods2[1], *testFoods2[2], *testFoods2[3], *testFoods2[4])
	fmt.Println(testFoods2[0], testFoods2[1], testFoods2[2], testFoods2[3], testFoods2[4])
	fmt.Println("-----------------------------------------------------------")
	testFoods3 := make([]*string, len(testFoods))
	for i := range testFoods {
		testFoods3[i] = &testFoods[i]
	}
	fmt.Println(testFoods[0], testFoods[1], testFoods[2], testFoods[3], testFoods[4])
	fmt.Println(*testFoods3[0], *testFoods3[1], *testFoods3[2], *testFoods3[3], *testFoods3[4])
	fmt.Println(testFoods3[0], testFoods3[1], testFoods3[2], testFoods3[3], testFoods3[4])

}
