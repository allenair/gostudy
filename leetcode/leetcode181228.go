package leetcode

import (
	"fmt"
	"math"
)

func leetcode016(nums []int, target int) int {
	diff := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				tmp := nums[i] + nums[j] + nums[k] - target
				if tmp < 0 {
					tmp *= -1
				}
				if tmp < diff {
					diff = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return diff
}

func leetcode017(digits string) []string {
	charMap := make(map[int][]byte)
	charMap[2] = []byte{'a', 'b', 'c'}
	charMap[3] = []byte{'d', 'e', 'f'}
	charMap[4] = []byte{'g', 'h', 'i'}
	charMap[5] = []byte{'j', 'k', 'l'}
	charMap[6] = []byte{'m', 'n', 'o'}
	charMap[7] = []byte{'p', 'q', 'r', 's'}
	charMap[8] = []byte{'t', 'u', 'v'}
	charMap[9] = []byte{'w', 'x', 'y', 'z'}

}

//LeetTest181228 Leetcode demo
func LeetTest181228() {
	fmt.Println(leetcode016([]int{-1, 2, 1, -4}, 1))
}
