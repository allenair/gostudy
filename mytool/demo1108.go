package mytool

import (
	"fmt"
)

func Main1108() {
	arr := [...]byte{1, 2, 3, 4, 5}
	fmt.Println(arr)

	b := append(make([]byte, 0), arr[:]...)
	fmt.Println(b)

	b = append(b, 5, 6, 7, 8, 9)
	fmt.Println(b)
}
