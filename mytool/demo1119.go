package mytool

import (
	"fmt"
	"sync"
)

var ws sync.WaitGroup

const routinNum = 10
const singleCount = 10

// Main1119 entry
func Main1119() {
	// findOne()
	// testone()
	// findTwo()
	findThree()
}

func testone() {
	resArr := make([]int, 0)
	for i := 1; i < 10; i++ {
		resArr = append(resArr, i)

	}
	fmt.Println(resArr)

	twoArr := []int{11, 12, 14}
	resArr = append(resArr, twoArr...)
	fmt.Println(resArr)
}

func findThree() {
	numch := make([]chan int, routinNum)
	for i := 0; i < routinNum; i++ {
		numch[i] = make(chan int)
	}

	// for {
	for i := 0; i < routinNum; i++ {
		go withSingleChannel(i*singleCount, (i+1)*singleCount, numch[i])
		// go func(index int) {
		fmt.Printf("%d  ", <-numch[i])
		// }(i)

	}
	// }

}

func withSingleChannel(start, end int, ch chan int) {
	var flag bool
	for i := start; i < end; i++ {
		if i == 0 || i == 1 {
			continue
		}
		flag = true
		for k := 2; k < i; k++ {
			if i%k == 0 {
				flag = false
				break
			}
		}
		if flag {
			ch <- i
		}
	}
	// close(ch)
}

func findTwo() {
	resArr := make([]int, 0)
	for i := 0; i < routinNum; i++ {
		resArr = append(resArr, <-withChannel(i*singleCount, (i+1)*singleCount)...)
	}
	fmt.Println(resArr)
}

func withChannel(start, end int) chan []int {
	ch := make(chan []int)
	go func(start, end int) {
		resArr := make([]int, 0)
		var flag bool
		for i := start; i < end; i++ {
			if i == 0 || i == 1 {
				continue
			}
			flag = true
			for k := 2; k < i; k++ {
				if i%k == 0 {
					flag = false
					break
				}
			}
			if flag {
				resArr = append(resArr, i)
			}
		}
		ch <- resArr
	}(start, end)
	return ch
}

func findOne() {
	ws.Add(routinNum)
	for i := 0; i < routinNum; i++ {
		go findPrimeNum(i*singleCount, (i+1)*singleCount)
	}
	ws.Wait()
	fmt.Println("\n======================")
}

func findPrimeNum(start, end int) {
	defer ws.Done()

	var flag bool
	for i := start; i < end; i++ {
		if i == 0 || i == 1 {
			continue
		}
		flag = true
		for k := 2; k < i; k++ {
			if i%k == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Printf("%d  ", i)
		}
	}
}
