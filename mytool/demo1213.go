package mytool

import (
	"fmt"
	"regexp"
)

// Main1213 Entrance function
func Main1213() {
	r, _ := regexp.Compile("(\\d+[.、]\\D+)")

	str := "12.ghhhjl哈哈哈;1、个哼哼3.2m唧唧；23.fghjk;34.34chjk"

	fmt.Println(r.FindAllString(str, -1))
}
