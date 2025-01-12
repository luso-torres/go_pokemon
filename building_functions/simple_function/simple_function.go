package main

import (
	"fmt"
)

func main() {
	var value int
	outputText("Insert a value between 0 and 10: ")
	fmt.Scan(&value)
	outputText(fmt.Sprint(value))
}

func outputText(text1 string) {
	fmt.Print(text1)
}
