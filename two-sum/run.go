// Based On: https://leetcode.com/problems/two-sum/?tab=Description
package main

import "fmt"

// Time complexity: O(n^2): because for each element we loop through the whole array
// Space complexity: O(1)
func twoSumBruteForce(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{0, 0}
}

// Time complexity: O(n): we build a hash of values O(n). then iterate the hash and do O(1) complement lookup
// Space complexity: O(n): we need to build a hash of all values
func twoSumTwoPassHash(nums []int, target int) []int {
	vals := make(map[int]int)

	for k, v := range nums {
		vals[v] = k
	}

	for k, v := range nums {
		diff := target - v

		i, ok := vals[diff]

		if ok && k != i {
			return []int{k, i}
		}
	}

	return []int{0, 0}
}

func main() {
	fmt.Print(twoSumBruteForce([]int{3, 2, 4}, 6))
	fmt.Println(twoSumBruteForce([]int{1, 1, 5}, 2))

	fmt.Print(twoSumTwoPassHash([]int{3, 2, 4}, 6))
	fmt.Print(twoSumTwoPassHash([]int{1, 1, 5}, 2))
}
