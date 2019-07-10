package leetcode

import (
	"fmt"
	"math"
)

func leetcode06(s string, numRows int) string {
	lens := len(s)
	row := 0
	row = row + lens/(numRows+numRows-2)*(1+numRows-2)
	row = row + (lens%(numRows+numRows-2))/numRows
	row = row + (lens%(numRows+numRows-2))%numRows

	charArr := make([][]int32, row)
	count := 1
	row = -1
	for _, c := range s {
		if count%(numRows+numRows-2) == 1 {
			count = 1
			row++
			charArr[row] = make([]int32, 4)
			charArr[row][0] = c
		} else if count <= numRows {
			charArr[row][count-1] = c
		} else {
			row++
			charArr[row] = make([]int32, 4)
			charArr[row][numRows-1-(count-numRows)] = c
		}
		count++
	}

	res := make([]int32, len(s))

	count = 0
	for k := 0; k < 4; k++ {
		for i := 0; i < len(charArr); i++ {
			if charArr[i][k] > 0 {
				res[count] = charArr[i][k]
				count++
			}

		}
	}

	return string(res)

}

func leetcode007(x int) int {
	flag := 1
	if x < 0 {
		flag = -1
	}

	x = x * flag

	numArr := []int{}
	for x > 0 {
		numArr = append(numArr, x%10)
		x = x / 10
	}

	var res int64 = 0
	for _, n := range numArr {
		res = res*10 + int64(n)
	}

	if res > math.MaxInt32 {
		return 0
	}

	return int(res) * flag
}

func leetcode008(str string) int {
	start := false
	flag := 1
	numArr := []int32{}
	for _, c := range str {
		switch {
		case !start && (c == 9 || c == 32):
			continue
		case c == 45:
			flag = -1
			start = true
		case c == 43:
			flag = 1
			start = true
		case c >= 48 && c <= 57:
			numArr = append(numArr, c)
			start = true
		default:
			if !start || (c == 9 || c == 32) && len(numArr) == 0 {
				return 0
			} else if len(numArr) > 0 {
				break
			}
		}
	}

	var res int64
	for _, n := range numArr {
		res = res*10 + int64(n-48)
	}

	if res > math.MaxInt32 {
		if flag > 0 {
			return math.MaxInt32
		}
		return math.MinInt32

	}

	return int(res) * flag
}

func leetcode009(x int) bool {
	if x < 0 {
		return false
	}

	numArr := []int{}

	for x > 0 {
		numArr = append(numArr, x%10)
		x = x / 10
	}

	for i, k := 0, len(numArr)-1; i < k; i++ {
		if numArr[i] != numArr[k] {
			return false
		}
		k--
	}
	return true
}

//LeetTest181224 Leetcode demo
func LeetTest181224() {
	// Leetcode06
	// fmt.Println(leetcode06("LEETCODEISHIRING", 3))

	// Leetcode007
	// fmt.Println(leetcode007(math.MaxInt32))

	// Leetcode008
	// fmt.Println(leetcode008("-91283472332"))

	// Leetcode009
	fmt.Println(leetcode009(12))
}
