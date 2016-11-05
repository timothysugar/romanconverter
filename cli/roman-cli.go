package main

import (
	"fmt"
	"github.com/timothysugar/romanconverter/converter"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		PrintUsage()
		os.Exit(0)
	}

	var arg = os.Args[1]
	if i, err := strconv.Atoi(arg); err != nil {
		PrintUsage()
	} else {
		var res = converter.Convert(i)
		fmt.Println(res)
	}

}

func PrintUsage() {
	fmt.Printf("Usage: Converts a single integer.\ne.g >./cli 10")
}
