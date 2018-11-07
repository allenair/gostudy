package mytool

import (
	"fmt"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// Print ...
func Print(str string) {
	fmt.Println("Tool:" + str)
}

// Test0816 ...
func Test0816() {
	mapt := map[string]int{"allen": 1, "aileen": 2, "bella": 3}
	for k, v := range mapt {
		fmt.Println(k, v)
	}
	simplifiedchinese.GBK.NewEncoder()
}
