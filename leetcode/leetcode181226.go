package leetcode

import (
	"fmt"
	"strings"
)

func leetcode011(height []int) int {
	maxVol := 0

	for i := 0; i < len(height); i++ {
		for k := i + 1; k < len(height); k++ {
			min := height[i]
			if min > height[k] {
				min = height[k]
			}

			tmp := min * (k - i)
			if maxVol < tmp {
				maxVol = tmp
			}

		}
	}
	return maxVol
}

func leetcode012(num int) string {
	numArr := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romaArr := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	resArr := make([]string, 0)

	for i := len(numArr) - 1; i >= 0; i-- {
		n := num / numArr[i]
		num = num % numArr[i]

		for k := 0; k < n; k++ {
			resArr = append(resArr, romaArr[i])
		}
	}

	return strings.Join(resArr, "")
}

func leetcode013(s string) int {
	romaMap := map[string]int{"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, "XL": 40, "L": 50, "XC": 90, "C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000}
	sum := 0
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		nc := ""
		if i+1 < len(s) {
			nc = c + string(s[i+1])
		}

		if val, ok := romaMap[nc]; ok {
			sum += val
			i++
		} else {
			sum += romaMap[c]
		}

	}
	return sum
}

func leetcode014(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	strArr := strings.Split(strs[0], "")
	for i := 1; i < len(strs); i++ {
		if len(strArr) == 0 {
			return ""
		}
		for k, c := range strs[i] {
			if k >= len(strArr) || string(c) != strArr[k] {
				strArr = strArr[:k]
				break
			}
		}
	}
	return strings.Join(strArr, "")
}

func leetcode015(nums []int) [][]int {
	resArr := make([][]int, 0)
	thirdMap := make(map[int]int)
	for i, n := range nums {
		thirdMap[n*-1] = i
	}

	for i := 0; i < len(nums); i++ {
		for k := i + 1; k < len(nums); k++ {
			if val, ok := thirdMap[nums[i]+nums[k]]; ok && val != i && val != k {
				tmp := []int{nums[i], nums[k], -1 * (nums[i] + nums[k])}
				if !compare(resArr, tmp) {
					resArr = append(resArr, tmp)
				}
			}
		}
	}

	return resArr
}

func compare(allRes [][]int, single []int) bool {
	for _, templ := range allRes {
		amap := make(map[int]int)
		count := 0
		for _, s := range templ {
			amap[s] = 1
		}
		for _, s2 := range single {
			if _, ok := amap[s2]; ok {
				count++
			} else {
				delete(amap, s2)
			}
		}
		if count > 2 {
			return true
		}
	}

	return false
}

//LeetTest181226 Leetcode demo
func LeetTest181226() {
	// Leetcode011
	fmt.Println(leetcode011([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	// Leetcode012
	fmt.Println(leetcode012(4))

	// Leetcode013
	fmt.Println(leetcode013("LVIII"))

	// Leetcode014
	fmt.Println(leetcode014([]string{"dog", "racecar", "car"}))

	// Leetcode015
	fmt.Println(leetcode015([]int{-1, 0, 1, 2, -1, -4}))
}
