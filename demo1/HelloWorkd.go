package main

import "fmt"

var str string = "hello world"

func printStr(str string) {
	fmt.Println(str)
}

func main() {
	printStr(str)
}
