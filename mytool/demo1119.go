package mytool

import (
	"fmt"
	"sync"
)

var ws sync.WaitGroup

const routinNum = 5
const singleCount = 1000

// Main1119 entry
func Main1119() {
	// testone()

	// findOne()
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

//=======================================================================
func findThree() {
	// numch := make([]chan int, routinNum)
	// for i := 0; i < routinNum; i++ {
	// 	numch[i] = make(chan int)
	// }

	// for i := 0; i < routinNum; i++ {
	// 	go withSingleChannel(i*singleCount, (i+1)*singleCount, numch[i])

	// 	go func(ch <-chan int) {
	// 		fmt.Printf("%d  ", <-ch)
	// 	}(numch[i])

	// }

	//----------------------------------------------
	// ch := make(chan int, 10)
	// for i := 0; i < routinNum; i++ {
	// 	go withSingleChannel(i*singleCount, (i+1)*singleCount, ch)
	// }

	// for res := range ch {
	// 	fmt.Printf("%d  ", res)
	// }

	//----------------------------------------------
	ch := make([]chan int, routinNum)
	for i := 0; i < routinNum; i++ {
		ch[i] = make(chan int)
	}
	for i := 0; i < routinNum; i++ {
		go withSingleChannel(i*singleCount, (i+1)*singleCount, ch[i])
	}

	for i := 0; i < routinNum; i++ {
		for res := range ch[i] {
			fmt.Printf("%d  ", res)
		}
	}

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
	close(ch)
}

//=======================================================================
func findTwo() {
	// resArr := make([]int, 0)

	// resArr = append(resArr, <-withChannel(i*singleCount, (i+1)*singleCount)...)

	select {
	case iarr0 := <-withChannel(0*singleCount, (0+1)*singleCount):
		fmt.Println(iarr0)
	case iarr1 := <-withChannel(1*singleCount, (1+1)*singleCount):
		fmt.Println(iarr1)
	case iarr2 := <-withChannel(2*singleCount, (2+1)*singleCount):
		fmt.Println(iarr2)
	case iarr3 := <-withChannel(3*singleCount, (3+1)*singleCount):
		fmt.Println(iarr3)
	case iarr4 := <-withChannel(4*singleCount, (4+1)*singleCount):
		fmt.Println(iarr4)
	}

	// for {
	// 	for i := 0; i < routinNum; i++ {
	// 		select {
	// 		case iarr0 := <-withChannel(i*singleCount, (i+1)*singleCount):
	// 			fmt.Println(iarr0)
	// 		default:
	// 			continue
	// 		}
	// 	}
	// }

	// fmt.Println(resArr)
}

func withChannel(start, end int) chan []int {
	ch := make(chan []int)
	go func(start, end int) {
		fmt.Println("============", start)
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

		// time.Sleep(1 * time.Second)

		ch <- resArr
	}(start, end)
	return ch
}

//=======================================================================
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
