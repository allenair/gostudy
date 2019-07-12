package main

import (
	"fmt"
	"gostudy/leetcode"
	"sort"
)

func main() {
	fmt.Println("=====START=====")
	//mytool.Main1113()
	// mygin.MainGin()
	// mytool.Main1213()
	// fmt.Println(runtime.NumCPU())
	leetcode.LeetTest190228()

	// test()
}

func test() {
	nums := []int{2, 4, 1, 3, 55, 2, 1, 3333, 21, 34}

	sort.Ints(nums)
	fmt.Println(nums)
}
