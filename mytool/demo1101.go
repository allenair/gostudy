package mytool

import (
	"encoding/base64"
	"fmt"

	"github.com/emirpasic/gods/lists/arraylist"
)

// Main1101 Entrance function
func Main1101() {
	fmt.Println("=====START=====")
	Test0816()

	list := arraylist.New()
	list.Add("a")

	myfunc(1, 2, 3, 4)

	fmt.Printf("Base64 String is %s", base64ToStr([]byte("Hello 王育才 Allen")))
}

func base64ToStr(bytes []byte) string {
	encodingRest := base64.StdEncoding.EncodeToString(bytes)

	return encodingRest
}

func myfunc(args ...int) {
	fmt.Println("11111111111")

	myfunc2(args...)
	myfunc2(args[2:]...)
}

func myfunc2(args ...int) {
	fmt.Println("222222222222")
	for _, data := range args {
		fmt.Println("2====", data)
	}
}

func cal() {
	arr := []int{0, 2, 3, 4, 5}
	cnt := 0
	maxVal := 0
	maxCnt := 0

	for _, a1 := range arr {
		for _, a2 := range arr {
			if a1 == a2 {
				continue
			}
			for _, a3 := range arr {
				if a1 == a3 || a2 == a3 {
					continue
				}
				for _, a4 := range arr {
					if a1 == a4 || a2 == a4 || a3 == a4 {
						continue
					}
					for _, a5 := range arr {
						if a1 == a5 || a2 == a5 || a3 == a5 || a4 == a5 {
							continue
						} else {
							x := a1*100 + a2*10 + a3
							y := a4*10 + a5
							if x > 100 && y > 10 {
								cnt++
								if maxVal < x*y {
									maxCnt = cnt
									maxVal = x * y
								}
								fmt.Printf("%d: (%d, %d), %d\n", cnt, x, y, x*y)
							}
						}
					}
				}

			}
		}
	}
	fmt.Printf("Max count is %d, Max is %d\n", maxCnt, maxVal)
}
