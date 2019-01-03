package leetcode

import (
	"fmt"
	"math"
	"strings"
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

var charMap map[string][]byte
var resList []string

func leetcode017(digits string) []string {
	charMap = make(map[string][]byte)
	charMap["2"] = []byte{'a', 'b', 'c'}
	charMap["3"] = []byte{'d', 'e', 'f'}
	charMap["4"] = []byte{'g', 'h', 'i'}
	charMap["5"] = []byte{'j', 'k', 'l'}
	charMap["6"] = []byte{'m', 'n', 'o'}
	charMap["7"] = []byte{'p', 'q', 'r', 's'}
	charMap["8"] = []byte{'t', 'u', 'v'}
	charMap["9"] = []byte{'w', 'x', 'y', 'z'}

	resList = make([]string, 0)
	func017(strings.Split(digits, ""), 0, make([]byte, 0))
	return resList
}

func func017(digArr []string, index int, resArr []byte) {
	if len(digArr) == index {
		resList = append(resList, string(resArr))
		return
	}
	for _, c := range charMap[digArr[index]] {
		resArr = append(resArr, c)
		func017(digArr, index+1, resArr)
		resArr = resArr[:len(resArr)-1]
	}
}

//LeetTest181228 Leetcode demo
func LeetTest181228() {
	// fmt.Println(leetcode016([]int{-1, 2, 1, -4}, 1))

	fmt.Println(leetcode017("234"))
}
