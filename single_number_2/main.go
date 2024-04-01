package main

import "fmt"

// xor -> return true if true appears only once
// e.g : true(1) and false(0) -> true
// e.g : true(1) and true(1) -> false
// So, make this operation with numbers:
// 2 -> 0010
// 4 -> 0100
// ^ ---------
// The result will be 6(0110)
// bitwise or(|) -> if we already have a true(1) will return true
// e.g: 2 (0010) will return true because have it a 1
// operation with numbers:
// 2 -> 0010
// 4 -> 0100
// | --------
// The result be 6(0110)
func main() {
	nums := []int{2, 3, 3, 3, 2, 2, 1}
	fmt.Println(single_number(nums))
}

func single_number(nums []int) int {
	ones := 0
	twos := 0
	for _, i := range nums {
		ones = (ones ^ i) & ^twos
		twos = (twos ^ i) & ^ones
	}
	return ones
}
