package main

import "fmt"

func main() {
	all := make([]string, 0, 10)
	a := "BR221736106023MU"
	all = append(all, a)
	fmt.Println(all)

	subSlsTnLists := make([]string, 0, 10)
	b := "BR221736106023MU"
	subSlsTnLists = append(subSlsTnLists, b)
	set := DifferenceSet(all, subSlsTnLists)

	fmt.Println(set)

}

func DifferenceSet(a []string, b []string) []string {
	var c []string
	temp := map[string]struct{}{}
	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}
	for _, val := range a {
		if _, ok := temp[val]; !ok {
			c = append(c, val)
		}
	}
	return c
}
