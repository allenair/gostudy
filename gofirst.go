package main

import (
	"fmt"
	"gostudy/mytool"
	"runtime"
)

func main() {
	fmt.Println("=====START=====")
	mytool.Main1108()
	// mygin.MainGin()

	fmt.Println(runtime.NumCPU())
}
