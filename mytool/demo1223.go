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

func LeetTest() {
	// fmt.Println(leetcode01([]int{3, 7, 2, 15}, 9))

	l1 := makeListnode([]int{2, 4, 3})
	l2 := makeListnode([]int{5, 6, 4})
	head := leetcode02(l1, l2)
	for head != nil {
		fmt.Print(head.val, " ")
		head = head.next
	}
}
