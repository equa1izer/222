package main

import "fmt"

func main() {
	case1 := []int{1, 2, 3, 4}
	out := TwoSum(case1, 5)
	fmt.Println(out)
}

func TwoSum(nums []int, target int) []int {

	x := len(nums)
	var result []int
	for i := 0; i < x; i++ {
		for j := i + 1; j < x; j++ {
			if nums[i]+nums[j] == target {
				result = []int{i, j}
				return result

			}
		}
	}
	return nil

}
