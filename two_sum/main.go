package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	indx := make([]int, 2)
	for i, num := range nums {
		if _, ok := m[num]; ok {
			indx[0] = m[num]
			indx[1] = i
			return indx
		}
		m[target-num] = i
	}
	return nil
}

func main() {
	nums := []int{3, 4, 2}
	target := 6
	fmt.Println(twoSum(nums, target))
}
