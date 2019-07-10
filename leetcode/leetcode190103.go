package leetcode

import (
	"fmt"
	"sort"
)

func leetcode018(nums []int, target int) [][]int {
	res := [][]int{}

	if len(nums) < 4 {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		sum3 := target - nums[i]
		for k := i + 1; k < len(nums)-2; k++ {
			if k > i+1 && nums[k] == nums[k-1] {
				continue
			}

			sum2 := sum3 - nums[k]
			left := k + 1
			right := len(nums) - 1

			for left < right {
				if nums[left]+nums[right] == sum2 {
					row := make([]int, 4)
					row[0] = nums[i]
					row[1] = nums[k]
					row[2] = nums[left]
					row[3] = nums[right]

					res = append(res, row)
					for {
						left++
						if !(left < right && nums[left-1] == nums[left]) {
							break
						}
					}
					for {
						right--
						if !(left < right && nums[right+1] == nums[right]) {
							break
						}
					}
				} else if nums[left]+nums[right] < sum2 {
					for {
						left++
						if !(left < right && nums[left-1] == nums[left]) {
							break
						}
					}

				} else {
					for {
						right--
						if !(left < right && nums[right+1] == nums[right]) {
							break
						}
					}
				}
			}
		}
	}

	return res
}

type linkNode struct {
	val  int
	next *linkNode
}

func leetcode019(head *linkNode, n int) *linkNode {
	ph := head
	pt := head

	count := 0
	for pt.next != nil {
		if count < n {
			pt = pt.next
			count++
		} else {
			ph = ph.next
			pt = pt.next
		}
	}

	ph.next = ph.next.next

	return head
}

func leetcode020(s string) bool {
	// sum := 0
	// for _, c := range s {
	// 	switch c {
	// 	case '{', '[', '(':
	// 		sum++
	// 	case '}', ']', ')':
	// 		sum--
	// 	}
	// 	if sum < 0 {
	// 		return false
	// 	}
	// }
	// if sum != 0 {
	// 	return false
	// }

	stack := make([]int32, 0)
	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, c)
		} else {
			last := stack[len(stack)-1]

			switch c {
			case '{', '[', '(':
				stack = append(stack, c)
			case '}', ']', ')':
				if c == '}' && last != '{' || c == ')' && last != ')' || c == ']' && last != '[' {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) > 0 {
		return false
	}

	return true
}

//LeetTest190103 Leetcode demo
func LeetTest190103() {
	// fmt.Println(leetcode018([]int{1, 0, -1, 0, -2, 2}, 0))

	// head := &linkNode{1, nil}
	// tmp := &linkNode{2, nil}
	// head.next = tmp

	// tmp.next = &linkNode{3, nil}
	// tmp = tmp.next

	// tmp.next = &linkNode{4, nil}
	// tmp = tmp.next

	// tmp.next = &linkNode{5, nil}
	// tmp = tmp.next

	// head = leetcode019(head, 2)

	// for head != nil {
	// 	fmt.Println(head.val)
	// 	head = head.next
	// }

	fmt.Println(leetcode020("[]{}"))
}
