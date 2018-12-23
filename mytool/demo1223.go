package mytool

import "fmt"

func leetcode01(numarr []int, target int) []int {
	resArr := make([]int, 0)
	arrMap := make(map[int]int)

	for index, value := range numarr {
		arrMap[value] = index
		if ntindex, ok := arrMap[target-value]; ok {
			resArr = append(resArr, ntindex, index)
		}
	}

	return resArr
}

type listNode struct {
	val  int
	next *listNode
}

func leetcode02(l1 *listNode, l2 *listNode) *listNode {
	pomotion := 0

	var head *listNode
	var rear *listNode

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.val
			l1 = l1.next
		}
		if l2 != nil {
			sum += l2.val
			l2 = l2.next
		}

		sum += pomotion
		if sum > 9 {
			pomotion = 1
			sum -= sum
		} else {
			pomotion = 0
		}

		node := &listNode{sum, nil}

		if head == nil {
			head = node
			rear = node
		} else {
			rear.next = node
			rear = node
		}
	}

	if pomotion > 0 {
		rear.next = &listNode{pomotion, nil}
	}

	return head
}

func makeListnode(numArr []int) *listNode {
	var head *listNode
	var rear *listNode

	for _, val := range numArr {
		node := &listNode{val, nil}

		if head == nil {
			head = node
			rear = node
		} else {
			rear.next = node
			rear = node
		}
	}

	return head
}

func leetcode03(s string) int {
	keyMap := make(map[int32]int)

	max := 1
	tmp := 0
	for _, c := range s {
		if _, ok := keyMap[c]; !ok {
			tmp = tmp + 1
			max = tmp

		} else {
			tmp = 1
			for k := range keyMap {
				delete(keyMap, k)
			}

		}
		keyMap[c] = 1
	}

	return max
}

func leetcode04(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	if len1 > len2 {
		nums1, nums2 = nums2, nums1
	}

	if (len1+len2)%2 == 0 {
		return (findMean(nums1, nums2, (len1+len2)/2) + findMean(nums1, nums2, (len1+len2)/2+1)) * 0.5

	} else {
		return findMean(nums1, nums2, (len1+len2+1)/2)
	}
}

func findMean(A []int, B []int, mean int) float64 {
	lenA := len(A)

	if lenA == 0 {
		return float64(B[mean-1])
	}
	if mean == 1 {
		if A[0] < B[0] {
			return float64(A[0])
		}
		return float64(B[0])
	}

	pa := mean / 2
	if pa > lenA {
		pa = lenA
	}

	pb := mean - pa

	if A[pa-1] < B[pb-1] {
		return findMean(A[pa:], B, pb)
	}

	return findMean(A, B[pb:], pa)
}

func leetcode05(str string) string {
	n := len(str)
	flag := true
	for i := n - 1; i >= 0; i-- {
		for k := 0; k+i < n; k++ {
			flag = true
			for s, e := k, i; s < e; {
				if str[s] != str[e] {
					flag = false
					break
				}
				s++
				e--
			}
			if flag {
				return str[k : i+1]
			}
		}
	}
	return "ERR"
}

//LeetTest Leetcode demo
func LeetTest() {
	// Leetcode01
	// fmt.Println(leetcode01([]int{3, 7, 2, 15}, 9))

	// Leetcode02
	l1 := makeListnode([]int{2, 4, 3})
	l2 := makeListnode([]int{5, 6, 4})
	head := leetcode02(l1, l2)
	for head != nil {
		fmt.Print(head.val, " ")
		head = head.next
	}

	// Leetcode03
	fmt.Println(leetcode03("pwwkew"))

	// Leetcode04
	fmt.Println(leetcode04([]int{1, 2}, []int{3}))

	// Leetcode05
	fmt.Println(leetcode05("cbasdfdsacd"))
}
